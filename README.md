# taskhop
Prototype. Remote Task Runner

[![ci](https://github.com/enuesaa/taskhop/actions/workflows/ci.yaml/badge.svg)](https://github.com/enuesaa/taskhop/actions/workflows/ci.yaml)

## 動機
Ubuntu on WSL2 で開発したいが Windows のキー配置に慣れてなく捗らない。そのため Mac から Ubuntu へアクセスしたい

## 仕組み
2つのコマンドを用意

```bash
taskhop
taskhop-runner
```

1. Mac で `taskhop` を実行しWebサーバを立ち上げる。
2. Ubuntu on WSL2 で `taskhop-runner` を実行し 1 のアドレスをポーリングする
3. `taskhop-runner` は `taskhop` より実行命令を受け取り実行する

## Example
```bash
# commander
taskhop

# commander (transfer assets)
taskhop --transfer

# runner
taskhop-runner --connect localhost:3000
```
