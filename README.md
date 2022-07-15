[In Chinese 中文版](README.zh_cn.md)

# goboot
Inspired by springboot framework, making it easy to create web application with go.


## Features
- simple, easy to use
- high performance（try to not use reflection）
- request log (in access_log.go) and response log (in response_log.go)
- trace every http request and response in controller, service and dao by trace id
- log files classified to default.log and error.log
- log desensitization

## Usage

### http response
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
- Use the value of 'traceId'（"n82gede29h"） in the rest api response, you can quickly find the root cause from the default.log or error.log 
- You just need to do: grep 'n82gede29h' default.log (errro.log)

### all tracing log 
```json
{"level":"info","ts":"2022-05-29T23:30:56.104+0800","caller":"middleware/access_log.go:43","msg":"AccessLog","Method":"POST","IP":"127.0.0.1","Path":"/v1/student/get_one","Header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Cache-Control":["no-cache"],"Connection":["keep-alive"],"Content-Length":["22"],"Content-Type":["application/json"],"Postman-Token":["ded26fff-52e5-4fca-8d62-534bbd8c2a21"],"User-Agent":["PostmanRuntime/7.29.0"]},"Query":"","UserAgent":"PostmanRuntime/7.29.0","Request":"{\n    \"studentId\": 1\n}","TraceID":"n82gede29h"}
{"level":"info","ts":"2022-05-29T23:30:56.114+0800","caller":"dao/student_dao.go:21","msg":"SqlInfoLog","elapsed":0.009854539,"rows":1,"sql":"SELECT * FROM `student` WHERE id =1 ","TraceID":"n82gede29h"}
{"level":"info","ts":"2022-05-29T23:30:56.114+0800","caller":"service/student_service.go:21","msg":"student=&{1 Ray 6 123 2022-05-29 14:13:01 +0800 CST 2022-05-29 14:13:01 +0800 CST}","TraceID":"n82gede29h"}
{"level":"debug","ts":"2022-05-29T23:30:56.114+0800","caller":"controller/student_controller.go:37","msg":"student=&{1 Ray 6 123 2022-05-29 14:13:01 +0800 CST 2022-05-29 14:13:01 +0800 CST}","TraceID":"n82gede29h"}
{"level":"info","ts":"2022-05-29T23:30:56.114+0800","caller":"middleware/response_log.go:51","msg":"ResponseLog","Status":200,"Path":"/v1/student/get_one","Response":"{\"code\":0,\"message\":\"Success\",\"data\":{\"Id\":1,\"Name\":\"Ray\",\"Age\":6,\"Address\":\"123\",\"Modified\":\"2022-05-29T14:13:01+08:00\",\"Created\":\"2022-05-29T14:13:01+08:00\"},\"details\":[],\"traceId\":\"n82gede29h\"}","Cost":0.010137309,"TraceID":"n82gede29h"}

```
- One request has one trace path: 
  access_log=>student_controller=>student_service=>student_dao=>response_log
- You can also see all sql log with the 'TraceID', which will help to trace the database operation
