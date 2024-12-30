# taskhop
Prototype. Remote Task Runner

## 目的
ubuntu on wsl2 でプログラミングをしたいが windows のキー配置に慣れてなく捗らない。そのため mac から ubuntu にアクセスしたい

## 方法
- mac でバイナリを実行 (commander)
- ubuntu on wsl2 でバイナリを実行 (runner)

commander はWebサーバを立ち上げる。runner は commander のアドレスをポーリングする。commander は「実行命令」をレスポンスするので runner はそれを実行する

## Commands
```bash
# commander
taskhop

# runner
taskhop-runner
```
