import numpy as np
import pandas as pd
import datetime
import time
import pymysql
from collections.abc import Iterable
import sys
import os
import json
__dir__ = os.path.dirname(os.path.abspath(__file__))
sys.path.append(__dir__)
sys.path.append(os.path.abspath(os.path.join(__dir__, '..')))
from pygrpc import deliver_pb2

class position:
    judge:bool # 标记是否持仓(即持仓量是否为0)
    Insid = ""  # 持仓id
    pos = 0 # 持仓数量
    posSide = "" # 持仓方向
    avgPx = 0. # 平均价格
    cTime = "" # 建仓时间
    uTime = [] # 仓位更新时间列表
    clOrdId_list = [] # 更新该仓位的订单名称 
    def __init__(self,Insid:str):
        self.Insid = Insid
        self.judge = False
    def _reset(self):
        self.judge = False # 标记是否持仓(即持仓量是否为0)
        self.pos = 0 # 持仓数量
        self.posSide = "" # 持仓方向
        self.avgPx = 0. # 平均价格
        self.cTime = "" # 建仓时间
        self.uTime = [] # 仓位更新时间列表
        self.clOrdId_list = [] # 更新该仓位的订单名称 
    def UpdatePosition(self,order_respon):
        self.uTime.append(str(datetime.datetime.now()))
        self.cTime = self.uTime[0]
        self.clOrdId_list.append(order_respon["clOrdId"])
        if order_respon["side"] == "buy":
            if not self.judge:
                self.posSide = "long"
                self.avgPx = float(order_respon["avgPx"])
                self.pos = float(order_respon["accFillSz"])
            else:
                if self.pos + float(order_respon["accFillSz"]) == 0:
                    self._reset()
                else:
                    self.avgPx = (self.avgPx*self.pos + float(order_respon["avgPx"])*float(order_respon["accFillSz"]))/(float(order_respon["accFillSz"])+self.pos)
                    self.pos += float(order_respon["accFillSz"])
                    if self.pos > 0:
                        self.posSide = "long"
                    else:
                        self.posSide = "short"
        else:
            if not self.judge:
                self.posSide = "short"
                self.avgPx = float(order_respon["avgPx"])
                self.pos = -float(order_respon["accFillSz"])
            else:
                if self.pos - float(order_respon["accFillSz"]) == 0:
                    self._reset()
                else:
                    self.avgPx = (self.avgPx*self.pos - float(order_respon["avgPx"])*float(order_respon["accFillSz"]))/(self.pos-float(order_respon["accFillSz"]))
                    self.pos -= float(order_respon["accFillSz"])
                    if self.pos > 0:
                        self.posSide = "long"
                    else:
                        self.posSide = "short"
        if self.pos != 0:
            self.judge = True
            
    def GenInfo(self):
        return [self.judge,self.Insid,self.pos,self.posSide,self.avgPx,self.cTime," ".join(self.uTime)," ".join(self.clOrdId_list)]
    


class ordertemplate:
    insId = ""
    tdMode = ""
    ccy = ""
    clOrdId = ""
    tag = ""
    side = ""
    posSide = ""
    ordType = ""
    sz = ""
    px = ""
    reduceOnly = "" # bool
    tgtCcy = ""
    banAmend = "" # bool
    tpTriggerPx = ""
    tpOrdPx = ""
    slTriggerPx = ""
    slOrdPx = ""
    tpTriggerPxType = ""
    slTriggerPxType = ""
    quickMgnType = ""
    brokerID = ""
    def genOrder(self):
        temp = deliver_pb2.Order()
        if self.insId != "":
            temp.insId = self.insId
        if self.tdMode != "":
            temp.tdMode = self.tdMode
        if self.ccy != "":
            temp.ccy = self.ccy
        if self.side != "":
            temp.side = self.side
        if self.clOrdId != "":
            temp.clOrdId = self.clOrdId
        if self.tag != "":
            temp.tag = self.tag
        if self.posSide != "":
            temp.posSide = self.posSide
        if self.ordType != "":
            temp.ordType = self.ordType
        if self.sz != "":
            temp.sz = self.sz
        if self.px != "":
            temp.px = self.px
        if self.reduceOnly != "":
            if self.reduceOnly == "true" or self.reduceOnly == "True":
                temp.reduceOnly = True
            else:
                temp.reduceOnly = False
        if self.tgtCcy != "":
            temp.tgtCcy = self.tgtCcy
        if self.banAmend != "":
            if self.banAmend == "true" or self.banAmend == "True":
                temp.banAmend = True
            else:
                temp.banAmend = False
        if self.tpTriggerPx != "":
            temp.tpTriggerPx = self.tpTriggerPx
        if self.tpOrdPx != "":
            temp.tpOrdPx = self.tpOrdPx
        if self.slTriggerPx != "":
            temp.slTriggerPx = self.slTriggerPx
        if self.slOrdPx != "":
            temp.slOrdPx = self.slOrdPx
        if self.tpTriggerPxType != "":
            temp.tpTriggerPxType = self.tpTriggerPxType
        if self.slTriggerPxType != "":
            temp.slTriggerPxType = self.slTriggerPxType
        if self.quickMgnType != "":
            temp.quickMgnType = self.quickMgnType
        if self.brokerID != "":
            temp.brokerID = self.brokerID
        return temp
    def genInfoSimple(self):
        return [self.insId,self.tdMode,self.ccy,self.tdMode,self.clOrdId,self.side,self.posSide,self.ordType,self.sz,self.px]

class config:
    strategyname = ""
    subtype = ""
    barcustom = ""
    tickInsid = ""
    barInsid = ""
    tickPort = ""
    barPort = ""
    accountPort = ""
    # go端服务端口
    portsubmit = ""
    portorder = ""
    initjson = ""
    def genLocalSubmit(self):
        return deliver_pb2.LocalSubmit(subtype=self.subtype,barcustom=self.barcustom,tickInsid=self.tickInsid,barInsid=self.barInsid,tickPort=self.tickPort,barPort=self.barPort,accountPort=self.accountPort,strategyname=self.strategyname,initjson=self.initjson)
    
class barinfo:
    Insid = ""
    Ts_open = 0
    Open_price = 0
    High_price = 0
    Low_price = 0
    Close_price = 0
    Vol = 0
    VolCcy = 0
    VolCcyQuote = 0
    # Oi = 0
    # OiCcy = 0
    # Ts_oi = 0
    # FundingRate = 0
    # NextFundingRate = 0
    # Ts_FundingRate = 0
    # TS_NextFundingRate = 0
    def __init__(self,bar_info=""):
        if not isinstance(bar_info,Iterable):
            self.Insid =  bar_info.Insid 
            self.Ts_open = bar_info.Ts_open
            self.Open_price = bar_info.Open_price
            self.High_price = bar_info.High_price
            self.Low_price = bar_info.Low_price
            self.Close_price = bar_info.Close_price
            self.Vol = bar_info.Vol
            self.VolCcy = bar_info.VolCcy
            self.VolCcyQuote = bar_info.VolCcyQuote

class tickinfo:
    Insid = ""
    Ts_Price = 0
    Ask1_price = 0
    Bid1_price = 0
    Ask1_volumn = 0
    Bid1_volumn = 0
    def __init__(self,tick_info=""):
        if not isinstance(tick_info,Iterable):
            self.Insid = tick_info.Insid
            self.Ts_Price = tick_info.Ts_Price
            self.Ask1_price = tick_info.Ask1_price
            self.Bid1_price = tick_info.Bid1_price
            self.Ask1_volumn = tick_info.Ask1_volumn
            self.Bid1_volumn = tick_info.Bid1_volumn

class StrategyInitInfo:
    TradingMode:str
    LeverageSet:map
    TradingInsid:str
    def __init__(self,TradingMode:str,LeverageSet:map,TradingInsid:str):
        self.TradingMode = TradingMode
        self.LeverageSet = LeverageSet
        self.TradingInsid = TradingInsid
    def GenJsonStr(self):
        temp_info = {}
        temp_info["TradingMode"] = self.TradingMode
        temp_info["LeverageSet"] = self.LeverageSet
        temp_info["TradingInsid"] = self.TradingInsid
        return str(json.dumps(temp_info))
        

class BarinfoArray():
    Array = []
    df = pd.DataFrame()
    max_length = 0
    Symbol = ""
    def __init__(self,Insid="",max_length=10000):
        self.Symbol = Insid
        self.max_length = max_length
        self.df = pd.DataFrame(columns=["Insid","Ts_open","Open_price","High_price","Low_price","Close_price","Vol","VolCcy","VolCcyQuote"])
    def Store(self,bar_info):
        if isinstance(bar_info,Iterable):
            self._addnum(bar_info)
        else:
            self._add(barinfo(bar_info))
    def _add(self,value:barinfo):
        if self.Symbol == "":
            self.Symbol = value.Insid
        if value.Insid != self.Symbol:
            print(value)
            print("add in wrong BarinfoArray,try building a new one")
            return
        self.Array.append(value)
        temp_list = [value.Insid,value.Ts_open,value.Open_price,value.High_price,value.Low_price,value.Close_price,value.Vol,value.VolCcy,value.VolCcyQuote]
        self.df.loc[len(self.df)] = temp_list
        if len(self.Array) > self.max_length:
            self.Array = self.Array[1:]
            self.df.drop(0,inplace=True)
            self.df.reset_index(inplace=True,drop=True)
    def _addnum(self,value):
        self.df.loc[len(self.df)] = value
        self.Array.append(barinfo(self.df.loc[len(self.df)-1]))
    def Getlength(self):
        return len(self.df)
    def GetInsid(self):
        return self.Symbol
    def GetTsByTail(self,limit:int):
        if limit > len(self.df):
            limit = len(self.df)
        return np.array(self.df["Ts_open"].tail(limit)).astype("int")
    def GetOpenPriceByTail(self,limit:int):
        if limit > len(self.df):
            limit = len(self.df)
        return np.array(self.df["Open_price"].tail(limit)).astype("float")
    def GetHighPriceByTail(self,limit:int):
        if limit > len(self.df):
            limit = len(self.df)
        return np.array(self.df["High_price"].tail(limit)).astype("float")
    def GetLowPriceByTail(self,limit:int):
        if limit > len(self.df):
            limit = len(self.df)
        return np.array(self.df["Low_price"].tail(limit)).astype("float")
    def GetClosePriceByTail(self,limit:int):
        if limit > len(self.df):
            limit = len(self.df)
        return np.array(self.df["Close_price"].tail(limit)).astype("float")
    def GetVolByTail(self,limit:int):
        if limit > len(self.df):
            limit = len(self.df)
        return np.array(self.df["Vol"].tail(limit)).astype("float")

class TickinfoArray():
    Array = []
    df = pd.DataFrame()
    max_length = 0
    Symbol = ""
    def __init__(self,Insid="",max_length=10000):
        self.Symbol = Insid
        self.max_length = max_length
        self.df = pd.DataFrame(columns=["Insid","Ts_Price","Ask1_price","Bid1_price","Ask1_volumn","Bid1_volumn"])
    
    def Store(self,tick_info):
        if isinstance(tick_info,Iterable):
            self._addnum(tick_info)
        else:
            self._add(tickinfo(tick_info))
    
    def _add(self,value:tickinfo):
        if self.Symbol == "":
            self.Symbol = value.Insid
        if value.Insid != self.Symbol:
            print("add in wrong TickinfoArray,try building a new one")
            return
        self.Array.append(value)
        temp_list = [value.Insid,value.Ts_Price,value.Ask1_price,value.Bid1_price,value.Ask1_volumn,value.Bid1_volumn]
        self.df.loc[len(self.df)] = temp_list
        if len(self.Array) > self.max_length:
            # a1 = datetime.datetime.now()
            self.Array = self.Array[1:]
            self.df.drop(0,inplace=True)
            self.df.reset_index(inplace=True,drop=True)
            # a2 = datetime.datetime.now()
            # print((a2-a1).microseconds)
    def _addnum(self,value):
        self.df.loc[len(self.df)] = value
        self.Array.append(tickinfo(self.df.loc[len(self.df)-1]))
    def Getlength(self):
        return len(self.df)
    def GetInsid(self):
        return self.Symbol
    def GetTsByTail(self,limit:int):
        if limit > len(self.df):
            limit = len(self.df)
        return np.array(self.df["Ts_Price"].tail(limit)).astype("int")
    def GetAsk1PriceByTail(self,limit:int):
        if limit > len(self.df):
            limit = len(self.df)
        return np.array(self.df["Ask1_price"].tail(limit)).astype("float")
    def GetBid1PriceByTail(self,limit:int):
        if limit > len(self.df):
            limit = len(self.df)
        return np.array(self.df["Bid1_price"].tail(limit)).astype("float")
    def GetAsk1volumnByTail(self,limit:int):
        if limit > len(self.df):
            limit = len(self.df)
        return np.array(self.df["Ask1_volumn"].tail(limit)).astype("float")
    def GetBid1volumnByTail(self,limit:int):
        if limit > len(self.df):
            limit = len(self.df)
        return np.array(self.df["Bid1_volumn"].tail(limit)).astype("float")
    

# 调用该方法时，策略必须声明GenHourBarCustom方法
def genhourbarCustom(strategy,bardf:BarinfoArray,end_min:str):
    """调用该方法时,策略必须声明GenHourBarCustom方法"""
    last_series = bardf.df.iloc[-1]
    length = 60
    if str(time.localtime(float(last_series["Ts_open"])/1000).tm_min) == end_min:
        if len(bardf.df) >= length:
            tempbar = barinfo()
            tempbar.Insid = last_series["Insid"]
            tempbar.Ts_open = list(bardf.df["Ts_open"][-length:])[0]
            tempbar.Open_price = list(bardf.df["Open_price"][-length:])[0]
            tempbar.High_price = bardf.df["High_price"][-length:].max()
            tempbar.Low_price = bardf.df["Low_price"][-length:].min()
            tempbar.Close_price = list(bardf.df["Close_price"][-length:])[-1]
            tempbar.Vol = bardf.df["Vol"][-length:].sum()
            tempbar.VolCcy = bardf.df["VolCcy"][-length:].sum()
            tempbar.VolCcyQuote = bardf.df["VolCcyQuote"][-length:].sum()
            strategy.GenHourBarCustom(tempbar)
            
# 调用该方法时，策略必须声明GenHourBarCustom方法和hour_bar_calculation对象
def genhourbarCustomQuick(strategy,bardf:BarinfoArray,end_min:str):
    """调用该方法时，策略必须声明GenHourBarCustom方法和hour_bar_calculation对象"""
    last_series = bardf.df.iloc[-1]
    length = 60
    if strategy.hour_bar_calculation[last_series["Insid"]]["time_start"] == 0:
        strategy.hour_bar_calculation[last_series["Insid"]]["time_start"] = int(last_series["Ts_open"])
        strategy.hour_bar_calculation[last_series["Insid"]]["Open_price"] = float(last_series["Open_price"])
    strategy.hour_bar_calculation[last_series["Insid"]]["Close_price"] = float(last_series["Close_price"])
    if float(last_series["High_price"]) > strategy.hour_bar_calculation[last_series["Insid"]]["High_price"]:
        strategy.hour_bar_calculation[last_series["Insid"]]["High_price"] = float(last_series["High_price"])
    if float(last_series["Low_price"]) < strategy.hour_bar_calculation[last_series["Insid"]]["Low_price"]:
        strategy.hour_bar_calculation[last_series["Insid"]]["Low_price"] = float(last_series["Low_price"])
    strategy.hour_bar_calculation[last_series["Insid"]]["Vol"] += float(last_series["Vol"])
    strategy.hour_bar_calculation[last_series["Insid"]]["VolCcy"] += float(last_series["VolCcy"])
    strategy.hour_bar_calculation[last_series["Insid"]]["VolCcyQuote"] += float(last_series["VolCcyQuote"])
        
    if str(time.localtime(float(last_series["Ts_open"])/1000).tm_min) == end_min:
        if (last_series["Ts_open"]-strategy.hour_bar_calculation[last_series["Insid"]]["time_start"])/1000/60 + 1 == length:
            tempbar = barinfo()
            tempbar.Insid = last_series["Insid"]
            tempbar.Ts_open = strategy.hour_bar_calculation[last_series["Insid"]]["time_start"]
            tempbar.Open_price = strategy.hour_bar_calculation[last_series["Insid"]]["Open_price"]
            tempbar.High_price = strategy.hour_bar_calculation[last_series["Insid"]]["High_price"]
            tempbar.Low_price = strategy.hour_bar_calculation[last_series["Insid"]]["Low_price"]
            tempbar.Close_price = strategy.hour_bar_calculation[last_series["Insid"]]["Close_price"]
            tempbar.Vol = strategy.hour_bar_calculation[last_series["Insid"]]["Vol"]
            tempbar.VolCcy = strategy.hour_bar_calculation[last_series["Insid"]]["VolCcy"]
            tempbar.VolCcyQuote = strategy.hour_bar_calculation[last_series["Insid"]]["VolCcyQuote"]
            strategy.hour_bar_calculation[last_series["Insid"]] = {"time_start":0,"Open_price":0,"High_price":0,"Low_price":float('+inf'),"Close_price":0,"Vol":0,"VolCcy":0,"VolCcyQuote":0}
            strategy.GenHourBarCustom(tempbar)
        elif (last_series["Ts_open"]-strategy.hour_bar_calculation[last_series["Insid"]]["time_start"])/1000/60 + 1 < length:
            strategy.hour_bar_calculation[last_series["Insid"]] = {"time_start":0,"Open_price":0,"High_price":0,"Low_price":float('+inf'),"Close_price":0,"Vol":0,"VolCcy":0,"VolCcyQuote":0}

    
# 调用此方法的strategy必须包含UpdateBarCustom方法
def Load1MBarFromLocalMysql(strategy,user,password,database,Insid,length=0):
    """调用此方法的strategy必须包含UpdateBarCustom方法"""
    con = pymysql.connect(host="127.0.0.1",user=user,password=password,db=database)
    # 读取sql
    temp_df = pd.read_sql("select * from `" + Insid + "`;",con)
    info_matrix = temp_df[["Insid","Ts_open","Open_price","High_price","Low_price","Close_price","Vol","VolCcy","VolCcyQuote"]].drop_duplicates(subset=["Ts_open"], keep="first").sort_values(by="Ts_open",ascending=True).values
    if length == 0:
        length = len(info_matrix)
    info_matrix = info_matrix[-length:]
    for i in range(length):
        strategy.UpdateBarCustom(info_matrix[i])

# 调用此方法的strategy必须包含UpdateBarCustom方法
def Load1MBarFromCsv(strategy,path,length=0):
    """调用此方法的strategy必须包含UpdateBarCustom方法"""
    temp_df = pd.read_csv("path")
    info_matrix = temp_df[["Insid","Ts_open","Open_price","High_price","Low_price","Close_price","Vol","VolCcy","VolCcyQuote"]].drop_duplicates(subset=["Ts_open"], keep="first").sort_values(by="Ts_open",ascending=True).values
    if length == 0:
        length = len(info_matrix)
    info_matrix = info_matrix[-length:]
    for i in range(length):
        strategy.UpdateBarCustom(info_matrix[i])
    
# 调用此方法用于更新本地id值
def UpdateOrderId(temp):
    """调用此方法用于更新本地id值"""
    temp[0] += 1  
    return "0" + str(temp[0])

