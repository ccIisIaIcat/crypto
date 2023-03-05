import tool

# 针对该demo策略，设置的简易订单模板
# 开多
odt_buylong = tool.ordertemplate()
odt_buylong.insId = "ETH-USDT-SWAP"
odt_buylong.posSide = "long"
odt_buylong.tdMode = "cross"
odt_buylong.side = "buy"
odt_buylong.ordType = "market"
odt_buylong.sz = "1"
odt_buylong.clOrdId = ""
# 平多
odt_selllong = tool.ordertemplate()
odt_selllong.insId = "ETH-USDT-SWAP"
odt_selllong.posSide = "long"
odt_selllong.tdMode = "cross"
odt_selllong.side = "sell"
odt_selllong.ordType = "market"
odt_selllong.sz = "1"
odt_selllong.clOrdId = ""
# 开空
odt_buyshort = tool.ordertemplate()
odt_buyshort.insId = "ETH-USDT-SWAP"
odt_buyshort.posSide = "short"
odt_buyshort.tdMode = "cross"
odt_buyshort.side = "buy"
odt_buyshort.ordType = "market"
odt_buyshort.sz = "1"
odt_buyshort.clOrdId = ""
# 平空
odt_sellshort = tool.ordertemplate()
odt_sellshort.insId = "ETH-USDT-SWAP"
odt_sellshort.posSide = "short"
odt_sellshort.tdMode = "cross"
odt_sellshort.side = "sell"
odt_sellshort.ordType = "market"
odt_sellshort.sz = "1"
odt_sellshort.clOrdId = ""

# def buylong(strategy,order:tool.ordertemplate()):
#     ordername = statrgy.