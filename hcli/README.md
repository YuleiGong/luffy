# 前置
- golang 发送http/https最佳实践实例，基于golang自带的 **net/http** 库
- 支持： 
   - **复用** http/https 客户端。
   - 超时控制。
   - 支持小文件上传。
# 设计思路

- 功能模块：
   - **IClientPool**：封装了http client连接池相关接口。
   - **IRequest**: 封装了发送请求的接口，目前只有Do接口。
   - **IResponse**: 封装了解析Response的相关接口。

![](https://cdn.nlark.com/yuque/0/2022/jpeg/2172986/1667227641258-9755d472-e088-4621-ab75-3251281c1732.jpeg)

- 设计思路：整体工具库设计，遵循了net/http的使用流程：
   - 通过ClientPool获取Client。
   - 生成Request。
   - 执行请求Do。
   - 获取结果Response。
- hcli 工具库所做的改变：
   - **client**:hcli 工具库，提供了线程安全的 **client_pool**，同时提供了TLS客户端。
   - **Request**: hcli 工具库，封装了超时，自定义Header, 文件上传。
   - **Response**: hcli 工具库，对Response 提供了结构化和非结构化的body获取接口，提供了header获取接口，提供了StatusCode获取接口。
- 三步走，即可完成一个http/https 请求的发送：
   - 从池子中获取 hcli 的 http client。
   - 构造出hcli 的 Resquest。
   - 根据实际需要，按需解析Response。
# 典型使用场景

- 执行测试案例前，先启动http测试服务
```go
cd example/server
go run server.go
```

- 执行HTTPS测试案例前，先启动https 测试服务
```go
cd example/https_server
go run server.go
```
## 发送一个http请求

- 测试案例详细见 **example/http_test.go**
```go
//获取一个http client pool
var cliPool hcli.IClientPool = hcli.GetClientPool()

func TestSendHTTP(t *testing.T) {
    var (
        r    hcli.IResponse
        err  error
        body []byte
    )
    if r, err = sendHTTP(); err != nil {
        t.Fatal(err)
    }
    defer f.Close()
    t.Logf("%d", r.GetStatusCode())
}

func sendHTTP() (r hcli.IResponse, err error) {
    var (
        method = "GET"
        url    = "http://127.0.0.1:8080/hello"
    )
    client := cliPool.GetOrCreateClient("http") //获取http client

    var resp *http.Response //发送一个请求
    resp, err = hcli.NewRequest(method, url).Do(hcli.NULL_BODY, client)

    return hcli.NewResponse(resp), err //初始化一个Response对象，便于解析结果
}
```
## 发送一个带有超时退出请求

- 测试案例详见 **example/http_timeout_test.go**
```go
var cliPool hcli.IClientPool = hcli.GetClientPool()

type Resp struct {
	Result string `json:"result"`
}

func TestSendHTTP(t *testing.T) {
	if _, err := sendTimeoutHTTP(); err != nil {
		if hcli.IsTimeout(err) {
			t.Logf("time out")
		} else {
			t.Fatal(err)
		}
	}
}

func sendTimeoutHTTP() (r hcli.IResponse, err error) {
	var (
		method = "GET"
		url    = "http://127.0.0.1:8080/timeout"
	)
	client := cliPool.GetOrCreateClient("http") //获取http client

	var resp *http.Response //发送一个请求
	resp, err = hcli.NewRequest(method, url, hcli.WithTimeout(1*time.Second)).Do(hcli.NULL_BODY, client)

	return hcli.NewResponse(resp), err //初始化一个Response对象，便于解析结果
}
```
## 解析一个Response

- 测试案例详见 **example/http_parse_test.go**
```go
//获取一个http client pool
var cliPool hcli.IClientPool = hcli.GetClientPool()

type Resp struct {
	Result string `json:"result"`
}

func TestSendHTTP(t *testing.T) {
	var (
		r   hcli.IResponse
		err error
	)
	if r, err = sendHTTP(); err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	resp := Resp{}
	if err = r.GetStructBody(&resp); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)

}

func sendHTTP() (r hcli.IResponse, err error) {
	var (
		method = "GET"
		url    = "http://127.0.0.1:8080/hello"
	)
	client := cliPool.GetOrCreateClient("http") //获取http client

	var resp *http.Response //发发送一个请求
	resp, err = hcli.NewRequest(method, url).Do(hcli.NULL_BODY, client)

	return hcli.NewResponse(resp), err //初始化一个Response对象，便于解析结果
}
```
## 发送一个https请求

- 客户端携带CA证书访问HTTPS服务端。测试案例详见 **example/https_test.go。**
- 测试证书存放在example/cert目录中。
```go
var cliPool hcli.IClientPool = hcli.GetClientPool()

type Resp struct {
	Result string `json:"result"`
}

func TestSendHTTPS(t *testing.T) {
	var (
		r   hcli.IResponse
		err error
	)
	if r, err = sendHTTPS(); err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	t.Logf("%d", r.GetStatusCode())
}

func sendHTTPS() (r hcli.IResponse, err error) {
	var (
		method = "GET"
		url    = "https://localhost:8080/hello"
	)

	var client *http.Client//创建https client
	if client, err = cliPool.GetOrCreateClientWithTLS("https", "cert/ca.crt"); err != nil {
		return
	}

	var resp *http.Response //发发送一个请求
	resp, err = hcli.NewRequest(method, url).Do(hcli.NULL_BODY, client)

	return hcli.NewResponse(resp), err //初始化一个Response对象，便于解析结果
}
```
## 上传一个文件

- 支持上传一个小文件到服务端。测试案例详见 **example/http_upload_test.go**
```go
var cliPool hcli.IClientPool = hcli.GetClientPool()

type Resp struct {
	Result string `json:"result"`
}

func TestUpload(t *testing.T) {
	var (
		r    hcli.IResponse
		err  error
		body []byte
	)
	if r, err = sendHTTP(); err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	if body, err = r.GetBody(); err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", string(body))

}

func sendHTTP() (r hcli.IResponse, err error) {
	var (
		method = "GET"
		url    = "http://127.0.0.1:8080/upload"
	)

	client := cliPool.GetOrCreateClient("upload")
	file := "upload.txt"
	field := "upload"

	var resp *http.Response //发发送一个请求
	resp, err = hcli.NewRequest(method, url, hcli.WithUploadFile(file, field)).Do(hcli.NULL_BODY, client)

	return hcli.NewResponse(resp), err //初始化一个Response对象，便于解析结果
}
```

