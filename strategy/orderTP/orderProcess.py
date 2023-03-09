import sys
import os
__dir__ = os.path.dirname(os.path.abspath(__file__))
sys.path.append(__dir__)
sys.path.append(os.path.abspath(os.path.join(__dir__, '..')))
from tool import ToolUtil as TU

# # 提供一些回测处理订单的模板函数
# class ordertemplate:
#     insId = ""
#     tdMode = ""
#     ccy = ""
#     clOrdId = ""
#     tag = ""
#     side = ""
#     posSide = ""
#     ordType = ""
#     sz = ""
#     px = ""
#     reduceOnly = "" # bool
#     tgtCcy = ""
#     banAmend = "" # bool
#     tpTriggerPx = ""
#     tpOrdPx = ""
#     slTriggerPx = ""
#     slOrdPx = ""
#     tpTriggerPxType = ""
#     slTriggerPxType = ""
#     quickMgnType = ""
#     brokerID = ""

# def ProcessOrderByMinuteBarNet(order:TU.ordertemplate,bar_info:TU.barinfo,sl_tp:map):
#     """根据bar信息和报单信息判断该bar是否成交"""
#     # 判断报单是否成交
    
#     # 判断仓位是否止盈止损
    
#     # 
    
# def DealOrderInBarNet(order:TU.ordertemplate):
    
    
    
    
    