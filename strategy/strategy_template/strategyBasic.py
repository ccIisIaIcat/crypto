import sys
import os
__dir__ = os.path.dirname(os.path.abspath(__file__))
sys.path.append(__dir__)
sys.path.append(os.path.abspath(os.path.join(__dir__, '..')))
import threading
import grpc
from pygrpc import deliver_pb2
from pygrpc import deliver_pb2_grpc
from pygrpc import pyserver
import datetime
import json
import numpy as np
import pandas as pd
from orderTP import orderSA
from collections.abc import Iterable
from tool import ToolUtil as TU

class strategy:
    # 交互参数
    porttick:str
    portbarcustom:str
    portaccount:str
    basic_conf:TU.config
    portsubmit:str
    portorder:str
    # 订单stud
    stub_order:deliver_pb2_grpc.OrerReceiverStub
    stub_submit:deliver_pb2_grpc.SubmitServerReceiverStub
    # 策略固定参数
    StrategyName = "sA" # 策略名称，用于标注订单
    OrderNumber = [0] # 累加订单id，每次发送订单时OrderNumber+1
    # 持仓信息（简化起见当前策略只有一笔持仓）
    order_record = {} # 未回执报单记录，只有当报单个数为0时，将trade_forbidden_signal标记为False
    trade_forbidden_signal = True # 禁止交易信号
    # 报单的本地存储
    order_df = pd.DataFrame()
    
    def __init__(self,conf:TU.config):
        self.LoadData()
        self.StrategyName = conf.strategyname
        self.basic_conf = conf
        if conf.tickPort != "":
            self.porttick = conf.tickPort
        if conf.barPort != "":
            self.portbarcustom = conf.barPort
        if conf.accountPort != "":
            self.portaccount = conf.accountPort
        self.portsubmit = conf.portsubmit
        self.portorder = conf.portorder
        self.order_df = pd.DataFrame(columns=["insId","tdMode","ccy","clOrdId","tag","side","posSide","ordType","sz","px"])
        channel = grpc.insecure_channel('localhost:'+self.portsubmit)
        self.stub_submit = deliver_pb2_grpc.SubmitServerReceiverStub(channel)
        channel2 = grpc.insecure_channel('localhost:'+self.portorder)
        self.stub_order = deliver_pb2_grpc.OrerReceiverStub(channel2)
        response = self.stub_submit.SubmitServerReceiver(conf.genLocalSubmit())
        print(response)
    
    def UpdateBarCustom(self,bar_info):
        pass
    
    def UpdateTick(self,tick_info):
        pass
        
    def LoadData(self):
        pass
    
    def Makeorder(self,order_info:TU.ordertemplate):
        if order_info.cancelOrder == "":
            order_info.clOrdId = self.StrategyName + TU.UpdateOrderId(self.OrderNumber)
            # self.order_df.loc[len(self.order_df)] = order_info.genInfoSimple()
            # self.order_df.to_csv("./order.csv",index=False)
            response = self.stub_order.OrerRReceiver(order_info.genOrder())
        else:
            response = self.stub_order.OrerRReceiver(order_info.genOrder())
        return order_info.clOrdId,response
    
    def UpdateAccount(self,account_info):
        format_info = json.loads(account_info)
        if "arg" in format_info:
            if format_info["arg"]["channel"] == "account":
                self.GatherAccount(format_info)
            if format_info["arg"]["channel"] == "positions":
                self.GatherPosition(format_info)
            if format_info["arg"]["channel"] == "orders":
                self.GatherOrder(format_info)
                
    def GatherAccount(self,ac_info):
        pass
    
    def GatherPosition(self,ac_info):
        pass
    
    def GatherOrder(self,ac_info):
        pass
        
    def Start(self):
        self.trade_forbidden_signal = False
        tick_thread = threading.Thread(target=pyserver.serveTick,args={self})
        barhour_thread = threading.Thread(target=pyserver.serveBarCustom,args={self})
        account_thread = threading.Thread(target=pyserver.serveAccount,args={self})
        pingpong_thread = threading.Thread(target=pyserver.servePingPong,args={self})
        tick_thread.start()
        barhour_thread.start()
        account_thread.start()
        pingpong_thread.start()
        