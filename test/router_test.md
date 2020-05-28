# router 测试
## 测试paramters

访问url path ，
``` shell script
request:

curl -X GET \
  'http://localhost:8080/v1/user/1?lastname=syc' \
  -H 'Postman-Token: 84e072b9-dafb-420f-acbe-15330f94f62e' \
  -H 'cache-control: no-cache'

response:
{
    "message": {
        "Id": "1",
        "FirstName": "Guest",
        "LastName": "syc"
    }
}
  
```

## 测试multipart/form
``` shell script
request:
curl -X POST \
  http://localhost:8080/v1/order/1/10010 \
  -H 'Postman-Token: 28a4554c-bf02-4b9a-bc6b-07b412b4eea8' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F message=msg
  
  
response:

{
    "fullPath": "/v1/order/:id/:itemId",
    "id": "1",
    "itemId": "10010",
    "message": "msg",
    "nick": "anonymous",
    "status": "posted"
}
```
