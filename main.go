package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mfx/gin_study/routers"
)

func main() {
	r := gin.Default()
	routers.LoadOrder(r) // 加载order 模块的router
	routers.LoadUser(r) // 加载 user 模块的路由
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
