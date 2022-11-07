# udp 服务端
* 支持接口形式的拆包和解包
* 基于单向数据流

# 功能模块
* IServer： 实现了udp conn 的和其他配置的封装，提供了拆包和解包接口，数据处理handler接口。
* IConn: 实现了udp conn的封装，并关联了IServer,调用数据的解包，封包，数据handler处理。
* IHandler: 实现了数据handler接口。
* IUnPack: 实现了数据拆包接口
