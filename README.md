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

## Feature Plan

- runner と commander を connect するときに、両方のコマンドをいくつか操作する必要があり、ちょっと頭が混乱する
- 大きな仕組みは変えず、CLI のフラグなど見直して改善したい
- デフォルトでファイル転送する (--transfer フラグの廃止)

```bash
$ taskhop
Taskhop Commander started!
┌─────────────────────────────────────────────────────────────────
│ Please start Taskhop Agent with following command:
│   taskhop-agent -c localhost:3000
└─────────────────────────────────────────────────────────────────

$ taskhop-agent -c localhost:3000
```

