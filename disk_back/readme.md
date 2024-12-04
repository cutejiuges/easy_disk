# disk_back服务
本仓库实现了网盘能力的服务端逻辑，拆分为主干的微服务，以及一些通用的工具和中间件等。
当`idl`发生变更后，要更新相关的文件生成代码，可以参考以下命令：

```shell
# 文件服务相关：
kitex -module github.com/cutejiuges/disk_back ../../idl/cloud_disk/back/file_server.thrift
# 用户服务相关:
kitex -module github.com/cutejiuges/disk_back ../../idl/cloud_disk/back/user_server.thrift
```

## file_server
`file_server`中实现了文件相关的主要操作，包括文件的上传、下载、查询、编辑等能力，计划支持文件秒传、断点续传、分块上传下载能力。

需要生成文件微服务的相关代码，可以参考：
```shell
cd micro_services/file_server
kitex -module github.com/cutejiuges/disk_back -service cutejiuge.disk.file -use github.com/cutejiuges/disk_back/kitex_gen/ ../../../../idl/cloud_disk/back/file_server.thrift
cd ../..
```

## user_server
`user_server`中实现了用户相关的主要操作，包括用户的注册、登陆验证、信息修改、帐号注销等能力

需要生成用户微服务的相关代码，可以参考:
```shell
cd micro_services/user_server
kitex -module github.com/cutejiuges/disk_back -service cutejiuge.disk.user -use github.com/cutejiuges/disk_back/kitex_gen/ ../../../../idl/cloud_disk/back/user_server.thrift
cd ../..
```