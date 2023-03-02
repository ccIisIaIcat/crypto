import pyserver
import threading
import grpc
import deliver_pb2
import deliver_pb2_grpc
import time
import tool
import json
import numpy as np
import pandas as pd
from collections.abc import Iterable


class strategy:
    # 交互参数
    porttick:str
    portbarcustom:str
    portaccount:str
    basic_conf:tool.config
    portsubmit:str
    portorder:str
    # 订单stud
    stub_order:deliver_pb2_grpc.OrerReceiverStub
    
    def __init__(self,conf:tool.config):
        self.basic_conf = conf
        if conf.tickPort != "":
            self.porttick = conf.tickPort
        if conf.barPort != "":
            self.portbarcustom = conf.barPort
        if conf.accountPort != "":
            self.portaccount = conf.accountPort
        self.portsubmit = conf.portsubmit
        self.portorder = conf.portorder
        channel = grpc.insecure_channel('localhost:'+self.portsubmit)
        stub = deliver_pb2_grpc.SubmitServerReceiverStub(channel)
        channel2 = grpc.insecure_channel('localhost:'+self.portorder)
        self.stub_order = deliver_pb2_grpc.OrerReceiverStub(channel2)
        response = stub.SubmitServerReceiver(conf.genLocalSubmit())
        print(response)
    
    # 声明存储对象
    tick_list = tool.TickinfoArray(Insid="ETH-USDT-SWAP",max_length=10)
    bar_list = tool.BarinfoArray(Insid="ETH-USDT-SWAP")
    bar_hour_list = tool.BarinfoArray(Insid="ETH-USDT-SWAP")
    
    
    def UpdateBarCustom(self,bar_info):
        if isinstance(bar_info,Iterable):
            self.bar_list.addnum(bar_info)
        else:
            self.bar_list.add(tool.barinfo(bar_info))
        tool.genhourbarCustom(self,self.bar_list,"59")
        # print("infogather barhour called",self.bar_list.df)
        
    def UpdateTick(self,tick_info):
        if isinstance(tick_info,Iterable):
            self.tick_list.addnum(tick_info)
        else:
            self.tick_list.add(tool.tickinfo(tick_info))
        print("infogather tick called",self.tick_list.df)
        
    def UpdateAccount(self,account_info):
        format_info = json.loads(account_info)
        print(format_info)
        if "arg" in format_info:
            # if format_info["arg"]["channel"] == "account":
            #     print("infogather account called",format_info["data"][0]["details"][0])
            # if format_info["arg"]["channel"] == "positions":
            #     print("infogather position called",format_info["data"][0])
            if format_info["arg"]["channel"] == "orders":
                print("oooooooooooooooooooooooorder")
                print(format_info["arg"])
                print("infogather orders called",format_info["data"][0])
        
    def GenHourBarCustom(self,bar_info):
        self.bar_hour_list.add(bar_info)
        print(self.bar_hour_list.df)
        
        
    def Makeorder(self,order_info:tool.ordertemplate):
        response = self.stub_order.OrerRReceiver(order_info.genOrder())
        print(response.response_me)
        return response
        
    def Start(self):
        # print(self.porttick,self.portbarcustom)
        tick_thread = threading.Thread(target=pyserver.serveTick,args={self})
        barhour_thread = threading.Thread(target=pyserver.serveBarCustom,args={self})
        account_thread = threading.Thread(target=pyserver.serveAccount,args={self})
        tick_thread.start()
        barhour_thread.start()
        account_thread.start()



if __name__ == '__main__':
    ############################################
    my_conf = tool.config()
    # 订阅对象.可选（tick,bar,account,position,order）
    my_conf.subtype = "order"
    # bar相关
    my_conf.barcustom = "1m"
    my_conf.barInsid = "ETH-USDT-SWAP"
    my_conf.barPort = "6001"
    # tick相关
    my_conf.tickInsid = "ETH-USDT-SWAP"
    my_conf.tickPort = "6002"
    # 账户相关
    my_conf.accountPort = "6003"
    # go端端口
    my_conf.portsubmit = "6101"
    my_conf.portorder = "6102"
    ############################################
    datagather = strategy(my_conf)
    # tool.Load1MBarFromLocalMysql(datagather,"root","zwj12345","crypto_swap","ETH-USDT-SWAP")
    odt = tool.ordertemplate()
    odt.insId = "ETH-USDT-SWAP"
    odt.posSide = "long"
    odt.tdMode = "cross"
    odt.side = "sell"
    odt.ordType = "market"
    odt.sz = "1"
    odt.clOrdId = "lalala558875"
    datagather.Start()
    time.sleep(10)
    tata = datagather.Makeorder(odt)
    print(tata)
    # time.sleep(10)
    # {"instId":"ETH-USDT-SWAP","posSide":"long","tdMode":"cross","side":"buy","ordType":"market","sz":"1"}
    # datagather.Makeorder(deliver_pb2.Order(insId="ETH-USDT-SWAP",posSide="long",tdMode="cross",side="buy",ordType="market",sz="1"))
    pyserver.NeverStop()