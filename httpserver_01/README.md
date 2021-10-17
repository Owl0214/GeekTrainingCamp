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
* 作业3：启动时指定了--log_dir，日志查看相应路径下文件内容
* 作业4：请求`http://localhost:8888/healthz`,返回`"{Code:200,Msg:'Serving'}"`



# 作业二
1. 编写Dockerfile将httpserver容器化
2. 将Dockerfile推送到官方镜像仓库
3. 通过Docker命令启动httpserver容器

## 作业解答说明
1. Dockerfile使用了golang:alpine，比较轻量级，通过makefile在linux中构建可执行文件。
构建方法：在linux系统中（CentOS7），通过mount，将windows的工程目录挂载到linux的文件夹中。【需先将windows相应的文件夹设置为共享，同时关闭防火墙】
>mount -t cifs -o username={windowsUserName},password={windowsPwd} //{windowsServerIp}/GeekTrainingCamp /home/windows/
2. 在centos中的对应文件夹下，执行make build进行编译
3. 执行make release，进行镜像构建
4. 启动容器
>docker run --name gohttpserver1 -itd -p 9010:8080  geektrainingcamp/httpserver:v1`
5. 最后，访问服务
> curl localhost:9010

返回服务响应结果
["/filter/getSystemGoVersion","/filter/getSystemEnv?sysConfigName= ","/","/healthz","/filter/getHeader"]

6. 镜像推送
> docker login myHarborServerIp:5000
> docker push geektrainingcamp/httpserver:v1