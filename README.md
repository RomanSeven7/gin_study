# gin_study

## gin简单的router 启动一个http服务

### 基本的router 注册

```go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello q1mi!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
```

优点：

适用于路由条目比较少的简单项目或者项目demo。

