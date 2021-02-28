package jmaxml

import (
	"io"
	"log"
	"math"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

func RunWatcher(persistent bool, afterHook string) {
	maxAge := regexp.MustCompile("\\d+")

	oldXML := ""
	c := time.After(1 * time.Second)
	for {
		select {
		case <-c:
			// XML とる
			retreiveTime := time.Now()
			resp, err := http.Get("http://www.data.jma.go.jp/developer/xml/feed/eqvol.xml")
			if err != nil {
				log.Printf("Retreive error: %v", err)
				continue
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Retreive error: %v", err)
				continue
			}

			xml := string(body)

			// 比較する
			if oldXML != xml {
				if oldXML != "" {
					log.Println("Detect update")
					hook(afterHook, xml)
				}

				oldXML = xml
			} else {
				log.Println("No update")
			}

			// Cache-control ヘッダの max-age を見て設定する
			cacheControl := resp.Header.Get("Cache-Control")
			maxAgeStr := maxAge.FindString(cacheControl)
			maxAge, err := strconv.ParseFloat(maxAgeStr, 64)

			interval := 40.0
			if err == nil {
				interval = maxAge + 2
			}
			interval = math.Max(10.0, interval-(time.Now().Sub(retreiveTime)).Seconds())

			log.Printf("Next run after %.0f seconds (maxAge: %s)\n", interval, maxAgeStr)
			c = time.After(time.Duration(interval) * time.Second)
		}
	}

}

func hook(command string, data string) {
	cmd := exec.Command(command)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Printf("hook execute error: %v", err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, data)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("hook output error: %v", err)
	}

	log.Println(string(out))
}
