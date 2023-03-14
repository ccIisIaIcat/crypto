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

sim_limit_buy = orderSA.sim_buy_limit
sim_limit_sell = orderSA.sim_sell_limit
cancel_order = orderSA.cancel_order

# 止损市价单
odt_long = orderSA.sim_buy
odt_short = orderSA.sim_sell


class strategyKnifeDay(strategyBasic.strategy):
    ret_BTC_1D_4 = [] # 对数收益率绝对值
    ret_ETH_1D_8 = []
    ret_ETH_1D_12 = []
    ret_LTC_1D_28 = []
    
    ema_BTC_1D_4 = [] # EMA均线
    ema_ETH_1D_8 = []
    ema_ETH_1D_12 = []
    ema_LTC_1D_28 = []
    
    signal_BTC_1D_4 = [] # 生成信号
    signal_ETH_1D_8 = []
    signal_ETH_1D_12 = []
    signal_LTC_1D_28 = []
    
    len_BTC_1D_4 = 4 # ema均线长度
    len1_BTC_1D_4 = 365
    len_ETH_1D_8 = 8
    len1_ETH_1D_8 = 365
    len_ETH_1D_12 = 12
    len1_ETH_1D_12 = 365
    len_LTC_1D_28 = 28
    len1_LTC_1D_28 = 365
    
    thre_BTC_1D_4 = 0.1
    thre_ETH_1D_8 = 0.1
    thre_ETH_1D_12 = 0.1
    thre_LTC_1D_28 = 0.06
    
    thre1_BTC_1D_4 = 0.1
    thre1_ETH_1D_8 = 0.1
    thre1_ETH_1D_12 = 0.04
    thre1_LTC_1D_28 = 0.12
    
    var1_BTC_1D_4 = [] # 变量var1
    var2_BTC_1D_4 = [] # 变量var2
    var1_ETH_1D_8 = [] # 变量var1
    var2_ETH_1D_8 = [] # 变量var2
    var1_ETH_1D_12 = [] # 变量var1
    var2_ETH_1D_12 = [] # 变量var2
    var1_LTC_1D_28 = [] # 变量var1
    var2_LTC_1D_28 = [] # 变量var2
    
    
    stop_lose_ratio = 0.05 # 仓位止损
    
    barday_list_BTC_1D_4 = TU.TickinfoArray(Insid="BTC-USDT-SWAP")
    barday_list_ETH_1D_8 = TU.TickinfoArray(Insid="ETH-USDT-SWAP")
    barday_list_ETH_1D_12 = TU.TickinfoArray(Insid="ETH-USDT-SWAP")
    barday_list_LTC_1D_28 = TU.TickinfoArray(Insid="LTC-USDT-SWAP")
    
    BTC_1D_4_position = TU.position(Insid="BTC-USDT-SWAP")
    ETH_1D_8_position = TU.position(Insid="ETH-USDT-SWAP")
    ETH_1D_12_position = TU.position(Insid="ETH-USDT-SWAP")
    LTC_1D_28_position = TU.position(Insid="LTC-USDT-SWAP")
    
    
    def __init__(self,conf:TU.config):
        super().__init__(conf)
    
    def UpdateBarCustom(self,bar_info):
        TU.genCustomBar(self,TU.barinfo(bar_info),60*24,0)
        temp_bar = TU.barinfo(bar_info)
        if temp_bar.Insid == "BTC-USDT-SWAP"
        self.bar_list.Store(bar_info)
        # 声明自定义bar方法，声明该方法时策略结构体必须包含GenDayBarCustom的类内函数
        TU.gendaybarCustom(self,self.bar_list)
    
    def GenDayBarCustom(self,bar_info):
        self.bar_day_list.Store(bar_info)
        print(self.order_record)
        
         # 未成交的订单撤单
        for k in self.order_record.keys():
            cancel_order.clOrdId = k
            self.Makeorder(cancel_order)
        # 全部平仓
        if self.ETH_USDT_position.judge and not self.trade_forbidden_signal:
            # 已成交多头平仓
            if self.ETH_USDT_position.posSide == "long":
                price = self.ETH_USDT_position.avgPx
                # 平多
                odt_short.sz = str(int(abs(self.ETH_USDT_position.pos)))
                self.trade_forbidden_signal = True
                self.Makeorder(odt_short)
                self.order_record[odt_short.clOrdId] = 1
            # 已成交空头平仓
            if self.ETH_USDT_position.posSide == "short":
                price = self.ETH_USDT_position.avgPx
                # 平空
                odt_long.sz = str(int(abs(self.ETH_USDT_position.pos)))
                self.trade_forbidden_signal = True
                self.Makeorder(odt_long)
                self.order_record[odt_long.clOrdId] = 1
        
        # 计算收益率
        self.ret.append(abs(np.log(self.bar_day_list.GetClosePriceByTail(1)[0])-np.log(self.bar_day_list.GetOpenPriceByTail(1)[0])))
        
        # 计算ema均线
        if len(self.ret) > self.len_:
            if len(self.ema) == 0:
                self.ema.append(np.mean(self.ret[-self.len_:]))
            else:
                self.ema.append(self.ema[-1]*((self.len_-1)/(self.len_+1))+self.ret[-1]*(2/(self.len_+1)))
                print(self.ema[-1])
                if self.ema[-1] < self.threshold_variance and not self.trade_forbidden_signal:
                    # 开仓
                    # 计算变量var1,var2
                    price_up = []
                    price_down = []
                    var1 = max(np.max(self.ret[-self.len_:])*0.8,self.threshold_1)
                    var2 = np.max(np.max(self.ret[-self.len1_:]))
                    # 测试用
                    # var1 = 1
                    # var2 = 0.1
                    # 确定比率
                    ratio_list = np.linspace(start = var1, stop = var2, num = 6)[1:]
                    for ratio in ratio_list:
                        price_up.append(self.bar_day_list.GetClosePriceByTail(1)[0]*(ratio+1))
                    ratio_list = np.linspace(start = var1, stop = var1-var2,num=6)[1:]
                    for ratio in ratio_list:
                        if ratio != -1:
                            price_down.append(self.bar_day_list.GetClosePriceByTail(1)[0]*(ratio+1))
                    
                    # 根据不同价格发单
                    for price in price_up:
                        sim_limit_sell.px = str(price)
                        self.Makeorder(sim_limit_sell)
                        self.order_record[sim_limit_sell.clOrdId] = 1
                    for price in price_down:
                        sim_limit_buy.px = str(price)
                        self.Makeorder(sim_limit_buy)
                        self.order_record[sim_limit_buy.clOrdId] = 1
                
    def UpdateTick(self,tick_info):
        self.tick_list.Store(tick_info)
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
        TU.Load1MBarFromLocalMysql(self,"root","","crypto_swap","ETH-USDT-SWAP",360*24*60+300)
        pass
    
    
    def GatherOrder(self, ac_info):
        for order_respon in ac_info["data"]:
            # print(order_respon)
            if order_respon["state"] == "canceled" or order_respon["state"] == "placed error":
                print(order_respon["state"], order_respon["clOrdId"])
                if order_respon["clOrdId"] in self.order_record:
                    del self.order_record[order_respon["clOrdId"]] 
            if order_respon["state"] == "live":
                continue
            if order_respon["state"] == "partially_filled" or order_respon["state"] == "filled":
                self.ETH_USDT_position.UpdatePosition(order_respon)
                # self.trade_df.loc[len(self.trade_df)] = self.ETH_USDT_position.GenInfo()
                # self.trade_df.to_csv("./trade_record.csv",index=False)
                if order_respon["state"] == "filled":
                    if order_respon["clOrdId"] in self.order_record:
                        del self.order_record[str(order_respon["clOrdId"])]
                continue
            

if __name__ == '__main__':
    ############################################
    my_conf = TU.config()
    # 订阅对象.可选（tick,bar,account,position,order）
    my_conf.subtype = "tick order bar"
    # 策略名
    my_conf.strategyname = "knifeETH1"
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
    datagather = strategyKnifeDay(my_conf)
    datagather.Start()
    pyserver.NeverStop()
