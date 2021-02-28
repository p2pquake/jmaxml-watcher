package cmd

import (
	"fmt"

	"github.com/p2pquake/jmaxml-watcher/jmaxml"
	"github.com/spf13/cobra"
)

// ビルド時に設定
var (
	Version = "develop"
	Commit  = "unknown"
	Date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:     "jmaxml-watcher",
	Short:   "気象庁防災情報 XML の更新監視",
	Version: fmt.Sprintf("%s (commit %s, built at %s)", Version, Commit, Date),
	Run: func(cmd *cobra.Command, args []string) {
		// jmaxml.RunWatcher(!nonPersistent, afterHook)
		jmaxml.RunWatcher(false, afterHook)
	},
}

var afterHook string
var nonPersistent bool

func init() {
	rootCmd.Flags().StringVarP(&afterHook, "after-hook", "a", "", "実行するコマンド (標準入力に XML を渡します)")
	// FIXME: 未実装
	// rootCmd.Flags().BoolVarP(&nonPersistent, "non-persistent", "n", false, "状態を永続化しない")
}

func Execute() error {
	return rootCmd.Execute()
}
