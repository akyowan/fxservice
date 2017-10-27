package main

import "fxservice/service/weixinpay/app"

func main() {
	addr := ":8010"
	app.Start(addr)
}
