
### Docker 常用命令

#### 1、docker 基础命令
1) 启动 docker 
```shell
systemctl start docker
```
2) 停止 docker
```shell
systemctl stop docker
```
3) 重启 docker
```shell
systemctl restart docker
```
4) docker 开机自启动
```shell
systemctl enable docker
```
5) 查看 docker 状态
```shell
systemctl status docker
```
6) 查看 docker 版本号和信息
```shell
docker version
docker info
```
7) docker 帮助命令
```shell
docker --help
```

#### 2、docker 镜像命令
1) 查看镜像列表
```shell
docker images
docker image ls
```
2) 搜索镜像
```shell
docker search image-name
// 搜索 STARS > 9000的 nginx 镜像
docker search --filter=STARS=9000 nginx
```
3) 拉取镜像
```shell
docker pull image-name
docker pull image-name:tag
```
4) 运行镜像
```shell
docker run image-name
docker run image-name:tag
```
5) 删除镜像
```shell
// 删除单个镜像
docker rmi -f image-name/imageID
// 删除全部镜像
docker rmi -f $(docker images -aq)
```
6) 保存镜像
```shell
docker save image-name/imageID -o image_path/image-name
```
7) 加载镜像
```shell
docker load -i image_path
```
8) 镜像标签
```shell
docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
```

#### 3、docker 容器命令
1) 查看运行的容器列表
```shell
docker ps -a
docker container ls
```
2) 运行容器
```shell
docker run -it -d --name container_name image_name:tag /bin/bash
```
3) 进入容器
```shell
// 进入容器
docker attach container_name/containerID
// 退出容器
exit 
ctrl + p + q
```
4) 停止容器
```shell
docker stop containerID/container_name
```
5) 重启容器
```shell
docker restart containerID/container_name
```
6) 启动容器
```shell
docker start containerID/container_name
```
7) kill 容器
```shell
docker kill containerID/container_name
```
8) 容器文件拷贝
```shell
// 从容器内 拷出
docker cp containerID/名称: 容器内路径  容器外路径
// 从外部 拷贝文件到容器内
docker cp 容器外路径 containerID/名称: 容器内路径
```
9) 查看容器日志
```shell
docker logs -f --tail=要查看末尾多少行 默认all containerID
```

#### 4、docker 运维命令
1) 查看 docker 工作目录
```shell
docker info | grep "Docker Root Dir"
```
2) 查看 Docker 磁盘占用情况
```shell
du -hs /var/lib/docker/ 
```
3) 查看 Docker 磁盘使用情况
```shell
docker system df
```
4) 删除无用的容器和镜像
```shell
// 删除异常停止的容器
docker rm `docker ps -a | grep Exited | awk '{print $1}'`
// 删除名称或标签为none的镜像
docker rmi -f  `docker images | grep '<none>' | awk '{print $3}'`
// 清除所有无容器使用的镜像
docker system prune -a
```
5) 查找大文件
```shell
// 查找 / 目录下的超过 100M 的文件
find / -type f -size +100M -print0 | xargs -0 du -h | sort -nr
// 查找指定docker使用目录下大于指定大小文件
// 查找 / 目录下的超过 100M 的文件
find / -type f -size +100M -print0 | xargs -0 du -h | sort -nr | grep '/var/lib/docker/overlay2/*'
```

---
**常见错误：**
1) no space left on device
```shell
docker: write /var/lib/docker/tmp/GetImageBlob325372670: no space left on device
```
一般这种错误是宿主机磁盘满导致的，需要进行磁盘清理

---
**参考文档：**  
1) https://blog.csdn.net/leilei1366615/article/details/106267225