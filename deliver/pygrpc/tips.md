关于仓位记录：
本地每一笔仓位由一个map记录，记录字段

> mgnMode	String	保证金模式， cross：全仓 isolated：逐仓
> posId	String	持仓ID
> posSide	String	持仓方向 long：双向持仓多头 short：双向持仓空头 net：单向持仓（交割/永续/期权：pos为正代表多头，pos为负代表空头。
> pos	String	持仓数量，逐仓自主划转模式下，转入保证金后会产生pos为0的仓位
> availPos	String	可平仓数量，适用于 币币杠杆,交割/永续（开平仓模式），期权（交易账户及保证金账户逐仓）。
> avgPx	String	开仓平均价
> upl	String	未实现收益
> uplRatio	String	未实现收益率
> instId	String	产品ID，如 BTC-USD-180216
> lever	String	杠杆倍数，不适用于期权卖方
> liqPx	String	预估强平价
> markPx	String	标记价格
> imr	String	初始保证金，仅适用于全仓
> margin	String	保证金余额，仅适用于逐仓，可增减
> mgnRatio	String	保证金率
> mmr	String	维持保证金
> interest	String	利息，已经生成未扣利息
> tradeId	String	最新成交ID
> notionalUsd	String	以美金价值为单位的持仓数量
> optVal	String	期权价值，仅适用于期权
> ccy	String	占用保证金的币种
> last	String	最新成交价
> usdPx	String	美金价格
> cTime	String	持仓创建时间，Unix时间戳的毫秒数格式，如 1597026383085
> uTime	String	最近一次持仓更新时间，Unix时间戳的毫秒数格式，如 1597026383085
> pTime	String	持仓信息的推送时间，Unix时间戳的毫秒数格式，如 1597026383085

关于订单更新：
> instType	String	产品类型
> instId	String	产品ID
> ccy	String	保证金币种，仅适用于单币种保证金账户下的全仓币币杠杆订单
> ordId	String	订单ID
> clOrdId	String	由用户设置的订单ID来识别您的订单
> tag	String	订单标签
> px	String	委托价格
> sz	String	原始委托数量，币币/币币杠杆，以币为单位；交割/永续/期权 ，以张为单位
> notionalUsd	String	委托单预估美元价值
> ordType	String	订单类型
market：市价单
limit：限价单
post_only： 只做maker单
fok：全部成交或立即取消单
ioc：立即成交并取消剩余单
optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续）
> side	String	订单方向，buy sell
> posSide	String	持仓方向
long：双向持仓多头
short：双向持仓空头
net：单向持仓
> tdMode	String	交易模式
保证金模式 isolated：逐仓 cross：全仓
非保证金模式 cash：现金
> tgtCcy	String	市价单委托数量sz的单位
base_ccy: 交易货币 quote_ccy：计价货币
> fillPx	String	最新成交价格
> tradeId	String	最新成交ID
> fillSz	String	最新成交数量
对于币币和杠杆，单位为交易货币，如 BTC-USDT, 单位为 BTC；对于市价单，无论tgtCcy是base_ccy，还是quote_ccy，单位均为交易货币；
对于交割、永续以及期权，单位为张。
> fillTime	String	最新成交时间
> fillFee	String	最新一笔成交的手续费金额或者返佣金额：
手续费扣除 为 ‘负数’，如 -0.01 ；
手续费返佣 为 ‘正数’，如 0.01
> fillFeeCcy	String	最新一笔成交的手续费币种
> execType	String	最新一笔成交的流动性方向 T：taker M：maker
> accFillSz	String	累计成交数量
对于币币和杠杆，单位为交易货币，如 BTC-USDT, 单位为 BTC；对于市价单，无论tgtCcy是base_ccy，还是quote_ccy，单位均为交易货币；
对于交割、永续以及期权，单位为张。
> avgPx	String	成交均价，如果成交数量为0，该字段也为0
> state	String	订单状态
canceled：撤单成功
live：等待成交
partially_filled： 部分成交
filled：完全成交
> lever	String	杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
> tpTriggerPx	String	止盈触发价
> tpTriggerPxType	String	止盈触发价类型
last：最新价格
index：指数价格
mark：标记价格
> tpOrdPx	String	止盈委托价，止盈委托价格为-1时，执行市价止盈
> slTriggerPx	String	止损触发价
> slTriggerPxType	String	止损触发价类型
last：最新价格
index：指数价格
mark：标记价格
> slOrdPx	String	止损委托价，止损委托价格为-1时，执行市价止损
> feeCcy	String	交易手续费币种
币币/币币杠杆：如果是买的话，收取的就是BTC；如果是卖的话，收取的就是USDT
交割/永续/期权 收取的就是保证金
> fee	String	订单交易累计的手续费与返佣
对于币币和杠杆，为订单交易累计的手续费，平台向用户收取的交易手续费，为负数。如： -0.01
对于交割、永续和期权，为订单交易累计的手续费和返佣
> rebateCcy	String	返佣金币种 ，如果没有返佣金，该字段为“”
> rebate	String	返佣累计金额，仅适用于币币和杠杆，平台向达到指定lv交易等级的用户支付的挂单奖励（返佣），如果没有返佣金，该字段为“”
> pnl	String	收益，适用于有成交的平仓订单，其他情况均为0
> source	String	订单来源
13:策略委托单触发后的生成的限价单
> cancelSource	String	订单取消的来源
有效值及对应的含义是：
0: 已撤单：系统撤单
1: 用户主动撤单
2: 已撤单：预减仓撤单，用户保证金不足导致挂单被撤回
3: 已撤单：风控撤单，用户保证金不足有爆仓风险，导致挂单被撤回
4: 已撤单：币种借币量达到平台硬顶，系统已撤回该订单
6: 已撤单：触发 ADL 撤单，用户保证金率较低且有爆仓风险，导致挂单被撤回
9: 已撤单：扣除资金费用后可用余额不足，系统已撤回该订单
13: 已撤单：FOK 委托订单未完全成交，导致挂单被完全撤回
14: 已撤单：IOC 委托订单未完全成交，仅部分成交，导致部分挂单被撤回
17: 已撤单：平仓单被撤单，由于仓位已被市价全平
20: 系统倒计时撤单
21: 已撤单：相关仓位被完全平仓，系统已撤销该止盈止损订单
22, 23: 已撤单：只减仓订单仅允许减少仓位数量，系统已撤销该订单
> category	String	订单种类分类
normal：普通委托订单种类
twap：TWAP订单种类
adl：ADL订单种类
full_liquidation：爆仓订单种类
partial_liquidation：减仓订单种类
delivery：交割
ddh：对冲减仓类型订单
> uTime	String	订单更新时间，Unix时间戳的毫秒数格式，如 1597026383085
> cTime	String	订单创建时间，Unix时间戳的毫秒数格式，如 1597026383085
> reqId	String	修改订单时使用的request ID，如果没有修改，该字段为""
> amendResult	String	修改订单的结果
-1： 失败
0：成功
1：自动撤单（因为修改失败导致订单自动撤销）
通过API修改订单时，如果cxlOnFail设置为false且修改失败后，则amendResult返回 -1
通过API修改订单时，如果cxlOnFail设置为true且修改失败后，则amendResult返回1
通过Web/APP修改订单时，如果修改失败后，则amendResult返回-1
> reduceOnly	String	是否只减仓，true 或 false
> quickMgnType	String	一键借币类型，仅适用于杠杆逐仓的一键借币模式
manual：手动，auto_borrow： 自动借币，auto_repay： 自动还币
> code	String	错误码，默认为0
> msg	String	错误消息，默认为""