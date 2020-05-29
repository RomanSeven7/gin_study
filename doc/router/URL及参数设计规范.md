# **URL及参数设计规范**

##1.uri设计规范

1) uri末尾不需要出现斜杠/
2) 在uri中使用斜杠/是表达层级关系的。
3) 在uri中可以使用连接符-, 来提升可读性。
比如 http://xxx.com/xx-yy 比 http://xxx.com/xx_yy中的可读性更好。
4) 在uri中不允许出现下划线字符_.
5) 在uri中尽量使用小写字符。
6) 在uri中不允许出现文件扩展名. 比如接口为 /xxx/api, 不要写成 /xxx/api.php 这样的是不合法的。
7) 在uri中使用复数形式。

具体可以看：（https://blog.restcase.com/7-rules-for-rest-api-uri-design/）

在RESTful架构中，每个uri代表一种资源，因此uri设计中不能使用动词，只能使用名词，并且名词中也应该尽量使用复数形式。使用者应该使用相应的http动词 GET、POST、PUT、PATCH、DELETE等操作这些资源即可。

那么在我们未使用RESTful规范之前，我们是如下方式来定义接口的，形式是不固定的，并且没有统一的规范。比如如下形式:

```html
http://xxx.com/api/getallUsers; // GET请求方式，获取所有的用户信息
http://xxx.com/api/getuser/1;   // GET请求方式，获取标识为1的用户信息
http://xxx.com/api/user/delete/1 // GET、POST 删除标识为1的用户信息
http://xxx.com/api/updateUser/1  // POST请求方式 更新标识为1的用户信息
http://xxx.com/api/User/add      // POST请求方式，添加新的用户
```

如上我们可以看到，在未使用Restful规范之前，接口形式是不固定的，没有统一的规范，下面我们来看下使用RESTful规范的接口如下，两者之间对比下就可以看到各自的优点了。

```
http://xxx.com/api/users;     // GET请求方式 获取所有用户信息
http://xxx.com/api/users/1;   // GET请求方式 获取标识为1的用户信息
http://xxx.com/api/users/1;   // DELETE请求方式 删除标识为1的用户信息
http://xxx.com/api/users/1;   // PATCH请求方式，更新标识为1的用户部分信息
http://xxx.com/api/users;     // POST请求方式 添加新的用户
```

## 2.HTTP请求规范

GET (SELECT): 查询；从服务器取出资源.
POST(CREATE): 新增; 在服务器上新建一个资源。
PUT(UPDATE): 更新; 在服务器上更新资源(客户端提供改变后的完整资源)。
PATCH(UPDATE): 更新；在服务器上更新部分资源(客户端提供改变的属性)。
DELETE(DELETE): 删除; 从服务器上删除资源。

## 3.参数命名规范

参数推荐采用下划线命名的方式。比如如下demo:

```
http://xxx.com/api/today_login // 获取今天登录的用户。
http://xxx.com/api/today_login&sort=login_desc // 获取今天登录的用户、登录时间降序排序。
```

## 4.http状态码相关的

### 状态码范围

客户端的每一次请求, 服务器端必须给出回应，回应一般包括HTTP状态码和数据两部分。

1xx: 信息，请求收到了，继续处理。
2xx: 代表成功. 行为被成功地接收、理解及采纳。
3xx: 重定向。
4xx: 客户端错误，请求包含语法错误或请求无法实现。
5xx: 服务器端错误.

#### 2xx 状态码

200 OK [GET]: 服务器端成功返回用户请求的数据。
201 CREATED [POST/PUT/PATCH]: 用户新建或修改数据成功。
202 Accepted 表示一个请求已经进入后台排队(一般是异步任务)。
204 NO CONTENT -[DELETE]: 用户删除数据成功。

#### 4xx状态码

400：Bad Request - [POST/PUT/PATCH]: 用户发出的请求有错误，服务器不理解客户端的请求，未做任何处理。
401: Unauthorized; 表示用户没有权限(令牌、用户名、密码错误)。
403：Forbidden: 表示用户得到授权了，但是访问被禁止了, 也可以理解为不具有访问资源的权限。
404：Not Found: 所请求的资源不存在，或不可用。
405：Method Not Allowed: 用户已经通过了身份验证, 但是所用的HTTP方法不在它的权限之内。
406：Not Acceptable: 用户的请求的格式不可得(比如用户请求的是JSON格式，但是只有XML格式)。
410：Gone - [GET]: 用户请求的资源被转移或被删除。且不会再得到的。
415: Unsupported Media Type: 客户端要求的返回格式不支持，比如，API只能返回JSON格式，但是客户端要求返回XML格式。
422：Unprocessable Entity: 客户端上传的附件无法处理，导致请求失败。
429：Too Many Requests: 客户端的请求次数超过限额。

#### 5xx 状态码

5xx 状态码表示服务器端错误。

500：INTERNAL SERVER ERROR; 服务器发生错误。
502：网关错误。
503: Service Unavailable 服务器端当前无法处理请求。
504：网关超时。

##5.统一返回数据格式

RESTful规范中的请求应该返回统一的数据格式。对于返回的数据，一般会包含如下字段:

1) code: http响应的状态码。
2) status: 包含文本, 比如：'success'(成功), 'fail'(失败), 'error'(异常) HTTP状态响应码在500-599之间为 'fail'; 在400-499之间为 'error', 其他一般都为 'success'。 对于响应状态码为 1xx, 2xx, 3xx 这样的可以根据实际情况可要可不要。

当status的值为 'fail' 或 'error'时，需要添加 message 字段，用于显示错误信息。

3) data: 当请求成功的时候, 返回的数据信息。 但是当状态值为 'fail' 或 'error' 时，data仅仅包含错误原因或异常信息等。

返回成功的响应JSON格式一般为如下:

```json
{
    "code": 200,
    "status": "success",
    "data": [{
        "userName": "tugenhua",
        "age": 31
    }]
}
```

返回失败的响应json格式为如下:

```json
{
    "code": 401,
    "status": "error",
    "message": '用户没有权限',
    "data": null
}
```