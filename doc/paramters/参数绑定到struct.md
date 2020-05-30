# gin框架的参数绑定
为了能够更方便的获取请求相关参数，提高开发效率，我们可以基于请求的Content-Type识别请求数据类型并利用反射机制自动提取请求中QueryString、form表单、JSON、XML等参数到结构体中。 下面的示例代码演示了.ShouldBind()强大的功能，它能够基于请求自动提取JSON、form表单和QueryString类型的数据，并把值绑定到指定的结构体对象。

## 使用
1、我们创建公共参数的结构体
```go
type CommonParams struct {
	Pkg string `json:"pkg" form:"pkg"`
	Vn string `json:"vn" form:"vn"`
	Ts int64 `json:"ts" form:"ts"`
}
```
2、在router中去使用
```go
var commonParams model.CommonParams
	if err:=c.ShouldBind(&commonParams);err!=nil{
		logrus.Error(err)
	}
```

## 测试
请求：
```shell script
curl --location --request GET 'http://localhost:8080/v1/user?pkg=test.1&vn=1.1.1&ts=123'

```
返回：
```shell script
{
    "commonParams": {
        "pkg": "test.1",
        "vn": "1.1.1",
        "ts": 123
    },
    "message": "load user success"
}
```

## 原理剖析
我们调用的Bind方法实际调用了两个方法binding.Default和c.MustBindWith，前者的主要作用是根据终端请求的Content-Type选择处理器，没办法，gin太强，支持的类型太多了。我们刚才的请求方式被理所应当的分配给了formBinding。
```shell script
var (
	JSON          = jsonBinding{}
	XML           = xmlBinding{}
	Form          = formBinding{}
	Query         = queryBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
	ProtoBuf      = protobufBinding{}
	MsgPack       = msgpackBinding{}
	YAML          = yamlBinding{}
	Uri           = uriBinding{}
	Header        = headerBinding{}
)
```

然后呢？在拿到了formBinding之后，就来到了c.MustBindWith方法，它的作用就是调用formBinding的Bind方法，
Bind方法主要就干了两件事情，第一是解析传过来的表单参数，第二是找到结构体里的tagform进行匹配赋值。看到这里我就明白了，原来只需要在Login后面加上一个tag

那么是如何解析tag的呢 ? 其实是通过反射获取struct中StructFiel获取tag,然后指定key至
```shell script

func mapping(value reflect.Value, field reflect.StructField, setter setter, tag string) (bool, error) {
	if field.Tag.Get(tag) == "-" { // just ignoring this field
		return false, nil
	}
...
}
```


