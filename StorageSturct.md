# 表结构设计
## 目录
- [Filer](#filer)
  - [账户信息](#filer/filer_account_info)
  - [账户登录](#filer/filer_account)
  - [账户平账信息](#filer/filer_account_income)
  - [账户余额信息](#filer/filer_balance_income)
  - [账户余额变动信息](#filer/filer_balance_flow)
  - [产品信息](#filer/filer_product)
  - [持仓信息](#filer/filer_order)
  - [持仓平账信息](#filer/filer_order_income)
  - [矿池信息](#filer/filer_pool)
  - [矿池平账信息](#filer/filer_pool_income)
  - [矿池状态表](#filer/filer_pool_status)
  - [爆块信息表](#filer/filer_block)
  - [爆块信息临时表](#filer/filer_block_temp)
  - [统计控制日志表](#filer/filer_stat_control)
  - [币种信息表](#filer/filer_currency)
  - [审核信息表](#filer/wallet_trans_approve)
- [Controller](#controller)
  - [后台用户表](#controller/controller_user)
  - [用户角色表](#controller/controller_role)
- [Wallet](#wallet)
  - [账户信息](#wallet/wallet_account)
  - [币种信息](#wallet/wallet_currency)
  - [地址管理](#wallet/wallet_address)
  - [余额信息](#wallet/wallet_balance)
  - [交易流转信息](#wallet/wallet_trans_flows)
  - [日运营统计](#wallet/wallet_stat_daily)
  - [实时运营统计](#wallet/wallet_stat_realtime)

<span id="filer"></span>  
## Filer


<span id="filer/filer_account_info"></span>
账户信息

filer_account_info

|name|说明|主键|                                                             
|---|---|---|  
|	filer_id	|	filer id	|	pk	
|	filer_name	|	filer名称	|		
|	filer_level	|	filer级别	|
|	acc_type	|	客户类型 0-customer 1-business	|		
|	area	|	地区	|		
|	reg_time	|	注册时间	|		
|	mobile	|	手机	|		
|	email	|	邮箱	|
|	is_valid             	|	启用标志 0-启用 1-废弃	|
------------

<span id="filer/filer_account"></span>
账户登录

filer_account

|name|说明|主键|                                                             
|---|---|---|  
|	filer_id	|	filer id	|	pk
|	pay_id |	filer名称	|
|	token	|	token	|
------------

<span id="filer/filer_account_income"></span>
账户平账信息

filer_account_income

|name|说明|主键|                                                             
|---|---|---|  
|	uu_id	|	uuid	|	pk
|	filer_id	|	filer id	|		
|	node_id	|	node id	|		
|	pledge_sum	|	质押金额	|
|	hold_power	|	持有算力	|
|	total_income	|	总收益	|		
|	freeze_income	|	冻结收益	|
|	total_available_income	|	累计已释放收益 = 昨日累计收益+当日释放收益	|
|	day_available_income	|  当日释放收益(当日直接25%+往日75%平账 ）    |
|   day_raise_income    |  日产出收益（当日直接25%+当日冻结75%） |
|	day_direct_income	|  当日直接收益（25%）	|
|	day_release_income	|  当日线性释放收益（往日75%平账）	|
|	stat_time	|	统计时间（yyyy-MM-dd)	|
|	create_time	|	创建时间 时间戳	|
------------

<span id="filer/filer_balance_income"></span>
账户余额信息

filer_balance_income

|name|说明|主键|                                                             
|---|---|---|  
|	uu_id	|	uuid	|	pk
|	filer_id	|	filer id	|
|	pledge_sum	|	质押金额	|
|	hold_power	|	持有算力	|
|	total_income	|	总收益	|
|	freeze_income	|	冻结收益	|
|	total_available_income	|	累计已释放收益 = 昨日累计收益+当日释放收益	|
|	balance	    |	余额 =昨日余额+ 当日释放收益 + 手工入账- 提币出账	|
|	day_available_income	|  当日释放收益(当日直接25%+往日75%平账 ）    |
|   day_raise_income    |  日产出收益（当日直接25%+当日冻结75%） |
|	day_direct_income	|  当日直接收益（25%）	|
|	day_release_income	|  当日线性释放收益（往日75%平账）	|
|	stat_time	|	统计时间（yyyy-MM-dd)	|
|	create_time	|	创建时间 时间戳	|
|   update_time |   更新时间 时间戳	|
------------

<span id="filer/filer_balance_flow"></span>
账户余额变动信息

filer_balance_flow

|name|说明|主键|                                                             
|---|---|---|  
|	uu_id	|	uuid	|	pk
|	filer_id	|	filer id	|
|   oper_type   |   操作类型 0-收益入账 1-手工入账 2-提现出账 |
|	amount	    |	金额	|
|	create_time	|	创建时间 时间戳	|
------------

<span id="filer/filer_product"></span>
产品信息

filer_product

|name|说明|主键|                                                             
|---|---|---|  
|	product_id	|	产品ID	|	pk	
|	product_name	|	产品名称	|	
|	node_id	|	节点	|		
|	cur_id	|	币种类型	|		
|	period	|	挖矿周期	|		
|	valid_plan	|	生效时间	|		
|	price	|	每T质押	|		
|	pledge_max	|	质押需求额	|		
|	service_rate	|	服务费率	|	
|   note1   |   说明（基本规则）|
|   note2   |   说明（联合挖矿说明）|
|	shelve_time	|	上架时间	|		
|	create_time	|	创建时间	|
|	update_time	|	更新时间	|
|	product_state	|	产品状态	0-进行中 9-已失效|		
|	is_valid             	|	启用标志 0--启用 1-废弃	|
------------

<span id="filer/filer_order"></span>
持仓信息

filer_order

|name|说明|主键|                                                             
|---|---|---|  
|	order_id	|	订单id	|	pk	
|	filer_id	|	filer id	|	
|   node_id     |   节点 id  |
|	pay_flow	|	交易流水（pay产生）	|		
|	product_id	|	产品id	|		
|	hold_power	|	持有算力	|		
|	pay_amount	|	支付金额	|		
|	order_time	|	下单时间	|	
|   update_time |   更新时间 |
|	valid_time	|	生效时间	|
|	end_time	|	结束时间	|
|	order_state	|	持仓状态	0-待生效 1-已生效 2-持仓已结束 9-收益结算完成|		
------------

<span id="filer/filer_order_income"></span>
持仓平账信息

filer_order_income

|name|说明|主键|                                                             
|---|---|---|  
|	uu_id	|	uuid	|	pk
|	order_id	|	订单id	|
|   filer_id    |   filer id (冗余) |
|   node_id     |   节点ID  |
|	total_income	|	总收益	|		
|	freeze_income	|	冻结收益	|
|	total_available_income	|	累计已释放收益 = 昨日累计收益+当日释放收益	|
|	day_available_income	|  当日释放收益(当日直接25%+往日75%平账 ）    |
|   day_raise_income    |  日产出收益（当日直接25%+当日冻结75%） |
|	day_direct_income	|  当日直接收益（25%）	|
|	day_release_income	|  当日线性释放收益（往日75%平账）	|
|	stat_time	|	统计时间	 （yyyy-MM-dd）|
|   create_time |   创建时间 时间戳 |
|   update_time |   更新时间 时间戳 |
------------
<span id="filer/filer_pool"></span>
矿池信息

filer_pool

|name|说明|主键|                                                             
|---|---|---|
|	node_id	|	节点id	|	pk	
|	node_name	|	节点名	|		
|	location	|	位置	|		
|	mobile	|	手机	|		
|	email	|	邮箱	|		
|	create_time	|	创建时间	|		
|	update_time	|	更新时间	|		
|	is_valid    |	启用标志	|
------------

<span id="filer/filer_pool_income"></span>
矿池平账信息

filer_pool_income

|name|说明|主键|                                                             
|---|---|---|  
|	uu_id	|	uuid	|	pk
|	node_id	|	节点id	|		
|   balance   |   矿池余额  |
|	pledge_sum	|	质押金额	|	
|   total_power |   总算力   |
|	total_income	|	总收益	|		
|	freeze_income	|	冻结收益	|		
|	available_income	|	可用收益	|
|	today_income_total	|	今日产出收益	|
|	today_income_freeze	|	今日冻结收益	|		
|	today_income_available	|	今日可用收益	|
|	stat_time	|	统计时间（yyyy-MM-dd）|
|	create_time	|	平账时间 时间戳|
------------
<span id="filer/filer_pool_status"></span>
矿池状态表

filer_pool_status（用于展示矿池待定）

|name|说明|主键|                                                             
|---|---|---|  
|	id	|	统计id（自增）	|	pk
|	node_id	|	节点id	|
|	valid_power	|	有效算力 T	|
|	power_raise	|	算力增量	T|
|	power_rate	|	算力增速	T/DAY|
|	win_num	|	出块数量	|
|	win_reward	|	出块奖励	|
|	efficiency	|	服务效率 FIL/T	|
|	lucky_rate	|	幸运值	%|
|	update_time	|	更新时间（具体时间）|
------------
<span id="filer/filer_block"></span>
爆块信息表

filer_block

|name|说明|主键|                                                             
|---|---|---|  
|	uu_id	|	uuid	|	pk
|	node_id	|	节点id	|
|	block_num	|	爆块数	|
|	block_gain	|	爆块奖励  FIL|
|   power   | 总算力   |
|   state   | 处理标志 0-待处理 1-线性释放 9-已处理|
|   gain_per_tib  |   每T收益  |
|	stat_time	|	统计时间（yyyy-MM-dd）	|
|   create_time  |   创建时间 时间戳 |
|	update_time	|	更新时间 时间戳|

<span id="filer/filer_block_temp"></span>
爆块信息临时表

filer_block_temp

|name|说明|主键|                                                             
|---|---|---|  
|	uu_id	|	uuid	|	pk
|	node_id	|	节点id	|
|	block_num	|	爆块数	|
|	block_gain	|	爆块奖励  FIL|
|	block_height	|	区块高度 |
|   power   | 总算力   |
|   state   | 处理标志 0-待处理 9-已处理|
|   chain_time  |   区块时间  |
|   create_time  |   创建时间  |
|	update_time	|	更新时间|

<span id="filer/filer_stat_control"></span>
统计控制日志表

filer_stat_control

|name|说明|主键|                                                             
|---|---|---|  
|	uu_id	|	uuid	|	pk
|	stat_type	|   统计类型 TASK_BLOCK_STAT,TASK_ORDER_STAT,TASK_ACCOUNT_STAT,TASK_POOL_STAT|
|	stat_time 	|	统计时间（yyyy-MM-dd)	|
|	node_id	    |	统计节点  |
|	stat_state	|	统计状态 0-新增 1-进行中，2-成功 3-失败 |
|   create_time   | 创建时间   |
|   update_time   | 更新时间  |
|   message     |   信息  |

<span id="filer/filer_currency"></span>
币种信息

filer_currency

注意：币种id和币种名与以太坊的配置一致

|name|说明|主键|                                                             
|---|---|---|  
|	cur_id	|	币种   唯一标识 列如：001，002    	|	pk	
|	cur_name       	|	币种名 列如：FIL             	|		
|	cur_desc             	|	币种介绍      列如：filcoin    	|		
|	withdraw_limit       	|	 最小提币数量                 	|		
|	staff_approve_limit  	|	 人工审核阈值                 	|		
|	manager_approve_limit	|	 高管审核阈值                 	|		
|	permit_withdraw      	|	是否允许提币 0-允许 1-禁止        	|		
|	permit_charge        	|	是否允许充值 0-允许 1-禁止         	|		
|	create_time	|	创建时间	|		
|	update_time	|	更新时间	|		
|	is_valid             	|	0-启动 1-废弃 （逻辑删除）        	|
------------

<span id="filer/wallet_trans_approve"></span>
审核信息表

wallet_trans_approve

|name|说明|主键|                                                             
|---|---|---|  
|	approve_id	|	审批id	|	pk	
|	filer_id	|	filer id	|		
|	cur_id	|	币种ID	|		
|	withdraw_amount  	|	提收益金额	|		
|	pay_flow	|	交易流水（pay产生）	|		
|	create_time	|	创建时间	|		
|	staff_approve_time 	|	交易员审核时间	|		
|	manager_approve_time	|	高管审核时间	|
|   trans_state   |   交易状态  0-待转账 1-已转账   |
|	approve_state	|	审核状态 0-新建待审核 1-待二级审核 2-通过 9-拒绝	|

<span id="system"></span>
## Controller

<span id="controller/controller_user"></span>
后台用户表

controller_user

|name|说明|主键|                                                             
|---|---|---|  
|	user_id	|	用户id	|	pk
|	role_id	|	角色id	|	
|	name	|	用户名	|	
|   encrypt |   加密密码 |
|   salt    |   盐      |
|	token	|	密码	|	
|	mobile	|	手机	|	
|	email	|	邮箱	|	
|	create_time	|	创建时间	|	
|	logon_time	|	登录时间	|	
|	is_valid	|	启用标志 0-启用 1-废弃	|	
------------

<span id="controller/controller_role"></span>
用户角色表

controller_role

|name|说明|主键|                                                             
|---|---|---|  
|	role_id	|	角色id	|	pk
|	role_name	|	角色名	|	
|	menu_items	|	菜单项	|	
|	create_time	|	创建时间	|	
|	update_time	|	更新时间	|	
|	is_valid	|	启用标志 0-启用 1-废弃	|	
------------


<span id="wallet"></span>
Wallet

<span id="wallet/account"></span>
账户信息

account

|name|说明|主键|                                                             
|---|---|---|
|	id	|	账户id	|	pk
|	phone	|	账户name	|	
|	mailbox	|	客户类型 0-customer 1-business	|	
|	cypher	|	地区	|	
|	salt	|	注册时间	|	
|	token	|	手机	|	
------------

<span id="wallet/wallet_currency"></span>
币种信息

wallet_currency

注意：币种id和币种名与以太坊的配置一致

|name|说明|主键|                                                             
|---|---|---|  
|	cur_id	|	币种   唯一标识 列如：001，002    	|	pk
|	cur_name       	|	币种名 列如：FIL             	|	
|	cur_desc             	|	币种介绍      列如：filcoin    	|	
|	create_time	|	创建时间	|	
|	update_time	|	更新时间	|	
|	is_valid             	|	0-启动 1-废弃 （逻辑删除）        	|	
------------
<span id="wallet/wallet_address"></span>
地址管理

wallet_address

|name|说明|主键|                                                             
|---|---|---|  
|	pay_id	|	账户id	|	pk
|	cur_id	|	币种	|	pk
|	address_type	|	地址类型	|	pk
|	trans_address	|	交易地址	|	
|	fee_address	|	手续费地址	|	
|	is_valid	|	启用标志 0-启用 1-废弃	|
------------

<span id="wallet/wallet_balance"></span>
余额信息

wallet_balance

|name|说明|主键|                                                             
|---|---|---|  
|	pay_id	|	账户id	|	pk
|	cur_id	|	币种	|	pk
|	trans_address	|	交易地址	|	
|	available_balance	|	可用余额	|	
|	freezen_balance	|	冻结余额	|	
|	total_balance	|	汇总余额	|	
|	update_time	|	更新时间	|	
|	lock	|	修改标志 0-可修改，1-不可修改	|
------------

<span id="wallet/wallet_trans_flows"></span>				
交易流转信息表		

wallet_trans_flows

|name|说明|主键|                                                             
|---|---|---|  
|	trans_id	|	交易id	|	pk
|	trans_type	|	交易类型  0-充值，1-提现，2-支付，3-提现到钱包	|	
|	trans_currency	|	交易币种	|	
|	trans_amount	|	交易金额	|	
|	trans_fee	|	交易手续费	|	
|	trans_from_account	|	出账账户	|	
|	trans_from_address	|	出账地址	|	
|	trans_to_account	|	入账账户	|	
|	trans_to_address	|	入账地址	|	
|	trans_hash	|	交易hash	|	
|	trans_state	|	交易转态	0-成功 9-失败|	
|	update_time	|	交易时间	|	
|	create_time	|	创建时间	|	
|	remark              	|	备注	|
------------

<span id="wallet/wallet_stat_daily"></span>				
日运营统计

wallet_stat_daily

注意：跑批 一日一次

|name|说明|主键|                                                             
|---|---|---|  
|	stat_id	|	统计Id	|	pk
|	cur_id	|	币种	|	
|	total_balance	|	汇总余额	|	
|	charge_account_num	|	充值用户数	|	
|	charge_total_times	|	充值笔数	|	
|	charge_total_amount	|	充值数量	|	
|	withdraw_account_num	|	提现用户数	|	
|	withdraw_total_times	|	提现笔数	|	
|	withdraw_total_amount	|	提现数量	|	
|	pledge_amount	|	矿池质押(矿池提现)数额	|	
|	total_income	|	矿池收入(矿池充值)数额	|	
|	stat_time	|	统计时间	|	
------------

<span id="wallet/wallet_stat_realtime"></span>					
实时运营统计

wallet_trans_stat_realtime

注意：跑批 一小时一次 频率待定

|name|说明|主键|                                                             
|---|---|---|  
|	stat_id	|	统计Id	|	pk
|	cur_id	|	币种	|	
|	total_balance	|	汇总余额	|	
|	charge_account_num	|	充值用户数	|	
|	charge_total_times	|	充值笔数	|	
|	charge_total_amount	|	充值数量	|	
|	withdraw_account_num	|	提现用户数	|	
|	withdraw_total_times	|	提现笔数	|	
|	withdraw_total_amount	|	提现数量	|	
|	pledge_amount	|	矿池质押(矿池提现)数额	|	
|	total_income	|	矿池收入(矿池充值)数额	|	
|	stat_time	|	统计时间	|
------------

