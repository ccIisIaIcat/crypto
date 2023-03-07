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

# 模拟盘中平多和开空的json一致,导入订单模板
# 开多
odt_long = orderSA.sim_buy
# 开空
odt_short = orderSA.sim_sell

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
    
    # 声明存储对象
    tick_list = TU.TickinfoArray(Insid="ETH-USDT-SWAP",max_length=10)
    bar_list = TU.BarinfoArray(Insid="ETH-USDT-SWAP")
    bar_hour_list = TU.BarinfoArray(Insid="ETH-USDT-SWAP")
    
    StrategyName = "sA" # 策略名称，用于标注订单
    OrderNumber = [0] # 累加订单id，每次发送订单时OrderNumber+1
    
    # 策略对象
    MA_hour = [] # 小时bar的MA均线
    Std_hour = [] # 小时bar对应MA的std
    MA_grad = [] # MA梯度（该梯度为当前MA和之前MA_gap个MA的梯度，结果乘以1000）
    MA_grad_mean = [] # MA的梯度的平均
    basic_signal = [] # 信号指标1：（MA梯度和梯度平均的差值的绝对值）up代表大于阈值，below代表小于阈值
    
    bolling_up = [] # 当前布林带的上方
    bolling_down = [] # 当前布林带的下方
    # 对于小时bar信号的数据的记录
    record_df = pd.DataFrame(columns=["MA","std","MA_grad","MA_grad_mean","basic_signal","price"])
    trade_df = pd.DataFrame(columns=["time","side","sz","bolling_up","bolling_down"])
    
    # 策略参数
    MA_length = 20 # MA均线长度
    MA_gap = 20 # 计算梯度时的间隔长度
    MA_grade_mean_length = 20 # MA均线移动平均长度
    basic_signal_threshold = 40 # 信号指标阈值
    stop_lose_ratio = 0.03 # 止损比例
    bolling_std_k = 2 # 布林带方差个数
    
    # 持仓信息（简化起见当前策略只有一笔持仓）
    position_record = [] # 仓位更改建议在UpdateAccount模块进行，该demo模块只记载持仓数
    order_record = {} # 未回执报单记录，只有当报单个数为0时，将trade_forbidden_signal标记为False
    trade_forbidden_signal = True # 禁止交易信号
    
    # 方便小时bar的快速计算
    hour_bar_calculation = {"ETH-USDT-SWAP":{"time_start":0,"Open_price":0,"High_price":0,"Low_price":float('+inf'),"Close_price":0,"Vol":0,"VolCcy":0,"VolCcyQuote":0}}
    
    # 测试
    test_a = 0
    
    
    def __init__(self,conf:TU.config):
        self.basic_conf = conf
        
    # 声明存储对象
    tick_list = TU.TickinfoArray(Insid="ETH-USDT-SWAP",max_length=10)
    bar_list = TU.BarinfoArray(Insid="ETH-USDT-SWAP")
    bar_hour_list = TU.BarinfoArray(Insid="ETH-USDT-SWAP")
    
    def UpdateBarCustom(self,bar_info):
        # 更新本地bar列表
        if isinstance(bar_info,Iterable):
            self.bar_list.addnum(bar_info)
        else:
            self.bar_list.add(TU.barinfo(bar_info))
        # 声明自定义bar方法，声明该方法时策略结构体必须包含GenHourBarCustom的类内函数
        
        if len(self.basic_signal) > 0 and self.basic_signal[-1]<self.basic_signal_threshold:
        # if  self.test_a == 30 or self.test_a == 180:
            if not self.trade_forbidden_signal: 
                # 当tick的ask小于下方布林带时,做多
                if list(self.bar_list.df["Low_price"])[-1] < self.bolling_down[-1]:
                # if self.test_a == 30 :
                    # 如果没有仓位，开仓
                    if len(self.position_record) == 0:
                        odt_long.clOrdId = self.StrategyName + str(TU.UpdateOrderId(self.OrderNumber))
                        odt_long.sz = "1"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_long)
                        self.trade_df.loc[len(self.trade_df)] = [list(self.bar_list.df["TS_open"])[-1],"get_signal","long"]
                        self.trade_df.to_csv("../trade_record.csv",index=False)
                    elif self.position_record[0]["posSide"] == "short":
                        odt_long.clOrdId = self.StrategyName + str(TU.UpdateOrderId(self.OrderNumber))
                        odt_long.sz = "2"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_long)
                        self.trade_df.loc[len(self.trade_df)] = [list(self.bar_list.df["TS_open"])[-1],"get_signal","long"]
                        self.trade_df.to_csv("../trade_record.csv",index=False)
                        
                # 当tick的bid大于上方布林带时,做空
                if list(self.bar_list.df["High_price"])[-1] > self.bolling_up[-1]:
                # if self.test_a == 180 :
                    # 如果没有仓位，开空
                    if len(self.position_record) == 0:
                        odt_short.clOrdId = self.StrategyName + str(TU.UpdateOrderId(self.OrderNumber))
                        odt_short.sz = "1"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_short)
                        self.trade_df.loc[len(self.trade_df)] = [list(self.bar_list.df["TS_open"])[-1],"get_signal","short"]
                        self.trade_df.to_csv("../trade_record.csv",index=False)
                    # 如果持有空仓，平多，开空
                    elif self.position_record[0]["posSide"] == "long":
                        odt_short.clOrdId = self.StrategyName + str(TU.UpdateOrderId(self.OrderNumber))
                        odt_short.sz = "2"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_short)
                        self.trade_df.loc[len(self.trade_df)] = [list(self.bar_list.df["TS_open"])[-1],"get_signal","short"]
                        self.trade_df.to_csv("../trade_record.csv",index=False)
        
        TU.genhourbarCustom(self,self.bar_list,"59")
        
                            
    def GenHourBarCustom(self,bar_info):
        self.bar_hour_list.add(bar_info)
        # 小时bar个数大于等于MA均线要求长度，生成MA
        if self.bar_hour_list.getlength() >= self.MA_length:
            temp_list = list(self.bar_hour_list.df["Close_price"])[-self.MA_length:]
            self.MA_hour.append(np.mean(temp_list))
            std = np.std(temp_list,ddof=1)
            self.Std_hour.append(std)
            self.bolling_down.append(self.MA_hour[-1] - self.bolling_std_k*std)
            self.bolling_up.append(self.MA_hour[-1] + self.bolling_std_k*std)
        # MA个数大于等于MA_gap时,生成MA均线梯度值
        if len(self.MA_hour) > self.MA_gap:
            self.MA_grad.append((self.MA_hour[-1]/self.MA_hour[-self.MA_gap-1]-1)*1000)
        # 当MA梯度个数大于要求的梯度长度，生成MA梯度均值
        if len(self.MA_grad) >= self.MA_grade_mean_length:
            self.MA_grad_mean.append(np.mean(self.MA_grad[-self.MA_grade_mean_length:]))
            # 当产生MA梯度平均时，计算信号指标
            self.basic_signal.append(abs(self.MA_grad_mean[-1]-self.MA_grad[-1]))
            print(self.basic_signal[-1])
            info_save = [self.MA_hour[-1],self.Std_hour[-1],self.MA_grad[-1],self.MA_grad_mean[-1],self.basic_signal[-1],list(self.bar_hour_list.df["Close_price"])[-1]]
            self.record_df.loc[len(self.record_df)] = info_save
            self.record_df.to_csv("./record.csv",index=False)
        # 当basic信号存在且大于阈值时考虑bar交易
        if len(self.basic_signal) > 0 and self.basic_signal[-1]>=self.basic_signal_threshold:
            # 若当前状况可交易
            if not self.trade_forbidden_signal:
                # 当当前bar的close大于上bolling up时，做多
                if float(self.bar_hour_list.df["Close_price"].iloc[-1]) > self.bolling_up[-1]:
                    # 如果没有仓位，开仓
                    if len(self.position_record) == 0:
                        odt_long.clOrdId = self.StrategyName + str(TU.UpdateOrderId(self.OrderNumber))
                        odt_long.sz = "1"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_long)
                    elif self.position_record[0]["posSide"] == "short":
                        odt_long.clOrdId = self.StrategyName + str(TU.UpdateOrderId(self.OrderNumber))
                        odt_long.sz = "2"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_long)
            # 当当前bar的close小于bolling down时，做空
                if float(self.bar_hour_list.df["Close_price"].iloc[-1]) < self.bolling_down[-1]:
                     # 如果没有仓位，开空
                    if len(self.position_record) == 0:
                        odt_short.clOrdId = self.StrategyName + str(TU.UpdateOrderId(self.OrderNumber))
                        odt_short.sz = "1"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_short)
                    # 如果持有空仓，平多，开空
                    elif self.position_record[0]["posSide"] == "long":
                        odt_short.clOrdId = self.StrategyName + str(TU.UpdateOrderId(self.OrderNumber))
                        odt_short.sz = "2"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_short)
                            
    def Makeorder(self,order_info:TU.ordertemplate):
        print(order_info.posSide)
        if len(self.position_record) != 0:
            self.position_record[0]["posSide"] = order_info.posSide
        else:
            self.position_record = [{}]
            self.position_record[0]["posSide"] = order_info.posSide
        self.trade_forbidden_signal = False

    
    def Start(self):
        self.trade_forbidden_signal = False
        TU.Load1MBarFromLocalMysql(self,"root","","crypto_swap","ETH-USDT-SWAP")
                    
if __name__ == '__main__':
    ############################################
    my_conf = TU.config()
    # 订阅对象.可选（tick,bar,account,position,order）
    ############################################
    datagather = strategy(my_conf)
    datagather.Start()
    pyserver.NeverStop()

