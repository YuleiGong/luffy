# Luffy
- [前置](#%E5%89%8D%E7%BD%AE)
- [最佳实践：daemon](#daemon)
- [最佳实践：graceful-HTTP](#graceful-HTTP)
- [工具：hcli](#hcli)
- [工具：glog](#glog)
## 前置

- **luffy** 项目旨在总结一些项目中的 **最佳实践**，或封装一些项目中实用的库包/工具。提升后端开发效率，提高开发质量。
- 代码主目录下，每个文件夹代表一个 **最佳实践** 或 **工具库**，每个文件夹里面的README包含了该项目或工具的具体说明。
## daemon

- 在golang服务中实现后台启动，集成 **命令行** 实现服务安全退出和配置热重载示例。
- [deamon](https://github.com/YuleiGong/luffy/tree/main/daemon)
## graceful-HTTP

- 优雅退出HTTP服务。
- [graceful-HTTP](https://github.com/YuleiGong/luffy/tree/main/graceful-HTTP)
## glog

- 基于 **logrus** 封装的日志库。可以在项目中直接使用
- [glog](https://github.com/YuleiGong/luffy/tree/main/glog)
## hcli

- golang 发送http/https 请求工具库。
- [hcli](https://github.com/YuleiGong/luffy/tree/main/hcli)
