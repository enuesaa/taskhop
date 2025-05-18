# taskhop
Prototype. Remote Task Runner

[![ci](https://github.com/enuesaa/taskhop/actions/workflows/ci.yaml/badge.svg)](https://github.com/enuesaa/taskhop/actions/workflows/ci.yaml)

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

- agent と commander を connect するときに、両方のコマンドをいくつか操作する必要があり、ちょっと頭が混乱する
- 大きな仕組みは変えず、CLI のフラグなど見直して改善したい
