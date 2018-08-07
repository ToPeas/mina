# mina
Golang 微信小程序 SDK

## done
登录和解密

## todo
生成二维码

## usage
```go
package mina

import (
	"github.com/solarhell/mina/debug"
	"net/http"
)

func main() {
	c := NewClient(&http.Client{
		Transport: &debug.DebugRequestTransport{
			RequestHeader:  true,
			RequestBody:    true,
			ResponseHeader: true,
			ResponseBody:   true,
		},
	})

	ui, err := c.Login("appid", "secret", "code")
	...
}
```