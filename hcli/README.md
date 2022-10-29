<a name="u4qNM"></a>
# 前置
- golang 发送http/https最佳实践实例，基于golang自带的 **net/http** 库
- 支持： 
   - **复用** http/https 客户端。
   - 超时控制。
   - 支持小文件上传。
<a name="oJxUo"></a>
# 设计思路

- 设计思路：整体工具库设计，遵循了net/http的使用流程：
   - 获取Client。
   - 生成Request
   - 执行请求Do
   - 获取结果Response
- hcli 工具库所做的改变：
   - **client**:hcli 工具库，提供了线程安全的 **client_pool**，同时提供了方便的TLS客户端生成接口。
   - **Request**: hcli 工具库，封装了超时，自定义Header, 文件上传。
   - **Response**: hcli 工具库，对Response 提供了结构化和非结构化的body获取接口，提供了header获取接口，提供了StatusCode获取接口。
- 三步走，即可完成一个http/https 请求的发送：
   - 从池子中获取 hcli 的 http client。
   - 构造出hcli 的 Resquest。
   - 根据实际需要，按需解析Response。
<a name="BGcvg"></a>
# 典型使用场景
<a name="fn3Au"></a>
## 发送一个http请求
<a name="Cgex2"></a>
## 发送一个带有超时的请求
<a name="QcLuF"></a>
## 解析一个Response
<a name="Fvw44"></a>
## 发送一个https请求
<a name="TXqzj"></a>
## 上传一个文件

