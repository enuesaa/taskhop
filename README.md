# taskhop
Prototype. Remote Task Runner

[![ci](https://github.com/enuesaa/taskhop/actions/workflows/ci.yaml/badge.svg)](https://github.com/enuesaa/taskhop/actions/workflows/ci.yml)

## 動機
Ubuntu on WSL2 で開発したいが Windows のキー配置に慣れてなく捗らない。そのため Mac から Ubuntu へアクセスしたい

## 仕組み
2つのコマンドを用意

```bash
taskhop
taskhop-agent
```

1. Mac で `taskhop` を実行しWebサーバを立ち上げる。
2. Ubuntu on WSL2 で `taskhop-agent` を実行し 1 のアドレスをポーリングする
3. `taskhop-agent` は `taskhop` より実行命令を受け取り実行する

## Example
```bash
# commander
taskhop

# runner
taskhop-agent localhost:3000
```

## Feature Plan
- agent を立ち上げる際に commander への接続情報が必要であり、両方を操作する必要があるので、ちょっと頭が混乱する
