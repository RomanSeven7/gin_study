#gin与 swagger的集成

[swagger的参考文档]: https://swagger.io/docs/specification/about/



## swagger的介绍

### 什么是swagger

Swagger 是一个规范和一套完整的框架，用于生成、描述、调用以及可视化 RESTful 风格的 Web 服务。

Swagger的总体目标是使客户端和文件系统服务器以同样的速度来更新，方法，参数和模型紧密集成到服务器端的代码中，允许API始终保持同步。

Swagger 让部署管理和使用API从未如此简单。

Swagger包括库、编辑器、代码生成器等很多部分，这里我们主要讲一下Swagger Editor。这是一个完全开源的项目，并且它也是一个基于Angular的成功案例，我们可以下载源码并自己部署它，也可以修改它或集成到我们自己的软件中。

在Swagger Editor中，我们可以基于YAML语法定义我们的RESTful API，然后它会自动生成一篇排版优美的API文档，并且提供实时预览。相信大多数朋友都遇到过这样一个场景：明明调用的是之前约定好的API，拿到的结果却不是想要的。

可能因为是有人修改了API的接口，却忘了更新文档；或者是文档更新的不及时；又或者是文档写的有歧义，大家的理解各不相同。总之，让API文档总是与API定义同步更新，是一件非常有价值的事。

### 自动文档的好处？

1. 不用手动写文档了，通过注解就可以自动化文档

2. 文档和代码同步更新，代码更新之后不需要再更新文档

3. 浏览器友好

4. 使用Swagger框架可以调试API，在浏览器端可以看到更多的`request`和`response`信息

### 自动化文档开发的初衷

我们需要开发一个API应用，然后需要和手机组的开发人员一起合作，当然我们首先想到的是文档先行，我们也根据之前的经验写了我们需要的API原型文档，我们还是根据github的文档格式写了一些漂亮的文档，但是我们开始担心这个文档如果两边不同步怎么办？因为毕竟是原型文档，变动是必不可少的。手机组有一个同事之前在雅虎工作过，他推荐我看一个swagger的应用，看了swagger的标准和文档化的要求，感觉太棒了，这个简直就是神器啊，通过swagger可以方便的查看API的文档，同时使用API的用户可以直接通过swagger进行请求和获取结果。所以我就开始学习swagger的标准，同时开始进行Go源码的研究，通过Go里面的AST进行源码分析，针对comments解析，然后生成swagger标准的json格式，这样最后就可以和swagger完美结合了。

这样做的好处有三个：

注释标准化
有了注释之后，以后API代码维护相当方便
根据注释自动化生成文档，方便调用的用户查看和测试

## 安装swagger

```go
go get -u github.com/swaggo/swag/cmd/swag
```

等待安装完成，在我们的终端中执行 `swag init`，目录为根目录，于 `main.go` 同目录。

执行完成后，会在根目录下新建一个 `docs` 文件夹。

```shell
docs
|
|-docs.go
|-swagger.json
|-swagger.yaml
```

接下来就可以完善项目了。

将下面两行放入 `initRouter` 中的 `import` 中。

```
swaggerFiles "github.com/swaggo/files"
ginSwagger "github.com/swaggo/gin-swagger"
复制代码
```

选择 `Sync packages of GinStudy`,此时 `IDE` 就会自动帮我下载，并添加到 `go.mod` 中。

## 集成swagger

对 `swagger` 安装完成后，我们就可以对项目进行集成了。

在 `initRouter` 中添加路由，这个路由是对 `swagger` 的访问地址来进行添加的

在`routers/routers.go`的`Init`方法中添加

```go
url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
```

其中 `url` 定义了 `swagger` 的 `doc.json` 路径，我们可以直接访问该 `json` 来进行查看。

接下来就是完善文档的时间。

在 `main.go` 中 `main` 方法上添加注释。同时引入我们生成 `docs.go`

```go
// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目
// @contact.name
// @contact.url https://youngxhui.top
// @contact.email youngxhui@g mail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	// 省略其他代码
}
```

上述的注释基本都是很好理解的，不做过多解释。

主要的项目介绍注释就是这些，接下来进行我们的接口方法注释。

在我们的 `handler` 中添加注释

打开 `app/user/handler.go` ,在 `CreateUser` 方法上添加。

```go
// @Summary 创建用户
// @Tags 用户模块
// @version 1.0
// @Accept application/x-www-form-urlencoded
// @Param name query string true "name"
// @Param age query int true "age"
// @Success 200 object model.UserModel 成功后返回值
// @Router  /v1/users [post]
func CreateUser(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	ageInt, _ := strconv.Atoi(age)
	basicHandle := app.BasicController{Ctx: c}
	basicHandle.Ok(userService.Create(name, ageInt, time.Now()))
}
```

- @Summary 是对该接口的一个描述
- @Tags 是对接口的标注，同一个 tag 为一组，这样方便我们整理接口
- @Version 表明该接口的版本
- @Accept 表示该该请求的请求类型
- @Param 表示参数 分别有以下参数 参数名词 参数类型 数据类型 是否必须 注释 属性(可选参数),参数之间用空格隔开。
- @Success 表示请求成功后返回，它有以下参数 请求返回状态码，参数类型，数据类型，注释
- @Failure 请求失败后返回，参数同上 (todo)
- @Router 该函数定义了请求路由并且包含路由的请求方式。


具体参数类型，数据类型等可以查看[官方文档](https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html)

其中文档中没有说明的地方这里说明一下，关于 `Param` 的参数类型有以下几种

- query 形如 `\user?username=Jack&age=18`
- body 需要将数据放到 body 中进行请求
- path 形如 `\user\1`
- formdata 接收的是form表单提交的参数

不同的参数类型对应的不同请求，请对应使用。

这样我们就完成了添加接口的文档注释。

我们对形如 `/v1/users/:id` 的接口，最后的 id 通过 `{}` 包裹。

细心的小伙伴可能会发现我们最后的返回结果为 `model.Result` ，这是为了我们统一返回结果而新建的一个结构体，方便前端进行解析。具体函数如下

```go
package model

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}

```

我们在对 `Result` 中的 `tag` 会有 `example` ,这个仍旧是 `swagger` 的标签，用来给该结构体一个示例。

同理，我们可以对之前的 `article` 进行注释。

当我们完成了所有的代码注释时，在控制台中重新执行 `swag init`，它会根据我们的注释生成 `docs.go` 及其对应的 json 和 yaml 文件。

启动我们的项目，访问 `hppt://localhost:8080/swagger/index.html` 就可以查看我们的文档,效果如下


![效果图](https://github.com/RomanSeven7/gin_study/blob/master/pic/swagger_demo.png)

