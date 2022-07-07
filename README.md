### golang简单打包测试项目，支持配置文件config.yaml
### 手动调试
go mod init helloserver  
go mod tidy -compat=1.17    #自动安装依赖包  
go run helloserver.go  

编译打包测试版本修改 helloserver.go 文件下面内容  
str := "Hello world ! friend01  

访问请求： curl 127.0.0.1:18088/hello
