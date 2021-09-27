### filer-app & controller's backend

## Interfaces

## 目录
- [表结构设计](StorageSturct.md)
- [Filer各模块功能点](./ModuleFactor.md)
- [wallet 钱包模块](#wallet)
    - `PAY`[<font color=#FF0000> /wallet/index 钱包首页数据</font>](#/wallet/index) `AUTH`
    - `PAY`[<font color=#FF0000> /wallet/transactionOne 单条详细的交易数据</font>](#/wallet/transactionOne) `AUTH`
    - `PAY`[<font color=#FF0000> /wallet/billList -APP 资金流水</font>](#/wallet/billList)`AUTH`
    - `PAY`[<font color=#FF0000> /wallet/transactionFrom 提现</font>](#/wallet/transactionFrom) `AUTH`

- [pool 矿池模块](#pool)
    - `FILER`[<font color=#FF0000> /pool/productList 产品列表</font>](#/pool/productList)
    - `FILER`[<font color=#FF0000> /pool/htFilNodeData 矿池数据 - app/联合挖矿</font>](#/pool/htFilNodeData)
    - `FILER`[<font color=#FF0000> /pool/userFilData 矿池用户数据 - app/联合挖矿</font>](#/pool/userFilData) `AUTH`
    - `FILER`[<font color=#FF0000> /pool/profitLogList 收益记录 - app/FIL 挖矿</font>](#/pool/profitLogList) `AUTH`
    - `FILER`[<font color=#FF0000> /pool/orderList 订单记录 - app/FIL 挖矿</font>](#/pool/orderList) `AUTH`
    - `FILER`[<font color=#FF0000> /pool/profitDrawing 提取收益 - app/充值到钱包</font>](#/pool/profitDrawing) `AUTH`
    - `FILER`[<font color=#FF0000> /pool/productOne 矿池产品详情 - app/挖矿详情</font>](#/pool/productOne)
    - `FILER`[<font color=#FF0000> /pool/purchaseFrom 购买矿池产品 - app/挖矿购买</font>](#/pool/purchaseFrom) `AUTH`

- [account 账户模块](#account)
    - `PAY`[<font color=#FF0000>/account/signInCheck 登录提交（第一步检查资格）</font>](#/account/signInCheck)
    - `PAY`[<font color=#FF0000>/account/signIn 登录提交（最终提交）</font>](#/account/signIn)
    - `PAY`[<font color=#FF0000>/account/register 注册账号 </font>](#/account/register)
    - `PAY`[/account/passwordReset 忘记密码](#/account/passwordReset)
    - `PAY`[/account/out 退出登录](#/account/out)`AUTH`
    - `PAY`[<font color=#FF0000>/account/binding 绑定手机号或邮箱  </font>](#/account/binding) `AUTH`

- [verify 验证模块](#verify)
    - `PAY`[<font color=#FF0000>/verify/sendCode 发送验证码  </font>](#/verify/sendCode)
    - `PAY`[<font color=#FF0000>/verify/signInKey 检查用户名是否存在  </font>](#/verify/signInKey)

- [market 市场模块](#market)
    - `FILER`[<font color=#FF0000>/market/price 市场币价  </font>](#/market/price)

- [user 用户模块](#user)
  - `PAY`[<font color=#FF0000>/user/index 用户首页数据  </font>](#/user/index)  `AUTH`

<span id="wallet"></span>

## [**APP**] wallet 钱包模块

<span id="/wallet/index"></span>
`post AUTH /wallet/index` 钱包首页数据

请求:nil

返回

|name|属性|例子|说明|
|---|----|----|----|
|wallet|arr|[ ]|钱包数组 - 固定两条 FIL 和 ETH|
|`wallet.currency`|fil|string|币种 fil,eth|
|`wallet.numAll`|string|1.001|全部余额|
|`wallet.numAvailable`|string|0.2102|可用余额|
|`wallet.address`|string|0x8e...|钱包地址|
|list|arr|[ ]|交易记录数组 - 最新前5条|
|`list.transactionId`|int|1|标识transaction id|
|`list.num`|string|1.2102|交易金额|
|`list.currency`|string|fil|币种 fil,eth|
|`list.types`|string|1|交易类型，充值 = 1,提现 = 2，转账 = 3|
|`list.time`|string|2021-07-16|交易日期|

------------
<span id="/wallet/transactionOne"></span>
`post AUTH /wallet/transactionOne` 单条详细的交易数据

types ：充值 = 1,提现 = 2，转账 = 3


请求

|name|属性|例子|说明|
|---|----|----|----|
|id|int|1|数据标识id|

返回 注：～提现，充值返回方式

|name|属性|例子|说明|
|---|----|----|----|
|types|string|1|交易类型 充值 = 1,提现 = 2，转账 = 3 |
|currency|string|fil|币种|
|serviceCharge|0.0001|string|手续费|
|fromAddress|string|0x8eE...|发送者-钱包地址|
|toAddress|string|0x852...|接收者-钱包地址|
|txId|string|0x96679...|交易hash|
|time|string|2021-02-24 12:00|交易时间|
|state|string|1|交易状态 成功=1，失败 = -1，审核中 = 2|

返回 注：～转账 返回方式

|name|属性|例子|说明|
|---|----|----|----|
|types|string|3|交易类型 充值 = 1,提现 = 2,转账 = 3|
|currency|string|fil|币种|
|fromName|string|13688888888|发送者-账号|
|toName|string|13666666666|接收者-账号|
|payId|string|0x96679...|流水号|
|time|string|2021-02-24 12:00|交易时间|
|state|string|1|交易状态 成功=1，失败 = -1，审核中 = 2|

------------
<span id="/wallet/billList"></span>
`post AUTH /wallet/billList`  -APP 资金流水

多条件查询

请求

|name|属性|例子|说明|
|---|----|----|----|
|currency|string|fil|币种 fil，eth，全部=all|
|types|string|all|交易类型 充值 = 1,提现 = 2 ,转账 = 3，全部=all|
|date|string|2019-01|日期 查询1月1日～1月30日 的数据|

返回

|name|属性|例子|说明|
|---|----|----|----|
|list|arr|[ ]|交易数据 数组|
|`list.transactionId`|int|1|数据标识id|
|`list.currency`|string|fil|币种 fil，eth，全部=all|
|`list.num`|1.0001|string|交易金额|
|`list.types`|string|1|交易类型 充值 = 1， 提现 = 2 ，转账 = 3|
|`list.balanceRecord`|string|1000.0213|余额变动记录|
|`list.time`|string|2021.03.01 20:32:52|交易时间|

------------
<span id="/wallet/transactionFrom"></span>
`post AUTH /wallet/transactionFrom` 提现

钱包余额提现，到站外erc20地址

请求

|name|属性|例子|说明|
|---|----|----|----|
|currency|string|fil|币种 fil，eth|
|num|string|231.321|提币数量|
|toAddress|string|0x96123...|提现地址 收账方|
|codeCheckPhone|string|02312|手机验证码 * 必须|
|codeTagPhone|string|hash123213|手机验证码标签 * 必须|
|codeCheckMailbox|string|02312|邮箱验证码 * 可选用户有绑定需认证|
|codeTagMailbox|string|hash123213|邮箱验证码标签 * 可选用户有绑定需认证|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|1 = 成功|
|~ |~ |~ | -1 = 完全失败|
|~ |~ |~ | -2 = 手机验证码错误|
|~ |~ |~ | -3 = 邮箱验证码错误|
|~ |~ |~ | -4 = 余额不足|
|~ |~ |~ | -5 = 收账账号异常|

------------
<span id="pool"></span>

## [**APP**] pool 矿池模块

<span id="/pool/productList"></span>
`post /pool/productList` 产品列表

请求 -无需登录

|name|属性|例子|说明|
|---|----|----|----|
|state|int|1|产品状态，1 = 进行中，2 = 已售馨，3 = 已结束|

返回

|name|属性|例子|说明|
|---|----|----|----|
|list| arr | [ ] |数据数组|
|`list.productId` |int |1 | 产品数据标识id|
|`list.currency`|string|fil|产品的 币种，目前只有fil|
|`list.name` |string |超算1期-FIL | 产品名称 |
|`list.pricePerT` |string | 6 | 产品价格 fil=默认每T价格|
|`list.pledgeLimit` |string |10600 | 质押上限 |
|`list.cycleDay` |string |540 | 挖矿周期（天） |
|`list.state` |int |1 | 产品状态，1 = 进行中，2 = 已售馨，3 = 已结束|

------------
<span id="/pool/htFilNodeData"></span>
`post /pool/htFilNodeData` 矿池数据 - app/联合挖矿

汉唐云app节点 矿池数据

请求 nil

返回

|name|属性|例子|说明|
|---|----|----|----|
|powerAll |string |10.1102 PiB | 总算力|
|produce24h |string |539.77 FIL | 24小时产出|
|produce24hAvg |string |0.077 FIL / TiB | 每T 24小时平均挖矿收益|
|produceAll |string |1539.77 FIL | 矿池总产出量|

------------
<span id="/pool/userFilData"></span>
`post AUTH /pool/userFilData` 矿池用户数据 - app/联合挖矿

请求:nil

返回

|name|属性|例子|说明|
|---|----|----|----|
|profitAll |string | 1300.001 | 累计收益 |
|profitBalance |string | 100.001 | 余额 |
|profitAvailable |string | 100.001 | 累计已释放收益|
|profitLock |string | 100.001 | 累计待释放收益|
|profitToday |string | 20.001 | 今日收益，今日发放的收益，今日还没发放返回0|
|pledgeAll |string | 900.001 | 质押数量|
|power |string | 12 | 我的算力（T）|

------------

<span id="/pool/profitLogList"></span>
`post AUTH /pool/profitLogList` 收益记录 - app/FIL 挖矿

请求:nil

返回

|name|属性|例子|说明|
|---|----|----|----|
|list |arr | [ ] | 收益记录|
|  `list.uuid`								 | string|  |数据标识     |  
|  `list.pledgeSum`            | string| 123 |质押金额     |
|  `list.holdPower`            | string| 123 |持有算力     |
|  `list.totalIncome`          | string| 123 |累计收益     |
|  `list.freezeIncome`         | string| 123 |累计冻结收益   |
|  `list.totalAvailableIncome` | string| 123 |累计已释放收益  |
|  `list.balance`              | string| 123 |可用余额     |
|  `list.dayAvailableIncome`   | string| 123 |当日释放收益   |
|  `list.dayRaiseIncome`       | string| 123 |日产出收益（当日直接25%+当日冻结75%） |
|  `list.dayDirectIncome`      | string| 123 |当日直接收益（25%）	|
|  `list.dayReleaseIncome`     | string| 123 |当日线性释放收益（往日75%平账）	|
|  `list.statTime`             | string| 2021-09-01 |统计时间（yyyy-MM-dd)|

<span id="/pool/profitExtract"></span>

------------

<span id="/pool/orderList"></span>
`post AUTH /pool/orderList` 订单记录 - app/FIL 挖矿

请求

|name|属性|例子|说明|
|---|----|----|----|
|state|int|1|订单状态 1 =挖矿中&进行中，2 = 结束|

返回

|name|属性|例子|说明|
|---|----|----|----|
|list|arr|[ ]|当前订单列表|
|`list.productId`|int|1|数据标识id|
|`list.currency`|string|fil|订单产品币种 fil，eth|
|`list.name`|string|超算1期-FIL|订单产品标题|
|`list.state`|int|1|订单状态 1 =挖矿中&进行中，2 = 结束|
|`list.userPledge`|string|100.00|用户的质押量|
|`list.userPower`|string|600.00|用户的算力（T）|
|`list.userProduceAll` |string |1539.77 | 用户购买产品产生的累计收益|
|`list.cycleDay`|string|540|产品周期（天） |
|`list.time`|string|2021.03.01 12:00:00|该产品买入时间|

------------

<span id="/pool/profitDrawing"></span>
`post AUTH /pool/profitDrawing` 提取收益 - app/充值到钱包

用户的矿池模块全部基本数据

请求

|name|属性|例子|说明|
|---|----|----|----|
|num|string|100|提取数量|
|currency|string|fil|币种 fil，eth|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1| 1 = 成功|
|~ |~ |~ | 2 = 成功且待审核|
|~ |~ |~ | -1 = 失败|
|~ |~ |~ | -2 = 失败2 待定|

------------

<span id="/pool/productOne"></span>
`post /pool/productOne` 矿池产品详情 - app/挖矿详情

用户的矿池模块全部基本数据（无需登录）

请求

|name|属性|例子|说明|
|---|----|----|----|
|productId|int|12|矿池产品id|

返回

|name|属性|例子|说明|
|---|----|----|----|
|productId |int |12 | 数据标识id|
|currency|string|fil|币种 fil，eth|
|state|string |1 | 产品状态，进行中 = 1，已售馨 = 2，已结束 = 3|
|name |string |超算1期-FIL | 产品名称|
|pricePerT |string | 6 | 产品价格 fil=默认每T价格|
|priceService |string | 40% | 技术服务费|
|Node1 |string | 文本 | 规则说明 |
|Node2 |string | 文本 | 挖矿规则|
|cycleDay |string | 540 | 产品周期=挖矿周期 （天）|
|guessProfitDay |string | 0.0012 | 预计日产出（T）|

------------

<span id="/pool/purchaseFrom"></span>
`post AUTH /pool/purchaseFrom` 购买矿池产品 - app/挖矿购买

请求

|name|属性|例子|说明|
|---|----|----|----|
|productId|int|12|矿池产品id|
|buyNum|int|2|买多少单位（T）|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode |int |1 | 1 = 成功 |
|～ |～ |～ | -1 =失败|
|～ |～ |～ | -2 = 余额不足|

-----
<span id="account"></span>

## [**APP**] account 账户模块

<span id="/account/signInCheck"></span>
`post /account/signInCheck` 登录提交（第一步检查资格）

用户登录有2步 这是第一步需下一步验证

请求

|name|属性|例子|说明|
|---|----|----|----|
|signInKey|string|13620000000|用户名|
|password|string|123456|密码|
|types|string|phone|登录类型 phone ，mailbox|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|1=成功 -1 = 失败|

------------
<span id="/account/signIn"></span>
`post /account/signIn` 登录提交（最终提交）

用户登录有2步 这是第一步需下一步验证

请求

|name|属性|例子|说明|
|---|----|----|----|
|signInKey|string|13620000000|用户名|
|password|string|123456|密码|
|codeCheck|string|0102|验证码|
|codeTag|string|hash2313|验证码标签|
|types|string|phone|登录类型 phone ，mailbox|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|1=成功 -1 = 失败|
|token|string|hash123213|登录成功分配token|

------------
<span id="/account/register"></span>
`post /account/register` 注册账号

注册成功 直接登录返回token

请求

|name|属性|例子|说明|
|---|----|----|----|
|registerKey|string|13620000000|用户名|
|password|string|123456|密码|
|codeCheck|string|0102|验证码|
|codeTag|string|hash1231|验证码标签|
|codeRecommend|string|12102|推荐码|
|types|string|phone|注册类型 phone ，mailbox|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|1=成功 -1 = 失败|
|token|string|hash123213|登录成功分配token|

------------
<span id="/account/passwordReset"></span>
`post /account/passwordReset` 忘记密码

请求

|name|属性|例子|说明|
|---|----|----|----|
|username|string|13620000000|用户名|
|passwordNew|string|123456|新密码|
|codeCheck|string|0102|验证码|
|codeTag|string|hash21312|验证码标签|
|types|string|phone|修改类型 phone ，mailbox|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|1=成功 |
|～|～|～|-1 = 失败|
|～|～|～|-2 = 验证码有误|

------------
<span id="/account/passwordUpdate"></span>
`post AUTH /account/passwordUpdate` 修改登录密码

请求

|name|属性|例子|说明|
|---|----|----|----|
|username|string|13620000000|用户名|
|passwordNew|string|123456|新密码|
|passwordOld|string|123456|旧密码|
|codeCheck|string|0102|验证码|
|codeTag|string|hash21312|验证码标签|
|types|string|phone|修改类型 phone ，mailbox|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|1=成功 |
|～|～|～|-1 = 失败|
|～|～|～|-2 = 验证码有误|

------------
<span id="/account/payPasswordSet"></span>
`post AUTH /account/payPasswordSet` 设置支付密码

请求

|name|属性|例子|说明|
|---|----|----|----|
|username|string|13620000000|用户名|
|passwordNew|string|123456|新密码|
|types|string|phone|修改类型 phone ，mailbox|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|1=成功 |
|～|～|～|-1 = 失败|

------------
<span id="/account/payPasswordUpdate"></span>
`post AUTH /account/payPasswordUpdate` 修改支付密码

请求

|name|属性|例子|说明|
|---|----|----|----|
|username|string|13620000000|用户名|
|passwordNew|string|123456|新密码|
|passwordOld|string|123456|旧密码|
|codeCheck|string|0102|验证码|
|codeTag|string|hash21312|验证码标签|
|types|string|phone|修改类型 phone ，mailbox|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|1=成功 |
|～|～|～|-1 = 失败|
|～|～|～|-2 = 验证码有误|

------------
<span id="/account/out"></span>
`post AUTH /account/out` 退出登录

请求:nil

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|1=成功 |
|～|～|～|-1 = 失败|

------------
<span id="/account/binding"></span>
`post AUTH /account/binding` 绑定手机号或邮箱

请求

|name|属性|例子|说明|
|---|----|----|----|
|bindingKey|string|13666666666|绑定 手机，邮箱|
|types|string|phone|绑定类型 phone ，mailbox|
|codeCheck|string|0121|验证码|
|codeTag|string|hash21312|验证码标签|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|1=成功 |
|～|～|～|-1 = 失败|

------------
<span id="verify"></span>

## [**APP**] verify 验证模块

<span id="/verify/sendCode"></span>
`post /verify/sendCode` 发送验证码

请求

|name|属性|例子|说明|
|---|----|----|----|
|signInKey|string|13620000000|用户名|
|types|string|phone|登录类型 phone ，mailbox|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|1=成功 -1 = 失败|
|codeTag|string|0121|验证码标签|
------------
<span id="/verify/signInKey"></span>
`post /verify/signInKey` 检查用户名是否存在

请求

|name|属性|例子|说明|
|---|----|----|----|
|signInKey|string|13620000000|用户名|
|types|string|phone|绑定类型 phone ，mailbox|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1| 1 = 存在|
|～|～|～|-1 = 不存在|
|～|～|～|-2 = 请求异常|

------------
<span id="market"></span>

## [**APP**] market 市场模块

<span id="/market/price"></span>
`post /market/price` 市场币价

币价默认 CNY换算

请求

|name|属性|例子|说明|
|---|----|----|----|
|currency|string |fil|币种 fil,eth,usdt|

返回

|name|属性|例子|说明|
|---|----|----|----|
|stateCode|int|1|返回状态 1 = 正常|
|～|～|-1|-1 = 异常|
|value|string|350.96|币价|

------------

<span id="user"></span>

## [**APP**] user 用户模块

<span id="/user/index"></span>
`post AUTH /user/index` 用户首页数据

请求:nil

返回

|name|属性|例子|说明|
|---|----|----|----|
|name|string|13620000000|用户名称|
|phone|string|13620000000|手机号 没绑定为 " " |
|mailbox|string|763999999@qq.com|邮箱 没绑定为 " " |

------------

<span id="/user/kycUser"></span>
`post AUTH /user/kycUser` 身份认证(KYC)

请求:nil

返回

|name|属性|例子|说明|
|---|----|----|----|
|name|string|13620000000|用户名称|
|state|int|0|身份认证状态 未认证-0 已认证-1|
|area|string|China|国家/地区|
|certifyType|int|1|证件类型 身份证-1 护照-2|
|realName|string|胡汉三|名字|
|certifyNum|string|530123456789|身份证号|
|birthday|string|1949/10/01|出生日期|
|address|string|北科大厦|居住地址|

------------

