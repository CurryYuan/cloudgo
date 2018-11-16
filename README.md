---
date: 2018-11-13 21:11
status: public
title: README.md
---

# cloudgo
## 框架选择 Martini 
Martini是一个强大为了编写模块化Web应用而生的GO语言框架.
### 功能列表
 - 使用极其简单.
 - 无侵入式的设计.
 - 很好的与其他的Go语言包协同使用.
 - 超赞的路径匹配和路由.
 - 模块化的设计 - 容易插入功能件，也容易将其拔出来.
 - 已有很多的中间件可以直接使用.
 - 框架内已拥有很好的开箱即用的功能支持.
 - 完全兼容*http.HandlerFunc*接口.

## 设计思路
 一个简单的go服务器应用，可以进行路由和返回静态网页
## 运行
>go run main.go -p 8080

## 测试
curl测试，-v可以显示客户端与服务端收发数据的详细信息
```
[root@centos_base etc]# curl -v http://localhost:8080
* About to connect() to localhost port 8080 (#0)
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET / HTTP/1.1
> User-Agent: curl/7.29.0
> Host: localhost:8080
> Accept: */*
> 
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Content-Length: 1609
< Content-Type: text/html; charset=utf-8
< Last-Modified: Tue, 13 Nov 2018 13:00:14 GMT
< Date: Tue, 13 Nov 2018 13:04:16 GMT
< 
<!DOCTYPE html>
<html>

<head>
    <link rel="stylesheet" type="text/css" href="index.css">
    <link rel="icon" href="http://courses.cs.washington.edu/courses/cse190m/09sp/homework/1/pie_icon.gif" type="image/x-icon">
    <meta charset="utf-8">
    <title>Welcome to pie</title>
</head>

<body>
    <h1>learn to make lemon meringue pie</h1>
    <div class="p1"><img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1505279867936&di=e8c49944ebd706dd5862f3dfd705e7a5&imgtype=jpg&src=http%3A%2F%2Fimg4.imgtn.bdimg.com%2Fit%2Fu%3D3402167733%2C2678319605%26fm%3D214%26gp%3D0.jpg"
            alt="pie1" width="500" height="400"><img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1505283313660&di=4d2842deb80a9eadf45ccadeaef540c8&imgtype=0&src=http%3A%2F%2Fim5.tongbu.com%2FArticleImage%2F3c0ae2d9-8.jpg%3Fw%3D480%2C343"
            alt="pie2" width="500" height="400"></div>
    <p class="p1"><a href="pie.html">Grandma's Lemon Meringue Pie</a></p>
    <div class="p1"><img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1505283467215&di=0361f0ac0d8fbe928e50d570632dd263&imgtype=0&src=http%3A%2F%2Fi5.xiachufang.com%2Fimage%2F600%2Fa8bce2fa6fd011e4865eb8ca3aeed2d7.jpg"
            alt="pie3" width="500" height="400"><img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1505283541313&di=aab3f2a7207ca1f1f0e09868bca11c40&imgtype=0&src=http%3A%2F%2Fmedia-cdn.tripadvisor.com%2Fmedia%2Fphoto-s%2F01%2Fd3%2Fd3%2F75%2Fbrennan-s-of-houston.jpg"
            alt="pie4" width="500" height="400"></div>
</body>

* Connection #0 to host localhost left intact
</html>
```
ab压力测试, -n 1000 -c 1000发送1000个请求，并发为100个
```
[root@centos_base etc]# ab -n 1000 -c 100 http://localhost:8080/
This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /
Document Length:        1609 bytes

Concurrency Level:      100     ###并发个数
Time taken for tests:   0.455 seconds   ###总请求时间
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      1795000 bytes
HTML transferred:       1609000 bytes
Requests per second:    2197.93 [#/sec] (mean)
Time per request:       45.497 [ms] (mean)
Time per request:       0.455 [ms] (mean, across all concurrent requests)
Transfer rate:          3852.82 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   3.0      1      17
Processing:     2   41  34.4     23     179
Waiting:        2   39  34.3     23     177
Total:          2   43  34.3     25     181

Percentage of the requests served within a certain time (ms)
  50%     25
  66%     53
  75%     77
  80%     80
  90%     98
  95%    108
  98%    114
  99%    120
 100%    181 (longest request)

```