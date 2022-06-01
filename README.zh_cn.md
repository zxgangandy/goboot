[In Chinese 中文版](README.zh_cn.md)

# goboot
受到springboot框架的启发, 让go语言创建web应用更容易.

## 功能
- 简单易用
- 高性能（尽量避免使用反射）
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
- 使用api返回值中的'traceId'（"n82gede29h"）, 你能快速在default.log或者是error.log文件中定位问题
- 你只需要执行：grep 'n82gede29h' default.log (errro.log)
