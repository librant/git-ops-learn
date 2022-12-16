
#### 从零开始构建镜像

1、编写源码，编译生成二进制   
2、编写 Dockerfile 文件   
3、根据 Dockerfile 构建镜像   
```shell
docker build . -t httpserver:0.0.1
```
4、镜像测试   
```shell
docker run -d -p 8080:8080 httpserver:0.0.1
```
5、推送到远程镜像仓库    
```shell
// 默认推送到 dockerhub 
docker push xxx/httpserver:0.0.1
```