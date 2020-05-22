package main

import (
	"fmt"
	"mfx/gin_study/app/order"
	"mfx/gin_study/app/user"
	"mfx/gin_study/routers"
)

func main() {
	// 加载多个app的路由配置
	routers.Include(order.Routers,user.Routers)
	// 初始化路由
	r:=routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
