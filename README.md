# goboot
Inspired by springboot framework, making it easy to create web application with go.


## design
- keep simple, easy to use
- keep high performance（try to not use reflection）
- trace every http request and response from controller to service and to dao by trace id
- support log files classified to default.log and error.log
- support log desensitization in global http request and response
