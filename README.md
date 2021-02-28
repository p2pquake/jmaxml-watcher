# jmaxml-watcher

気象庁防災情報 XML の更新監視

## 主な機能

- 更新を検出すると、 `--after-hook` で指定したコマンドの標準入力に XML を渡します
- `Cache-Control` HTTP ヘッダを見て、最適な更新間隔でチェックします

## 使用方法

[Releases](https://github.com/p2pquake/jmaxml-watcher/releases) から実行可能なバイナリが入手可能です。

```sh
$ ./jmaxml-watcher  --help
気象庁防災情報 XML の更新監視

Usage:
  jmaxml-watcher [flags]

Flags:
  -a, --after-hook string   実行するコマンド (標準入力に XML を渡します)
  -h, --help                help for jmaxml-watcher

$ ./jmaxml-watcher  --after-hook run_updated_actions.sh
2021/03/01 01:17:04 No update
2021/03/01 01:17:04 Next run after 55 seconds (maxAge: 53)
2021/03/01 01:17:58 Detect update
2021/03/01 01:18:11 Next run after 10 seconds (maxAge: 14)
2021/03/01 01:18:21 No update
2021/03/01 01:18:21 Next run after 41 seconds (maxAge: 39)

# (Ctrl+C で停止します)
```

## 参考

- [気象庁 | 気象庁防災情報XMLフォーマット](http://xml.kishou.go.jp/)

## ライセンス

- MIT License
