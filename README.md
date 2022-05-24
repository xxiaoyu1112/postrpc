### Post RPC


#### 总体目录结构
```text
idl: proto3的接口定义文件,分别应该定义四个模块
post_*: rpc各个模块服务Repo
sdk: rpc服务sdk的生成
sh: 更新、生成rpc服务和sdk的脚本
```
#### 依赖

> go 1.17.x \
> grpc \
> proto3 \
> grpcui

#### 运行RPC项目

编译
```bash
cd post_model_manage
sh build.sh
```
生成测试平台
```bash
grpcui -plaintext [host:port]
# 例如
grpcui -plaintext localhost:50051
```
#### 端口

目前未加入consul + envoy,没有service mesh功能
post_data_collect   port:   50051
post_data_manage    port:   50052
post_model_manage   port:   50053
post_model_predict  port:   50054

#### 一次操作四个Repo
##### 编译

```bash
# 方法一
./build_all.sh
# 方法二
bash build_all.sh
# 错误方法
sh build_all.sh
```
##### 运行

```bash
# 方法一
./bootstrap.sh
# 方法二
bash bootstrap.sh
# 错误方法
sh bootstrap.sh
```