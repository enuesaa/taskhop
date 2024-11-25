# designdoc

## 目的
ubuntu on wsl2 でプログラミングをしたいが windows のキー配置に慣れてなく捗らない。
そのため mac から ubuntu にアクセスしたい

## 方法
- ubuntu on wsl2 でバイナリを実行 (runner)
- mac でもバイナリを実行 (commander)

runner は commander のIPアドレスをポーリングして、もし命令がレスポンスされてきたら、実行する

## Stacks
### echo
- シンプルに安定性を求めるため

### gqlgen
- GraphQL でファイルアップロードをどうすべきか知識がないので、gqlgen で試してみたい
- see https://gqlgen.com/reference/file-upload/

### sveltekit
- ちょっと凝ったUIを試してみたい
