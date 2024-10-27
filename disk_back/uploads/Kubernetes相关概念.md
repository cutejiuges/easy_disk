## Kubernetes组件
一个kubernetes集群主要是由**控制节点**、**工作节点(node)** 构成，每个节点上会安装不同的组件。
![k8s组件](./image/k8s组件.png "k8s组件")

**master**节点：**集群的控制平面，负责集群的决策**
> **ApiServer**：api网关，所有的操作都需要通过ApiServer 
> **Scheduler**：负责集群的资源调度，按照预定的调度策略将pod调度到对应的node节点上
> **ControllerManager**：负责维护集群状态，譬如程序部署、故障检测、自动扩展、滚动更新等
> **ETCD**：负责存储集群中资源对象的信息，k-v数据库

**Node**节点：**集群的数据平面，负责为容器提供运行环境**
>**Kubelet**：负责维护容器的生命周期，控制docker来创建、更新、销毁容器
>**KubeProxy**：负责提供集群内的服务发现和负载均衡
>**Docker**：负责进行具体节点上的各种操作


### 实例说明kebernetes的调度关系
如果需要部署一个MySQL服务，kebernetes各组件之间的调度关系如下：
1. 启动kubernetes，master和node节点都会将自己的信息存储到etcd中；
2. MySQL的安装请求首先会被发送到master节点的ApiServer组件；
3. ApiServer调用Scheduler，决定将MySQL安装到哪一个node节点上面。这时，Scheduler会从etcd中拿取节点信息，根据算法决策，将结果告知ApiServer；
4. ApiServer再调用ControllerManager调度node执行MySQL的安装；
5. Kubelet接收到指令，会通知Docker，由Docker来启动一个MySQL的pod。pod时kubernetes中的最小操作单元，容器必须运行在pod中；
6. 至此，MySQL服务就运行起来了，如果外部用户需要访问此MySQL服务，需要通过kubeProxy代理来访问pod，外界用户就能访问pod中的服务了。