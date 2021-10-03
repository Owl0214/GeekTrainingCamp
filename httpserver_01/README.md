# 作业一

1.接收客户端 request，并将 request 中带的 header 写入 response header

2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header

3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出

4.当访问 localhost/healthz 时，应返回200

## 启动服务器

启动命令：
* -p: 指定端口，不指定则默认8080
* --log_dir：指定日志存放路径
> go run httpserver_01/main.go -p 8888 --log_dir=httpserver_01/logs

## 服务清单
访问根路径，返回所有path路径

## 作业说明
* 作业1：请求`/filter/getHeader`，如请求访问'http://localhost:8888/filter/getHeader'
* 作业2：
    * 请求`/filter/getSystemGoVersion`,获取Go的Version
    * 请求`/filter/getSystemEnv?sysConfigName=`，获取指定系统配置项名
* 作业3：启动时指定了--log_dir，日志查看响应路径下文件内容
* 作业4：请求`http://localhost:8888/healthz`,返回`"{Code:200,Msg:'Serving'}"`