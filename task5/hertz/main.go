// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"mission/task5/hertz/biz/handler"
)

func main() {
	handler.InitConn()
	h := server.Default()

	register(h)
	h.Spin()
}
