[In Chinese 中文版](README.zh_cn.md)

# goboot
受到springboot框架的启发, 让go语言创建web应用更容易.

## 功能
- 简单易用
- 高性能（尽量避免使用反射）
- 请求日志 (详见access_log.go) 和请求返回日志 (详见response_log.go)
- 通过trace id在controller, service 和dao层跟踪每个请求和相应
- 日志分级保存到default.log 文件和error.log文件
- 日志脱敏

## 使用

### http返回值
```json
{
    "code":0,
    "message":"Success",
    "data":{
        "Id":1,"Name":"Ray","Age":6,"Address":"123","Modified":"2022-05-29T14:13:01+08:00","Created":"2022-05-29T14:13:01+08:00"
    },
    "details":[],
    "traceId":"n82gede29h"
}
```
- 在每个用户的http返回值中，都返回了traceId
- 使用api返回值中的'traceId'（"n82gede29h"）, 你能快速在default.log或者是error.log文件中定位问题
- 你只需要执行：grep 'n82gede29h' default.log (errro.log)

### 全链路日志
```json
{"level":"info","ts":"2022-05-29T23:30:56.104+0800","caller":"middleware/access_log.go:43","msg":"AccessLog","Method":"POST","IP":"127.0.0.1","Path":"/v1/student/get_one","Header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Cache-Control":["no-cache"],"Connection":["keep-alive"],"Content-Length":["22"],"Content-Type":["application/json"],"Postman-Token":["ded26fff-52e5-4fca-8d62-534bbd8c2a21"],"User-Agent":["PostmanRuntime/7.29.0"]},"Query":"","UserAgent":"PostmanRuntime/7.29.0","Request":"{\n    \"studentId\": 1\n}","TraceID":"n82gede29h"}
{"level":"info","ts":"2022-05-29T23:30:56.114+0800","caller":"dao/student_dao.go:21","msg":"SqlInfoLog","elapsed":0.009854539,"rows":1,"sql":"SELECT * FROM `student` WHERE id =1 ","TraceID":"n82gede29h"}
{"level":"info","ts":"2022-05-29T23:30:56.114+0800","caller":"service/student_service.go:21","msg":"student=&{1 Ray 6 123 2022-05-29 14:13:01 +0800 CST 2022-05-29 14:13:01 +0800 CST}","TraceID":"n82gede29h"}
{"level":"debug","ts":"2022-05-29T23:30:56.114+0800","caller":"controller/student_controller.go:37","msg":"student=&{1 Ray 6 123 2022-05-29 14:13:01 +0800 CST 2022-05-29 14:13:01 +0800 CST}","TraceID":"n82gede29h"}
{"level":"info","ts":"2022-05-29T23:30:56.114+0800","caller":"middleware/response_log.go:51","msg":"ResponseLog","Status":200,"Path":"/v1/student/get_one","Response":"{\"code\":0,\"message\":\"Success\",\"data\":{\"Id\":1,\"Name\":\"Ray\",\"Age\":6,\"Address\":\"123\",\"Modified\":\"2022-05-29T14:13:01+08:00\",\"Created\":\"2022-05-29T14:13:01+08:00\"},\"details\":[],\"traceId\":\"n82gede29h\"}","Cost":0.010137309,"TraceID":"n82gede29h"}
```
- 每一个请求都有一个trace链路:
  access_log=>student_controller=>student_service=>student_dao=>response_log
- 所有的sql都用'TraceID'关联起来, 这样有助于跟踪每个请求的sql执行
