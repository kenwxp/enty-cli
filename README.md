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

<span id="filer"></span>
## Filer


<span id="filer/filer_account_info"></span>
账户信息

filer_account_info

|name|说明|主键|                                                             
|---|---|---|  
|	filer_id	|	filer id	|	pk
|	filer_name	|	filer名称	|
|	area	|	地区	|
|	reg_time	|	注册时间	|
|	mobile	|	手机	|
|	email	|	邮箱	|
|	is_valid             	|	启用标志 0-启用 1-废弃	|
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

