# Luffy

* [前置](#前置)
* [daemon](#daemon)
* [graceful-HTTP](#graceful-HTTP)
* [task-queue](#task-queue)

## 前置

* __luffy__ 项目旨在总结一些项目中的 __最佳实践__，或整合一些实用的库包/工具。提升后端开发效率，提高开发质量。
* 代码主目录下，每个文件夹代表一个 __最佳实践__ 或 __工具库__，每个文件夹里面的README包含了具体说明。

## daemon

* 在golang服务中实现后台启动，集成 __命令行__ 实现服务安全退出和配置热重载示例。
* [deamon](https://github.com/YuleiGong/luffy/tree/main/daemon "daemon")

## graceful-HTTP
* 优雅退出HTTP服务。
* [graceful-HTTP](https://github.com/YuleiGong/luffy/tree/main/graceful-HTTP "优雅退出http服务")

## task-queue

* golang 编码的类似 python __Celery__ 单节点异步任务队列。
* 支持：
    * 任务重试
    * 设置任务超时
    * 任务执行状态查询

