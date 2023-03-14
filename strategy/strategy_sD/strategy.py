import sys
import os
__dir__ = os.path.dirname(os.path.abspath(__file__))
sys.path.append(__dir__)
sys.path.append(os.path.abspath(os.path.join(__dir__, '..')))
from strategy_template import strategyBackward
from tool import ToolUtil as TU
from pygrpc import pyserver
from orderTP import orderSA
import json
from collections.abc import Iterable
import pandas as pd
import numpy as np

# 开多
odt_long = orderSA.sim_buy
# 开空
odt_short = orderSA.sim_sell

class strategy(strategyBackward.strategy):
    """回测模块市价单也需提交报单价格"""
     # 策略对象
    MA_hour = [] # 小时bar的MA均线
    Std_hour = [] # 小时bar对应MA的std
    MA_grad = [] # MA梯度（该梯度为当前MA和之前MA_gap个MA的梯度，结果乘以1000）
    MA_grad_mean = [] # MA的梯度的平均
    basic_signal = [] # 信号指标1：（MA梯度和梯度平均的差值的绝对值）up代表大于阈值，below代表小于阈值
    
    bolling_up = [] # 当前布林带的上方
    bolling_down = [] # 当前布林带的下方
    # 对于小时bar信号的数据的记录
    trade_df = pd.DataFrame(columns=["time","side","sz"])
    
    # 策略参数
    MA_length = 20 # MA均线长度
    MA_gap = 20 # 计算梯度时的间隔长度
    MA_grade_mean_length = 20 # MA均线移动平均长度
    basic_signal_threshold = 40 # 信号指标阈值
    stop_lose_ratio = 0.03 # 止损比例
    bolling_std_k = 2 # 布林带方差个数
    
    signal_df = pd.DataFrame(columns=["record_time","MA","Signal"])
    
    def __init__(self,conf:TU.config):
        super().__init__(conf)
    # 声明存储对象
    tick_list = TU.TickinfoArray(Insid="ETH-USDT-SWAP",max_length=10)
    bar_list = TU.BarinfoArray(Insid="ETH-USDT-SWAP")
    bar_hour_list = TU.BarinfoArray(Insid="ETH-USDT-SWAP")
    # 声明仓位对象
    ETH_USDT_position = TU.position(Insid="ETH-USDT-SWAP")
    
    
    def UpdateBarCustom(self,bar_info):
        self.Pre_process(bar_info=TU.barinfo(bar_info))
        
        # 更新本地bar列表
        self.bar_list.Store(bar_info)
        # 声明自定义bar方法，声明该方法时策略结构体必须包含GenHourBarCustom的类内函数
        TU.genhourbarCustom(self,self.bar_list,"59")

        pass
    
    def GenHourBarCustom(self,bar_info):
        self.bar_hour_list.Store(bar_info)
        # 小时bar个数大于等于MA均线要求长度，生成MA
        if self.bar_hour_list.Getlength() >= self.MA_length:
            self.bar_hour_list.GetClosePriceByTail(self.MA_length)
            temp_list = self.bar_hour_list.GetClosePriceByTail(self.MA_length)
            self.MA_hour.append(temp_list.mean())
            std = temp_list.std(ddof=1)
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
            self.signal_df.loc[len(self.signal_df)] = [self.bar_hour_list.GetTsByTail(1)[0],self.MA_hour[-1],self.basic_signal[-1]]
            self.signal_df.to_csv("./signal.csv")

        # 当basic信号存在且大于阈值时考虑bar交易
        if len(self.basic_signal) > 0 and self.basic_signal[-1]>=self.basic_signal_threshold:
            # 若当前状况可交易
            if not self.trade_forbidden_signal:
                # 当当前bar的close大于上bolling up时，做多
                if self.bar_hour_list.GetClosePriceByTail(1)[0] > self.bolling_up[-1]:
                    # 如果没有仓位，开仓
                    if not self.ETH_USDT_position.judge:
                        odt_long.sz = "1"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_long)
                        self.order_record[odt_long.clOrdId] = 1
                    elif self.ETH_USDT_position.posSide == "short":
                        odt_long.sz = "2"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_long)
                        self.order_record[odt_long.clOrdId] = 1
            # 当当前bar的close小于bolling down时，做空
                if self.bar_hour_list.GetClosePriceByTail(1)[0] < self.bolling_down[-1]:
                     # 如果没有仓位，开空
                    if not self.ETH_USDT_position.judge:
                        odt_short.sz = "1"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_short)
                        self.order_record[odt_short.clOrdId] = 1
                    # 如果持有空仓，平多，开空
                    elif self.ETH_USDT_position.posSide == "long":
                        odt_short.sz = "2"
                        self.trade_forbidden_signal = True
                        self.Makeorder(odt_short)
                        self.order_record[odt_short.clOrdId] = 1
                        
    