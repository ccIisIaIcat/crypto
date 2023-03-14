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
    StrategyName = "demo" # 策略名称，用于标注订单
    OrderNumber = [0] # 累加订单id，每次发送订单时OrderNumber+1
    # 持仓信息（简化起见当前策略只有一笔持仓）
    trade_forbidden_signal = True # 禁止交易信号
    # 已提交订单信息
    order_submit = []
    order_deal_submit = []
    
    def __init__(self,conf:TU.config):
        self.StrategyName = conf.strategyname
        self.basic_conf = conf
        
    def UpdateBarCustom(self,bar_info):
        pass
    
    def UpdateTick(self,tick_info):
        pass
        
    def LoadData(self):
        pass
    
    def Makeorder(self,order_info:TU.ordertemplate,order_deal:bool):
        """回测信号分为两种,一种是立即成交信号,一种是挂单信号,成交信号做标记,等待在之后的bar内成交"""
        order_info.clOrdId = self.StrategyName + TU.UpdateOrderId(self.OrderNumber)
        if order_deal:
            self.order_deal_submit.append(order_info)
        else:
            self.order_submit.append(order_info)
        return "order placed"
    
    def Pre_process(self,bar_info:TU.barinfo,Position:TU.position):
        # 处理deal订单，把deal订单处理为仓位
        for order_deal in self.order_deal_submit:
            Position.UpdateBackwardOrder(order_deal,bar_info.Ts_open)
        self.order_deal_submit = []
        undeal_list = []
        for order_undeal in self.order_submit:
            if order_undeal.BarInfoOpenJudge(bar_info):
                Position.UpdateBackwardOrder(order_deal,bar_info.Ts_open)
            else:
                undeal_list.append(order_undeal)
        self.order_submit = undeal_list
            

    def Start(self):
        self.LoadData()