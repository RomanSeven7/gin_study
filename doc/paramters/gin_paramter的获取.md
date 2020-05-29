# gin获取请求的参数信息
【注】c 就是func中的参数(c  *gin.Contex)



## 从urlpath中获取



请求:

``` shell script

curl -X GET \
  http://localhost:8080/v1/order/1 \
  -H 'Postman-Token: e4d0bf85-df6b-4ab0-bcca-1feb5a76f531' \
  -H 'cache-control: no-cache'  

```

获取urlpath参数

``` go
id:=c.Param("id")
```

##  从paramters中获取

请求：

```shell
curl -X GET \
  http://localhost:8080/v1/order?firstname=Jack&lastname=ma \
  -H 'Postman-Token: e4d0bf85-df6b-4ab0-bcca-1feb5a76f531' \
  -H 'cache-control: no-cache'  
```

获取paramters参数

```go
firstName := c.DefaultQuery("firstname", "Guest")  //如果读取不到firstname会给赋值一个默认的值 Guest
lastname := c.Query("lastname")// 读取 lastname的值，如果读取不到返回空字符串
// c.Query("lastname") 等价于 c.Request.URL.Query().Get("lastname")
```

##从Multipart/Urlencoded Form中获取



请求:

```shell
curl -X POST \
  'http://localhost:8080/v1/order/1/10010?pkg=test.1' \
  -H 'Postman-Token: 49319e4b-fd9a-416e-a69b-5ffb54b12413' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F message=msg
```

获取Multipart/Urlencoded Form参数

```go
message := c.PostForm("message") // 读取 message的值，如果读取不到返回空字符串
nick := c.DefaultPostForm("nick", "anonymous") // 读取nick的值，如果读取不到会给nick赋默认值anonymous
```

## 从paramter/form 读取map或者array

请求:

```shell
curl -X PUT \
  'http://localhost:8080/v1/order/1?idMap[a]=aaa&idMap[b]=bbb&idArr=c1,c2' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Postman-Token: a2f3a545-3b19-4a5d-a542-8b83c19ee2f9' \
  -H 'cache-control: no-cache' \
  -d 'nameArr=name1&nameArr=name2&nameMap%5Ba%5D=namea&nameMap%5Bb%5D=nameb&undefined='
```



从paramter获取map,array,从form获取map array

```go
idMap := c.QueryMap("idMap") //从paramter获取map
idArr := c.QueryArray("idArr")//从paramter获取array
nameArr := c.PostFormArray("nameArr") // 从form获取map
nameMap := c.PostFormMap("nameMap")// 从form获取array

```

