# echo-query-params

## 複数パラメータ
- サンプルコード中の `Slices` フィールドのように、スライスの構造体を定義することで、
複数パラメータのBindが可能となる。
- パラメータ例: http://localhost:8080/?slices=a&slices=b&slices=c
- スライスではなく、通常の型(stringとか)の場合は最初に送られてきた `a` のみがBindされ、 `b`, `c` は無視される。

## 構造体の埋め込み
- サンプルコード中の `Embed` フィールドのように、構造体の埋め込みを行ってBindすることも可能。
- 埋め込む構造体は exported である必要がある。
- パラメータ例: http://localhost:8080/?slices=a&slices=b&slices=c&embed=d

## Bind処理を複数回行う
- サンプルコードのように、一度 `query` 構造体へのBindを行った後に、 `another` 構造体へのBindを行うことも可能。
- でもバリデーションも複数回やらないといけなくなって面倒な気もする。
- パラメータ例: http://localhost:8080/?slices=a&slices=b&slices=c&embed=d&other=e
