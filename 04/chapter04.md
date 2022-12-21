### 第四讲：GitOps 实现应用秒级自动发布和回滚

#### 1、传统 K8s 应用发布流程

1) 更新应用方式
- 使用 kubectl set image 命令
```shell
kubectl set image deployment/httpserver-test librant/httpserver:v0.0.2
```
- 修改本地的 Manifest
```shell
kubectl apply -f httpserver-deploy-v2.yaml
```
- 修改集群内 Manifest
```shell
kubectl edit deployment httpserver-test
```

#### 2、从零开始搭建 GitOps 发布工作流

1) 安装 FluxCD 并创建工作流
- 负责监听 Git 仓库变化，自动部署的工具

---
参考文档：
- https://blog.csdn.net/tao12345666333/article/details/121112980
- https://fluxcd.io/flux/installation/
- https://github.com/kubevela/kubevela
- http://kubevela.net/zh/docs/reference/addons/fluxcd
- https://www.jianshu.com/p/84256ae4abc5

