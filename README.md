# goboot
Inspired by springboot framework, making it easy to create web application with go.


## design
- keep simple, avoid over-design
- keep high performance（try to not use reflection）
- trace every http request and response from controller to service and to dao by trace id
- log files classify by default and error
