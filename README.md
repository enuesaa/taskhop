# taskhop
Prototype. Remote Task Runner

## 目的
ubuntu on wsl2 でプログラミングをしたいが windows のキー配置に慣れてなく捗らない。
そのため mac から ubuntu にアクセスしたい

## 方法
- ubuntu on wsl2 でバイナリを実行 (runner)
- mac でもバイナリを実行 (commander)

commander はWebサーバを立ち上げ、runner がそのIPアドレスをポーリングする。必要な時に commander は「コマンド実行命令」をレスポンスするので、runner はそれを実行する
