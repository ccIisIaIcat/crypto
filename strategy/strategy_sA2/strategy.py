import sys
import os
__dir__ = os.path.dirname(os.path.abspath(__file__))
sys.path.append(__dir__)
sys.path.append(os.path.abspath(os.path.join(__dir__, '..')))
from strategy_template import strategyBasic
from tool import ToolUtil as TU
from pygrpc import pyserver
from orderTP import orderSA
import json
from collections.abc import Iterable
import pandas as pd
import numpy as np
import datetime

# 开多
odt_long = orderSA.sim_buy
# 开空
odt_short = orderSA.sim_sell



class strategySc(strategyBasic.strategy):
     # 策略对象
    MA_hour = [] # 小时bar的MA均线
    Std_hour = [] # 小时bar对应MA的std
    MA_grad = [] # MA梯度（该梯度为当前MA和之前MA_gap个MA的梯度，结果乘以1000）
    MA_grad_mean = [] # MA的梯度的平均
    basic_signal = [] # 信号指标1：（MA梯度和梯度平均的差值的绝对值）up代表大于阈值，below代表小于阈值
    
    bolling_up = [] # 当前布林带的上方
    bolling_down = [] # 当前布林带的下方
    
    waiting_cross_up = False
    waiting_cross_down = False
    # 对于小时bar信号的数据的记录
    signal_df = pd.DataFrame(columns=["record_time","MA","Signal"])
    trade_df = pd.DataFrame(columns=["judge","Insid","pos","posSide","avgPx","cTime","uTime","clOrdId_list"])
    
    # 策略参数
    MA_length = 20 # MA均线长度
    MA_gap = 20 # 计算梯度时的间隔长度
    MA_grade_mean_length = 20 # MA均线移动平均长度
    basic_signal_threshold = 40 # 信号指标阈值
    stop_lose_ratio = 0.03 # 止损比例
    bolling_std_k = 2 # 布林带方差个数
    
    def __init__(self,conf:TU.config):
        super().__init__(conf)
    # 声明存储对象
    tick_list = TU.TickinfoArray(Insid="ETH-USDT-SWAP",max_length=10)
    bar_list = TU.BarinfoArray(Insid="ETH-USDT-SWAP")
    bar_hour_list = TU.BarinfoArray(Insid="ETH-USDT-SWAP")
    # 声明仓位对象
    ETH_USDT_position = TU.position(Insid="ETH-USDT-SWAP")

    # 测试用
    test_a = 0
    
    def UpdateBarCustom(self,bar_info):
        # 更新本地bar列表
        self.bar_list.Store(bar_info)
        # 声明自定义bar方法，声明该方法时策略结构体必须包含GenHourBarCustom的类内函数
        TU.genhourbarCustom(self,self.bar_list,"59")

    
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
            self.signal_df.loc[len(self.signal_df)] = [str(datetime.datetime.now()),self.MA_hour[-1],self.basic_signal[-1]]
            self.signal_df.to_csv("./signal.csv")
            
        # 判断小时bar是否突破布林带
        if self.bar_hour_list.GetClosePriceByTail(1)[0] < self.bolling_down[-1]:
            # 判断是否大于阈值
            if self.basic_signal[-1] > self.basic_signal_threshold:
                # 开空
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
            if self.basic_signal[-1] < self.basic_signal_threshold:
                # 等待tick大于bolling_down，若大于，开多
                self.waiting_cross_down = True

            
        if self.bar_hour_list.GetClosePriceByTail(1)[0] > self.bolling_up[-1]:
            # 判断是否大于阈值
            if self.basic_signal[-1] > self.basic_signal_threshold:
                # 开多
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
            if self.basic_signal[-1] < self.basic_signal_threshold:
                # 等待tick小于bolling_up,若小于，开空
                self.waiting_cross_up = True

        
    def UpdateTick(self,tick_info):
        self.tick_list.Store(tick_info)
        # print(self.tick_list.GetAsk1PriceByTail(10))
        # print("/////////",self.tick_list.GetAsk1PriceByTail(1)[0])
        
        # 测试用
        self.test_a += 1
        print(self.test_a)
        print(self.order_record)
        print(self.ETH_USDT_position.GenInfo())
        print(self.trade_forbidden_signal)
        
        # 等待tick上穿
        if self.waiting_cross_down:
            if self.tick_list.GetAsk1PriceByTail(1)[0] > self.bolling_down[-1]:
                # 做多
                if not self.trade_forbidden_signal: 
                    if not self.ETH_USDT_position.judge:
                        odt_long.sz = "1"
                        self.trade_forbidden_signal = True
                        print(odt_long.genOrder())
                        self.Makeorder(odt_long)
                        self.order_record[odt_long.clOrdId] = 1
                    elif self.ETH_USDT_position.posSide == "short":
                        odt_long.sz = "2"
                        self.trade_forbidden_signal = True
                        res = self.Makeorder(odt_long)
                        self.order_record[odt_long.clOrdId] = 1
        
        # 等待tick下穿
        if self.waiting_cross_up:
            if self.tick_list.GetBid1PriceByTail(1)[-1] > self.bolling_up[-1]:
                # 做空
                if not self.trade_forbidden_signal: 
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
                        
        # 止盈止损判断
        if self.ETH_USDT_position.judge and not self.trade_forbidden_signal:
            # 多头止损判断
            if self.ETH_USDT_position.posSide == "long":
                price = self.ETH_USDT_position.avgPx
                # 触发止损
                if self.tick_list.GetAsk1PriceByTail(1)[0] < price*(1-self.stop_lose_ratio) :
                    # 平多
                    odt_short.sz = str(int(abs(self.ETH_USDT_position.pos)))
                    self.trade_forbidden_signal = True
                    self.Makeorder(odt_short)
                    self.order_record[odt_short.clOrdId] = 1
            # 空头止损判断
            if self.ETH_USDT_position.posSide == "short":
                price = self.ETH_USDT_position.avgPx
                if self.tick_list.GetBid1PriceByTail(1)[0] > price*(1+self.stop_lose_ratio):
                    # 平空
                    odt_long.sz = str(int(abs(self.ETH_USDT_position.pos)))
                    self.trade_forbidden_signal = True
                    self.Makeorder(odt_long)
                    self.order_record[odt_long.clOrdId] = 1
        
    def LoadData(self):
        TU.Load1MBarFromLocalMysql(self,"root","","crypto_swap","ETH-USDT-SWAP",5000)
        pass
    
    # def UpdateAccount(self, account_info):
    #     return super().UpdateAccount(account_info)
    
    def GatherOrder(self, ac_info):
        for order_respon in ac_info["data"]:
            print(order_respon)
            if order_respon["state"] == "canceled" or order_respon["state"] == "placed error":
                if order_respon["clOrdId"] in self.order_record:
                    del self.order_record[order_respon["clOrdId"]] 
            if order_respon["state"] == "live":
                continue
            if order_respon["state"] == "partially_filled" or order_respon["state"] == "filled":
                self.ETH_USDT_position.UpdatePosition(order_respon)
                self.trade_df.loc[len(self.trade_df)] = self.ETH_USDT_position.GenInfo()
                self.trade_df.to_csv("./trade_record.csv",index=False)
                if order_respon["state"] == "filled":
                    if order_respon["clOrdId"] in self.order_record:
                        del self.order_record[str(order_respon["clOrdId"])]
                    self.waiting_cross_up = False
                    self.waiting_cross_down = False
                continue
            
        if len(self.order_record) == 0:
            self.trade_forbidden_signal = False



if __name__ == '__main__':
    ############################################
    my_conf = TU.config()
    # 订阅对象.可选（tick,bar,account,position,order）
    my_conf.subtype = "tick order bar"
    # 策略名
    my_conf.strategyname = "sA2"
    # bar相关
    my_conf.barcustom = "1m"
    my_conf.barInsid = "ETH-USDT-SWAP"  # （多个品种用空格隔开）
    # tick相关
    my_conf.tickInsid = "ETH-USDT-SWAP" # （多个品种用空格隔开）
    # 端口相关
    my_conf.barPort = "6004"
    my_conf.tickPort = "6005"
    my_conf.accountPort = "6006"
    # go端端口
    my_conf.portsubmit = "6101"
    my_conf.portorder = "6102"
    ############################################
    datagather = strategySc(my_conf)
    datagather.Start()
    pyserver.NeverStop()