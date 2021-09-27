# Filer各模块功能点
## 目录

- [pool 矿池模块](#pool)
    - [<font color=orange> 联合挖矿</font>](#/pool/main)
        - [FIL挖矿记录&订单](#/pool/userFilData)
            1. [收益记录](#/pool/profitLogList)
            2. [当前订单](#/pool/orderList)
            3. [历史订单](#/pool/orderList)
        
    - [<font color=orange> 矿池数据</font>](#/pool/htFilNodeData)
        - [矿池数据](#/pool/htFilNodeData)
      
    - [<font color=orange> 优选产品</font>](#/pool/productList)
        - [产品列表](#/pool/productList)
            1. [优选产品详情](#/pool/productOne)
            2. [立即挖矿](#/pool/productOne)
            3. [充值](#/pool/profitDrawing)
    

- [user 用户模块-个人中心](#user)
    - [<font color=orange> 登录</font>](#/account/login)
    - [<font color=orange> 注册</font>](#/account/register)
    - [<font color=orange> 账号与安全</font>](#/account/safe)
        - [绑定手机/邮箱](#/account/binding)
        - [修改登录密码](#/account/passwordUpdate)
        - [支付密码](#/account/payPasswordSet)
    - [<font color=orange> 身份认证</font>](#/account/kyc)


## [**APP**] pool 矿池模块

### <矿池> 页面（点击矿池页签）

#### <span id="/pool/main">联合挖矿</span>

<span id="/pool/userFilData"></span>
`post AUTH /pool/userFilData` 矿池用户数据 - app/联合挖矿

- <span id="/pool/userFilData">FIL挖矿记录&订单</span>

    <span id="/pool/userFilData"></span>
    `post AUTH /pool/userFilData` 矿池用户数据 - app/联合挖矿

1. 收益记录

    <span id="/pool/profitLogList"></span>
    `post AUTH /pool/profitLogList` 收益记录 - app/FIL 挖矿
   

2. 当前订单

   <span id="/pool/orderList"></span>
   `post AUTH /pool/orderList` 订单记录 - app/FIL 挖矿
   

3. 历史订单

   <span id="/pool/orderList"></span>
   `post AUTH /pool/orderList` 订单记录 - app/FIL 挖矿

#### <span id="/pool/htFilNodeData">矿池数据</span>

<span id="/pool/htFilNodeData"></span>
`post /pool/htFilNodeData` 矿池数据 - app/联合挖矿

#### <span id="/pool/productList">优选产品</span>

<span id="/pool/productList"></span>
`post /pool/productList` 产品列表

- 优选产品详情（点击优选产品）

    <span id="/pool/productOne"></span>
    `post /pool/productOne` 矿池产品详情 - app/挖矿详情


- 立即挖矿（点击立即挖矿）

    <span id="/pool/productOne"></span>
    `post /pool/productOne` 矿池产品详情 - app/挖矿详情
  
  
- 充值

    <span id="/pool/profitDrawing"></span>
    `post AUTH /pool/profitDrawing` 提取收益 - app/充值到钱包

#### 集群算力
- [ ] TODO

## [**APP**] user 账户模块

### <span id="/account/userdata"><个人中心> 页面（点击我的页签</span>

#### <span id="/account/login">登录</span>
1. <span id="/account/signInCheck"></span>
   `post /account/signInCheck` 登录提交（第一步检查资格）
   
2. <span id="/account/signIn"></span>
   `post /account/signIn` 登录提交（最终提交）
   
3. <span id="/verify/sendCode"></span>
   `post /verify/sendCode` 发送验证码
   
4. <span id="/verify/signInKey"></span>
   `post /verify/signInKey` 检查用户名是否存在

#### <span id="/account/register">注册</span>

1. <span id="/account/register"></span>
   `post /account/register` 注册账号

2. <span id="/verify/sendCode"></span>
   `post /verify/sendCode` 发送验证码

3. <span id="/verify/signInKey"></span>
   `post /verify/signInKey` 检查用户名是否存在
   
#### <span id="/account/safe">账号与安全</span>

- 绑定手机/邮箱

    <span id="/account/binding"></span>
    `post AUTH /account/binding` 绑定手机号或邮箱
  

- 修改登录密码

    <span id="/account/passwordUpdate"></span>
    `post AUTH /account/passwordUpdate` 修改登录密码
  

- 支付密码

1. <span id="/account/payPasswordSet"></span>
   `post AUTH /account/payPasswordSet` 设置支付密码
       
2. <span id="/account/payPasswordUpdate"></span>
   `post AUTH /account/payPasswordUpdate` 修改支付密码

- 谷歌身份验证
- [ ] TODO

#### <span id="/account/kyc">身份认证</span>

<span id="/user/kycUser"></span>
`post AUTH /user/kycUser` 身份认证(KYC)

## [**APP**] wallet 钱包模块

### <钱包> 页面（主页面）

#### 首页
1. `post AUTH /wallet/index` 钱包首页数据
2. `post /market/price` 市场币价

#### 资金流水
1. `post AUTH /wallet/billList`  -APP 资金流水

#### 交易详情
1. `post AUTH /wallet/transactionOne`  单条详细的交易数据

#### 提现
1. `post AUTH /user/index` 查看用户绑定情况
2. `post /verify/sendCode` 发送验证码
3. `post AUTH /wallet/transactionFrom`  提现