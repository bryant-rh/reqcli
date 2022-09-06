# reqcli
reqcli is a tool for detecting interface time-consuming
用于调试接口，打印调试日志、性能跟踪，甚至可以 dump 完整的请求和响应内容。

# Usage
```Bash
reqcli is a tool for detecting interface time-consuming

Usage:
  reqcli [flags]

Flags:
  -d, --data string      HTTP POST data
  -H, --header strings   Pass custom header(s) to server
  -h, --help             help for reqcli
  -X, --request string   指定method,支持 GET、POST、DELETE、PUT (default "GET")
```

# Demo
## GET
```Bash
reqcli "https://httpbin.org/get"
```
输出
```
2022/09/05 22:33:34.599859 DEBUG [req] HTTP/2 GET https://httpbin.org/get
:authority: httpbin.org
:method: GET
:path: /get
:scheme: https
user-agent: reqcli api client
accept-encoding: gzip

:status: 200
date: Mon, 05 Sep 2022 14:33:34 GMT
content-type: application/json
content-length: 270
server: gunicorn/19.9.0
access-control-allow-origin: *
access-control-allow-credentials: true

{
  "args": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Host": "httpbin.org", 
    "User-Agent": "reqcli api client", 
    "X-Amzn-Trace-Id": "Root=1-631608be-13b1f17534af809e0fd21565"
  }, 
  "origin": "116.169.14.4", 
  "url": "https://httpbin.org/get"
}


the request total time is 1.329588s, and costs 631.631958ms on tls handshake
-------------------------------
TotalTime         : 1.329588s
DNSLookupTime     : 106.520875ms
TCPConnectTime    : 281.74025ms
TLSHandshakeTime  : 631.631958ms
FirstResponseTime : 304.512875ms
ResponseTime      : 931.125µs
IsConnReused:     : false
RemoteAddr        : 52.87.105.151:443
```

## POST
```Bash
reqcli "https://httpbin.org/post" \
-X POST  \
-H 'Content-Type: application/json' \
-d '{"username": "imroc"}'
```
输出
```
2022/09/05 22:42:26.464906 DEBUG [req] HTTP/2 POST https://httpbin.org/post
:authority: httpbin.org
:method: POST
:path: /post
:scheme: https
content-type:  application/json
user-agent: reqcli api client
content-length: 20
accept-encoding: gzip

{"username":"imroc"}

:status: 200
date: Mon, 05 Sep 2022 14:42:26 GMT
content-type: application/json
content-length: 453
server: gunicorn/19.9.0
access-control-allow-origin: *
access-control-allow-credentials: true

{
  "args": {}, 
  "data": "{\"username\":\"imroc\"}", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Content-Length": "20", 
    "Content-Type": "application/json", 
    "Host": "httpbin.org", 
    "User-Agent": "reqcli api client", 
    "X-Amzn-Trace-Id": "Root=1-63160ad2-07305fc1377a02ca4bdc3969"
  }, 
  "json": {
    "username": "imroc"
  }, 
  "origin": "116.169.14.4", 
  "url": "https://httpbin.org/post"
}


the request total time is 1.530676875s, and costs 817.091041ms on tls handshake
-------------------------------
TotalTime         : 1.530676875s
DNSLookupTime     : 79.809959ms
TCPConnectTime    : 323.585083ms
TLSHandshakeTime  : 817.091041ms
FirstResponseTime : 303.903666ms
ResponseTime      : 1.46225ms
IsConnReused:     : false
RemoteAddr        : 52.87.105.151:443
```

# Trace Info
```
TotalTime:         is a duration that total request took end-to-end.
      
DNSLookupTime:     is a duration that transport took to perform

ConnectTime:       is a duration that took to obtain a successful connection.

TCPConnectTime:    is a duration that took to obtain the TCP connection.

TLSHandshakeTime:  is a duration that TLS handshake took place.

FirstResponseTime: is a duration that server took to respond first byte since.
// connection ready (after tls handshake if it's tls and not a reused connection)

ResponseTime:      is a duration since first response byte from server to

IsConnReused:      is whether this connection has been previously
// used for another HTTP request.

IsConnWasIdle:     is whether this connection was obtained from an

ConnIdleTime:      is a duration how long the connection was previously

RemoteAddr:        returns the remote network address.
```