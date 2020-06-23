# echo-request-id
`labstack/echo` + `labstack/gommon/log` を利用してログにリクエストIDを出力するためのサンプルコードです。

## 使い方
1. `e.Use(middleware.RequestID())` を追加します。
2. `e.Use(middlewares.RequestIDInLogger())` を追加します。

## 注意点
- `e.Use(middleware.RequestID())` で生成したIDを利用しているため、これを忘れるとIDが出力されなくなります。
