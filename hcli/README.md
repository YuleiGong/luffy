# hcli
* golang 发送http/https最佳实践实例，基于golang自带的 __net/http__ 库
* 支持：
    * __复用__ http/https 客户端。
    * 超时控制。
    * 支持小文件上传。
* 设计思路：整体工具库设计，遵循了net/http的使用流程，既获取Client，生成Request，执行请求Do，获取结果Response


