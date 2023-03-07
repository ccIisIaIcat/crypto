import numpy as np
import pandas as pd
import datetime
import time
import pymysql
from collections.abc import Iterable
import sys
import os
__dir__ = os.path.dirname(os.path.abspath(__file__))
sys.path.append(__dir__)
sys.path.append(os.path.abspath(os.path.join(__dir__, '..')))
from pygrpc import deliver_pb2

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
    def genLocalSubmit(self):
        return deliver_pb2.LocalSubmit(subtype=self.subtype,barcustom=self.barcustom,tickInsid=self.tickInsid,barInsid=self.barInsid,tickPort=self.tickPort,barPort=self.barPort,accountPort=self.accountPort,strategyname=self.strategyname)
    
class barinfo:
    InsID = ""
    TS_open = 0
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
            self.InsID =  bar_info.Insid 
            self.TS_open = bar_info.Ts_open
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

class BarinfoArray():
    Array = []
    df = pd.DataFrame()
    max_length = 0
    Symbol = ""
    def __init__(self,Insid="",max_length=10000):
        self.Symbol = Insid
        self.max_length = max_length
        self.df = pd.DataFrame(columns=["Insid","TS_open","Open_price","High_price","Low_price","Close_price","Vol","VolCcy","VolCcyQuote"])

    def add(self,value:barinfo):
        if self.Symbol == "":
            self.Symbol = value.InsID
        if value.InsID != self.Symbol:
            print("add in wrong BarinfoArray,try building a new one")
            return
        self.Array.append(value)
        temp_list = [value.InsID,value.TS_open,value.Open_price,value.High_price,value.Low_price,value.Close_price,value.Vol,value.VolCcy,value.VolCcyQuote]
        self.df.loc[len(self.df)] = temp_list
        if len(self.Array) > self.max_length:
            self.Array = self.Array[1:]
            self.df.drop(0,inplace=True)
            self.df.reset_index(inplace=True,drop=True)
    def addnum(self,value):
        # print(value)
        self.df.loc[len(self.df)] = value
        self.Array.append(barinfo(self.df.loc[len(self.df)-1]))
    def getlength(self):
        return len(self.df)

class TickinfoArray():
    Array = []
    df = pd.DataFrame()
    max_length = 0
    Symbol = ""
    def __init__(self,Insid="",max_length=10000):
        self.Symbol = Insid
        self.max_length = max_length
        self.df = pd.DataFrame(columns=["Insid","Ts_Price","Ask1_price","Bid1_price","Ask1_volumn","Bid1_volumn"])
    
    def add(self,value:tickinfo):
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
    def addnum(self,value):
        self.df.loc[len(self.df)] = value
        self.Array.append(tickinfo(self.df.loc[len(self.df)-1]))
    def getlength(self):
        return len(self.df)

# 调用该方法时，策略必须声明GenHourBarCustom方法
def genhourbarCustom(strategy,bardf:BarinfoArray,end_min:str):
    last_series = bardf.df.iloc[-1]
    length = 60
    if str(time.localtime(float(last_series["TS_open"])/1000).tm_min) == end_min:
        if len(bardf.df) >= length:
            tempbar = barinfo()
            tempbar.InsID = last_series["Insid"]
            tempbar.TS_open = list(bardf.df["TS_open"][-length:])[0]
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
    last_series = bardf.df.iloc[-1]
    length = 60
    if strategy.hour_bar_calculation[last_series["Insid"]]["time_start"] == 0:
        strategy.hour_bar_calculation[last_series["Insid"]]["time_start"] = int(last_series["TS_open"])
        strategy.hour_bar_calculation[last_series["Insid"]]["Open_price"] = float(last_series["Open_price"])
    strategy.hour_bar_calculation[last_series["Insid"]]["Close_price"] = float(last_series["Close_price"])
    if float(last_series["High_price"]) > strategy.hour_bar_calculation[last_series["Insid"]]["High_price"]:
        strategy.hour_bar_calculation[last_series["Insid"]]["High_price"] = float(last_series["High_price"])
    if float(last_series["Low_price"]) < strategy.hour_bar_calculation[last_series["Insid"]]["Low_price"]:
        strategy.hour_bar_calculation[last_series["Insid"]]["Low_price"] = float(last_series["Low_price"])
    strategy.hour_bar_calculation[last_series["Insid"]]["Vol"] += float(last_series["Vol"])
    strategy.hour_bar_calculation[last_series["Insid"]]["VolCcy"] += float(last_series["VolCcy"])
    strategy.hour_bar_calculation[last_series["Insid"]]["VolCcyQuote"] += float(last_series["VolCcyQuote"])
        
    if str(time.localtime(float(last_series["TS_open"])/1000).tm_min) == end_min:
        if (last_series["TS_open"]-strategy.hour_bar_calculation[last_series["Insid"]]["time_start"])/1000/60 + 1 == length:
            tempbar = barinfo()
            tempbar.InsID = last_series["Insid"]
            tempbar.TS_open = strategy.hour_bar_calculation[last_series["Insid"]]["time_start"]
            tempbar.Open_price = strategy.hour_bar_calculation[last_series["Insid"]]["Open_price"]
            tempbar.High_price = strategy.hour_bar_calculation[last_series["Insid"]]["High_price"]
            tempbar.Low_price = strategy.hour_bar_calculation[last_series["Insid"]]["Low_price"]
            tempbar.Close_price = strategy.hour_bar_calculation[last_series["Insid"]]["Close_price"]
            tempbar.Vol = strategy.hour_bar_calculation[last_series["Insid"]]["Vol"]
            tempbar.VolCcy = strategy.hour_bar_calculation[last_series["Insid"]]["VolCcy"]
            tempbar.VolCcyQuote = strategy.hour_bar_calculation[last_series["Insid"]]["VolCcyQuote"]
            strategy.hour_bar_calculation[last_series["Insid"]] = {"time_start":0,"Open_price":0,"High_price":0,"Low_price":float('+inf'),"Close_price":0,"Vol":0,"VolCcy":0,"VolCcyQuote":0}
            strategy.GenHourBarCustom(tempbar)
        elif (last_series["TS_open"]-strategy.hour_bar_calculation[last_series["Insid"]]["time_start"])/1000/60 + 1 < length:
            strategy.hour_bar_calculation[last_series["Insid"]] = {"time_start":0,"Open_price":0,"High_price":0,"Low_price":float('+inf'),"Close_price":0,"Vol":0,"VolCcy":0,"VolCcyQuote":0}

    
# 调用此方法的strategy必须包含UpdateBarCustom方法
def Load1MBarFromLocalMysql(strategy,user,password,database,Insid,length=0):
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
    temp_df = pd.read_csv("path")
    info_matrix = temp_df[["Insid","Ts_open","Open_price","High_price","Low_price","Close_price","Vol","VolCcy","VolCcyQuote"]].drop_duplicates(subset=["Ts_open"], keep="first").sort_values(by="Ts_open",ascending=True).values
    if length == 0:
        length = len(info_matrix)
    info_matrix = info_matrix[-length:]
    for i in range(length):
        strategy.UpdateBarCustom(info_matrix[i])
    
# 调用此方法用于更新本地id值
def UpdateOrderId(temp):
    temp[0] += 1  
    return "0" + str(temp[0])

