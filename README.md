##### influen verse
#### 引入私有仓库

> 参考教程：https://www.jianshu.com/p/50884e7bcf97
> go get github.com/Super-NFT/warehouses

#### 架构目录架构

```
api 业务项目代码
api/game nft 商场
api/game/controller 控制器
api/game/middleware 中间件
api/game/service 业务逻辑层，核心代码都写到这里
api/game/router 路由
api/game/server 启动
api/game/common 项目公用
api/game/config 项目配置
api/game/form 输入和输出结构体 req|resp
api/game/output 输入和输出逻辑处理
```

#### 层级直接调用

```
common: 禁止调用其他层
controller 可以调用service、form、put、common。禁止调用model
form: 可以调用common
middleware: 不可调用其他层
model: 可调用common、form
put: 可form、common
router: 可controller
server: 禁止其他层调用
service: 可调用model, 子包内部相互调用
test: 禁止其他层调用
```
''

### 一些参考的文章

#内置合约相关的，代付相关的合约，可以查看下
https://juejin.cn/post/6876330619798814728

#测试网上的代付，当前是可以正常的运行了
#正式的代付，需要注意，是否需要申请 //这个要问下conflux相关的工作人员
//关键字  下面链接下的 代付规则 查看下相关说明
https://forum.conflux.fun/t/conflux/11949

#conflux的开发文档
https://developer.confluxnetwork.org/conflux-doc/docs/json_rpc/

#这个是相关的在线ide
https://chainide.com/s/conflux/2b64fb17acff4e57849a278a1496fe3c?open=ERC721AA.sol&type=file


#相关的一些可以参考的conflux的开发的案例
https://github.com/conflux-fans/go-conflux-sdk-examples

https://github.com/conflux-fans/conflux-contracts

#相关的介绍文档
https://forum.conflux.fun/t/conflux-2022-5-18-721-20-721-1155-nft/8781