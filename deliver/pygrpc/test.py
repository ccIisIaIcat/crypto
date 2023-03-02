import talib
import numpy as np
import pandas as pd
import deliver_pb2
import deliver_pb2_grpc
import grpc
import tool
import datetime
import time
import json
from collections.abc import Iterable

a = tool.barinfo()

print(isinstance(a,Iterable))
print((1677232800000-1677229200000)/1000/60)

# my_conf = tool.config()
# my_conf.barInsid = "1"
# print(my_conf)

# print(str(time.localtime(1677662100).tm_min))
# if bool("False"):
#     print("a")

# my_conf = tool.config()
# my_conf.subtype("bar")
# my_conf.barcustom("1m")
# my_conf.barInsid("ETH-USDT-SWAP")
# my_conf.barPort("3904")
# datagather = strategy(my_conf)

# channel = grpc.insecure_channel('localhost:4352')
# # 客户端实例
# stub = deliver_pb2_grpc.SubmitServerReceiverStub(channel)
# # 调用服务端方法
# order_info = deliver_pb2.LocalSubmit(subtype = "tick")
# response = stub.SubmitServerReceiver(order_info)
# print("Greeter client received: " + response.response_me)

# df = pd.DataFrame(columns=["a","b","c"])
# df.loc[len(df)] = [1,2,3]
# print(df)
# print(df.loc[0].a)
# df["a"] = [1,2,3]
# df.loc[len(df.index)] = [2]

# df.drop(0,inplace=True)
# df.reset_index(inplace=True,drop=True)
# print(df)


# a = '''{"arg":{"channel":"account","uid":"156282789136867328"},"data":[{"adjEq":"","details":[{"availBal":"120.61512162667509","availEq":"120.61512162667509","cashBal":"120.61512162667509","ccy":"USDT","coinUsdPrice":"1","crossLiab":"","disEq":"120.61512162667509","eq":"120.61512162667509","eqUsd":"120.61512162667509","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1677650639840","upl":"0"},{"availBal":"0.0019991034","availEq":"0.0019991034","cashBal":"0.0019991034","ccy":"ETH","coinUsdPrice":"1649.38","crossLiab":"","disEq":"3.2972811658920005","eq":"0.0019991034","eqUsd":"3.2972811658920005","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1676903082349","upl":"0"},{"availBal":"0.005574","availEq":"0.005574","cashBal":"0.005574","ccy":"THETA","coinUsdPrice":"1.166","crossLiab":"","disEq":"0.003249642","eq":"0.005574","eqUsd":"0.006499284","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1617378821231","upl":"0"},{"availBal":"0.00001854","availEq":"0.00001854","cashBal":"0.00001854","ccy":"MLN","coinUsdPrice":"25.5","crossLiab":"","disEq":"0","eq":"0.00001854","eqUsd":"0.00047277","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1616145377671","upl":"0"},{"availBal":"0.000392","availEq":"0.000392","cashBal":"0.000392","ccy":"XRP","coinUsdPrice":"0.37911","crossLiab":"","disEq":"0.000126319452","eq":"0.000392","eqUsd":"0.00014861112","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1617379662057","upl":"0"},{"availBal":"0.009464","availEq":"0.009464","cashBal":"0.009464","ccy":"CHAT","coinUsdPrice":"0.005445","crossLiab":"","disEq":"0","eq":"0.009464","eqUsd":"0.00005153148","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1615570752542","upl":"0"},{"availBal":"0.00000096","availEq":"0.00000096","cashBal":"0.00000096","ccy":"BSV","coinUsdPrice":"41.26","crossLiab":"","disEq":"0.00003564864","eq":"0.00000096","eqUsd":"0.0000396096","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1616847274518","upl":"0"},{"availBal":"0.002758","availEq":"0.002758","cashBal":"0.002758","ccy":"SOC","coinUsdPrice":"0.001549","crossLiab":"","disEq":"0","eq":"0.002758","eqUsd":"0.000004272142","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1616092775173","upl":"0"},{"availBal":"0.0005632","availEq":"0.0005632","cashBal":"0.0005632","ccy":"MITH","coinUsdPrice":"0.00336","crossLiab":"","disEq":"0","eq":"0.0005632","eqUsd":"0.000001892352","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1615907630794","upl":"0"},{"availBal":"0.000000745","availEq":"0.000000745","cashBal":"0.000000745","ccy":"PERP","coinUsdPrice":"0.972","crossLiab":"","disEq":"0.00000036207","eq":"0.000000745","eqUsd":"0.00000072414","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1616134360243","upl":"0"},{"availBal":"0.00003632","availEq":"0.00003632","cashBal":"0.00003632","ccy":"IQ","coinUsdPrice":"0.00717","crossLiab":"","disEq":"0","eq":"0.00003632","eqUsd":"0.0000002604144","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1616129691693","upl":"0"},{"availBal":"0.00006144","availEq":"0.00006144","cashBal":"0.00006144","ccy":"XPR","coinUsdPrice":"0.00164","crossLiab":"","disEq":"0","eq":"0.00006144","eqUsd":"0.0000001007616","fixedBal":"0","frozenBal":"0","interest":"","isoEq":"0","isoLiab":"","isoUpl":"0","liab":"","maxLoan":"","mgnRatio":"","notionalLever":"0","ordFrozen":"0","spotInUseAmt":"","stgyEq":"0","twap":"0","uTime":"1616121852182","upl":"0"}],"imr":"","isoEq":"0","mgnRatio":"","mmr":"","notionalUsd":"","ordFroz":"","totalEq":"123.91979383790284","uTime":"1677727382994"}]}'''
# b = json.loads(a)
# print(b["data"][0]["details"][0])
