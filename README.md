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

## Future plan
- cmds.yml を用意するのが面倒なので、廃止
- 代わりに prompt にする
- Session Manager ライクに。

```bash
taskhop -transfer
# show address

taskhop-runner -connect localhost:3000
```
