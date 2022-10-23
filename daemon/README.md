# daemon

* 服务的daemon模式下启动，同时集成命令行的安全退出和配置重载。
* 单体应用中，这样的服务启动模式会很好用。
* 例子详细见:[sample.go](https://github.com/YuleiGong/luffy/blob/main/daemon/example/sample.go)
* 使用：
    ```go
    cd example
    go build -o server
    ./server
    ./server -s reload
    ./server -s quit
    ./server -s stop
    ```

