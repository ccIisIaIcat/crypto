import sys
import os
__dir__ = os.path.dirname(os.path.abspath(__file__))
sys.path.append(__dir__)
sys.path.append(os.path.abspath(os.path.join(__dir__, '..')))
from tool import ToolUtil

# 针对该demo策略，设置的简易订单模板
# 开多
odt_buylong = ToolUtil.ordertemplate()
odt_buylong.insId = "ETH-USDT-SWAP"
odt_buylong.posSide = "long"
odt_buylong.tdMode = "cross"
odt_buylong.side = "buy"
odt_buylong.ordType = "market"
odt_buylong.sz = "1"
odt_buylong.clOrdId = ""
# 平多
odt_selllong = ToolUtil.ordertemplate()
odt_selllong.insId = "ETH-USDT-SWAP"
odt_selllong.posSide = "long"
odt_selllong.tdMode = "cross"
odt_selllong.side = "sell"
odt_selllong.ordType = "market"
odt_selllong.sz = "1"
odt_selllong.clOrdId = ""
# 开空
odt_buyshort = ToolUtil.ordertemplate()
odt_buyshort.insId = "ETH-USDT-SWAP"
odt_buyshort.posSide = "short"
odt_buyshort.tdMode = "cross"
odt_buyshort.side = "buy"
odt_buyshort.ordType = "market"
odt_buyshort.sz = "1"
odt_buyshort.clOrdId = ""
# 平空
odt_sellshort = ToolUtil.ordertemplate()
odt_sellshort.insId = "ETH-USDT-SWAP"
odt_sellshort.posSide = "short"
odt_sellshort.tdMode = "cross"
odt_sellshort.side = "sell"
odt_sellshort.ordType = "market"
odt_sellshort.sz = "1"
odt_sellshort.clOrdId = ""

# 针对模拟交易，不存在odt_sellshort.posSide的long和short，因此新模板
# 买多
sim_buy = ToolUtil.ordertemplate()
sim_buy.insId = "ETH-USDT-SWAP"
sim_buy.posSide = "long"
sim_buy.tdMode = "cross"
sim_buy.side = "buy"
sim_buy.ordType = "optimal_limit_ioc"
sim_buy.sz = "1"
sim_buy.clOrdId = ""
# 买空
sim_sell = ToolUtil.ordertemplate()
sim_sell.insId = "ETH-USDT-SWAP"
sim_sell.posSide = "short"
sim_sell.tdMode = "cross"
sim_sell.side = "sell"
sim_sell.ordType = "optimal_limit_ioc"
sim_sell.sz = "1"
sim_sell.clOrdId = ""

