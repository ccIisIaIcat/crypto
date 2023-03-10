import sys
import os
__dir__ = os.path.dirname(os.path.abspath(__file__))
sys.path.append(__dir__)
sys.path.append(os.path.abspath(os.path.join(__dir__, '..')))
import datetime
import json
import numpy as np
import pandas as pd
from orderTP import orderSA
from collections.abc import Iterable
from tool import ToolUtil as TU

class strategy:
    # 策略固定参数
    StrategyName = "sA" # 策略名称，用于标注订单
    OrderNumber = [0] # 累加订单id，每次发送订单时OrderNumber+1
    # 持仓信息（简化起见当前策略只有一笔持仓）
    position_record = [] # 仓位更改建议在UpdateAccount模块进行，该demo模块只记载持仓数
    position = {}
    order_record = {} # 未回执报单记录，只有当报单个数为0时，将trade_forbidden_signal标记为False
    trade_forbidden_signal = True # 禁止交易信号
    # 已提交订单信息
    order_submit = []
    
    def __init__(self,conf:TU.config):
        self.StrategyName = conf.strategyname
        self.basic_conf = conf
        
    def UpdateBarCustom(self,bar_info):
        pass
    
    def UpdateTick(self,tick_info):
        pass
        
    def LoadData(self):
        pass
    
    def Makeorder(self,order_info:TU.ordertemplate,TradeInBar:bool,Position:TU.position):
        """回测信号分为两种,一种是立即成交信号,发送后立即转换为对应仓位,一种是挂单信号,等待在之后的bar内成交"""
        order_info.clOrdId = self.StrategyName + TU.UpdateOrderId(self.OrderNumber)
        if TradeInBar:
            self._tradeInBar(order_info)
        else:
            self._tradeAfterBar(order_info)
        return "order placed"
    
    def _tradeInBar(self,order_info:TU.ordertemplate,position:TU.position):
        
        pass
    
    def _tradeAfterBar(self,order_info:TU.ordertemplate):
        self.order_submit.append
        pass
    
    def Process_order(self,barinfo:TU.barinfo):
        """放置在类内对应函数中用于处理订单"""
        pass
    
    def Start(self):
        self.LoadData()