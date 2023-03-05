import pyserver
import threading
import grpc
import deliver_pb2
import deliver_pb2_grpc
import datetime
import tool
import json
import numpy as np
import pandas as pd
from collections.abc import Iterable

# 针对该demo策略，设置的简易订单模板
# 开多
odt_buylong = tool.ordertemplate()
odt_buylong.insId = "ETH-USDT-SWAP"
odt_buylong.posSide = "long"
odt_buylong.tdMode = "cross"
odt_buylong.side = "buy"
odt_buylong.ordType = "market"
odt_buylong.sz = "1"
odt_buylong.clOrdId = ""
# 平多
odt_selllong = tool.ordertemplate()
odt_selllong.insId = "ETH-USDT-SWAP"
odt_selllong.posSide = "long"
odt_selllong.tdMode = "cross"
odt_selllong.side = "sell"
odt_selllong.ordType = "market"
odt_selllong.sz = "1"
odt_selllong.clOrdId = ""
# 开空
odt_buyshort = tool.ordertemplate()
odt_buyshort.insId = "ETH-USDT-SWAP"
odt_buyshort.posSide = "short"
odt_buyshort.tdMode = "cross"
odt_buyshort.side = "buy"
odt_buyshort.ordType = "market"
odt_buyshort.sz = "1"
odt_buyshort.clOrdId = ""
# 平空
odt_sellshort = tool.ordertemplate()
odt_sellshort.insId = "ETH-USDT-SWAP"
odt_sellshort.posSide = "short"
odt_sellshort.tdMode = "cross"
odt_sellshort.side = "sell"
odt_sellshort.ordType = "market"
odt_sellshort.sz = "1"
odt_sellshort.clOrdId = ""

# simple demo
# 只考虑涨跌趋势，当处于上涨趋势时，无开单开多或者平空仓开多，处于下降趋势时，无开单开空或平多单开孔
class strategy:
    # 交互参数
    porttick:str
    portbarcustom:str
    portaccount:str
    basic_conf:tool.config
    portsubmit:str
    portorder:str
    # 订单stub
    stub_order:deliver_pb2_grpc.OrerReceiverStub
    
    def __init__(self,conf:tool.config):
        # tool.Load1MBarFromLocalMysql(self,"root","zwj12345","crypto_swap","ETH-USDT-SWAP")
        self.trade_forbidden_signal = True
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
        self.stub_order = deliver_pb2_grpc.SubmitServerReceiverStub(channel2)
        response = stub.SubmitServerReceiver(conf.genLocalSubmit())
        print(response)
    
    # 声明存储对象
    tick_list = tool.TickinfoArray(Insid="ETH-USDT-SWAP",max_length=10)
    bar_list = tool.BarinfoArray(Insid="ETH-USDT-SWAP")
    bar_hour_list = tool.BarinfoArray(Insid="ETH-USDT-SWAP")
    
    StrategyName = "sA" # 策略名称，用于标注订单
    OrderNumber = 0 # 累加订单id，每次发送订单时OrderNumber+1
    
    # 策略对象
    MA_hour = [] # 小时bar的MA均线
    Std_hour = [] # 小时bar对应MA的std
    MA_grad = [] # MA梯度（该梯度为当前MA和之前MA_gap个MA的梯度，结果乘以1000）
    MA_grad_mean = [] # MA的梯度的平均
    basic_signal = [] # 信号指标1：（MA梯度和梯度平均的差值的绝对值）up代表大于阈值，below代表小于阈值
    
    bolling_up = [] # 当前布林带的上方
    bolling_down = [] # 当前布林带的下方
    
    # 策略参数
    MA_length = 20 # MA均线长度
    MA_gap = 20 # 计算梯度时的间隔长度
    MA_grade_mean_length = 20 # MA均线移动平均长度
    basic_signal_threshold = 40 # 信号指标阈值
    stop_lose_ratio = 0.03 # 止损比例
    bolling_std_k = 2 # 布林带方差个数
    
    # 持仓信息（简化起见当前策略只有一笔持仓）
    position_record = [] # 仓位更改建议在UpdateAccount模块进行，确保订单完成成交
    # (当前简化仓位只保存订单clOrdId,开仓方向posSide,成交均价avgPx,成交时间fillTime,资金费率fee)
    order_record = {} # 未回执报单记录，只有当报单个数为0时，将trade_forbidden_signal标记为False
    trade_forbidden_signal = False # 禁止交易信号
    
    
    def UpdateBarCustom(self,bar_info):
        # 更新本地bar列表
        if isinstance(bar_info,Iterable):
            self.bar_list.addnum(bar_info)
        else:
            self.bar_list.add(tool.barinfo(bar_info))
        # 声明自定义bar方法，声明该方法时策略结构体必须包含GenHourBarCustom的类内函数
        tool.genhourbarCustom(self,self.bar_list,"59")

        # print("infogather barhour called",self.bar_list.df)
        
    def UpdateTick(self,tick_info):
        # 更新本地tick列表
        if isinstance(tick_info,Iterable):
            self.tick_list.addnum(tick_info)
        else:
            self.tick_list.add(tool.tickinfo(tick_info))
        # 只有当basic信号存在且小于阈值时考虑tick开仓平仓
        if len(self.basic_signal) > 0 and self.basic_signal[-1]<self.basic_signal_threshold:
            # 若当前状况可交易
            if not self.trade_forbidden_signal: 
                # 当tick的ask小于下方布林带时,做多
                if list(self.tick_list.df["Ask1_price"])[-1] < self.bolling_down[-1]:
                    # 如果没有仓位，开仓
                    if len(self.position_record) == 0:
                        odt_buylong.clOrdId = self.StrategyName + str(self.OrderNumber)
                        self.OrderNumber += 1
                        self.order_record[odt_buylong.clOrdId] = 1
                        self.trade_forbidden_signal = True
                        res = self.Makeorder(odt_buylong)
                        if res == "1":
                            print("发单失败")
                            del self.order_record[odt_buylong.clOrdId]
                            self.trade_forbidden_signal = False
                    # 如果持有空仓，平空，开多
                    elif self.position_record[0]["posSide"] == "short":
                        odt_sellshort.clOrdId = self.StrategyName + str(self.OrderNumber)
                        self.OrderNumber += 1
                        self.order_record[odt_sellshort.clOrdId] = 1
                        self.trade_forbidden_signal = True
                        res = self.Makeorder(odt_sellshort)
                        if res == "1":
                            print("发单失败")
                            del self.order_record[odt_sellshort.clOrdId]
                            self.trade_forbidden_signal = False
                        else:
                            odt_buylong.clOrdId = self.StrategyName + str(self.OrderNumber)
                            self.OrderNumber += 1
                            self.order_record[odt_buylong.clOrdId] = 1
                            self.trade_forbidden_signal = True
                            res = self.Makeorder(odt_buylong)
                            if res == "1":
                                print("发单失败")
                                del self.order_record[odt_buylong.clOrdId]
                                self.trade_forbidden_signal = False
                # 当tick的bid大于上方布林带时,做空
                if list(self.tick_list.df)["Bid1_price"][-1] > self.bolling_up[-1]:
                    # 如果没有仓位，开空
                    if len(self.position_record) == 0:
                        odt_buyshort.clOrdId = self.StrategyName + str(self.OrderNumber)
                        self.OrderNumber += 1
                        self.order_record[odt_buyshort.clOrdId] = 1
                        self.trade_forbidden_signal = True
                        res = self.Makeorder(odt_buyshort)
                        if res == "1":
                            print("发单失败")
                            del self.order_record[odt_buyshort.clOrdId]
                            self.trade_forbidden_signal = False
                    # 如果持有空仓，平多，开空
                    elif self.position_record[0]["posSide"] == "long":
                        odt_selllong.clOrdId = self.StrategyName + str(self.OrderNumber)
                        self.OrderNumber += 1
                        self.order_record[odt_selllong.clOrdId] = 1
                        self.trade_forbidden_signal = True
                        res = self.Makeorder(odt_selllong)
                        if res == "1":
                            print("发单失败")
                            del self.order_record[odt_selllong.clOrdId]
                            self.trade_forbidden_signal = False
                        else:
                            odt_buyshort.clOrdId = self.StrategyName + str(self.OrderNumber)
                            self.OrderNumber += 1
                            self.order_record[odt_buyshort.clOrdId] = 1
                            self.trade_forbidden_signal = True
                            res = self.Makeorder(odt_buyshort)
                            if res == "1":
                                print("发单失败")
                                del self.order_record[odt_buyshort.clOrdId]
                                self.trade_forbidden_signal = False
        # 止损检查
        if len(self.position_record) > 0:
            if self.position_record[0]["posSide"] == "long":
                if list(self.tick_list.df["Bid1_price"])[-1]/self.position_record[0]["avgPx"]-1 < -self.stop_lose_ratio:
                    # 止损平多
                    odt_selllong.clOrdId = self.StrategyName + str(self.OrderNumber)
                    self.OrderNumber += 1
                    self.order_record[odt_selllong.clOrdId] = 1
                    self.trade_forbidden_signal = True
                    res = self.Makeorder(odt_selllong)
                    if res == "1":
                        print("发单失败")
                        del self.order_record[odt_selllong.clOrdId]
                        self.trade_forbidden_signal = False
            if self.position_record[0]["posSide"] == "short":
                if list(self.tick_list.df["Bid1_price"])[-1]/self.position_record[0]["avgPx"]-1 > self.stop_lose_ratio:
                    # 止损平空
                    odt_sellshort.clOrdId = self.StrategyName + str(self.OrderNumber)
                    self.OrderNumber += 1
                    self.order_record[odt_sellshort.clOrdId] = 1
                    self.trade_forbidden_signal = True
                    res = self.Makeorder(odt_sellshort)
                    if res == "1":
                        print("发单失败")
                        del self.order_record[odt_sellshort.clOrdId]
                        self.trade_forbidden_signal = False
            
        # print("infogather tick called",self.tick_list.df)
        
    def GenHourBarCustom(self,bar_info):
        self.bar_hour_list.add(bar_info)
        # 小时bar个数大于等于MA均线要求长度，生成MA
        if self.bar_hour_list.getlength() >= self.MA_length:
            temp_list = list(self.bar_hour_list.df["Close_price"])[-self.MA_length:]
            self.MA_hour.append(np.mean(temp_list))
            std = np.std(temp_list,ddof=1)
            self.bolling_down.append(self.MA_hour[-1] - self.bolling_std_k*std)
            self.bolling_up.append(self.MA_hour[-1] + self.bolling_std_k*std)
        # MA个数大于等于MA_gap时,生成MA均线梯度值
        if len(self.MA_hour) > self.MA_gap:
            self.MA_grad.append((self.MA_hour[-1]/self.MA_hour[-self.MA_gap-1]-1)*1000)
        # 当MA梯度个数大于要求的梯度长度，生成MA梯度均值
        if len(self.MA_grad) >= self.MA_grade_mean_length:
            self.MA_grad_mean.append(np.mean(self.MA_grad[-self.MA_grade_mean_length:]))
            # 当产生MA梯度平均时，计算信号指标
            self.basic_signal.append(np.abs(self.MA_grad_mean[-1]-self.MA_grad[-1]))
            print(self.basic_signal[-1])
        # 当basic信号存在且大于阈值时考虑bar交易
        if len(self.basic_signal) > 0 and self.basic_signal[-1]>=self.basic_signal_threshold:
            # 若当前状况可交易
            if not self.trade_forbidden_signal:
                # 当当前bar的close大于上bolling up时，做多
                if self.bar_hour_list.df["Close_price"] > self.bolling_up[-1]:
                    # 如果没有仓位，开仓
                    if len(self.position_record) == 0:
                        odt_buylong.clOrdId = self.StrategyName + str(self.OrderNumber)
                        self.OrderNumber += 1
                        self.order_record[odt_buylong.clOrdId] = 1
                        self.trade_forbidden_signal = True
                        res = self.Makeorder(odt_buylong)
                        if res == "1":
                            print("发单失败")
                            del self.order_record[odt_buylong.clOrdId]
                            self.trade_forbidden_signal = False
                    # 如果持有空仓，平空，开多
                    elif self.position_record[0]["posSide"] == "short":
                        odt_sellshort.clOrdId = self.StrategyName + str(self.OrderNumber)
                        self.OrderNumber += 1
                        self.order_record[odt_sellshort.clOrdId] = 1
                        self.trade_forbidden_signal = True
                        res = self.Makeorder(odt_sellshort)
                        if res == "1":
                            print("发单失败")
                            del self.order_record[odt_sellshort.clOrdId]
                            self.trade_forbidden_signal = False
                        else:
                            odt_buylong.clOrdId = self.StrategyName + str(self.OrderNumber)
                            self.OrderNumber += 1
                            self.order_record[odt_buylong.clOrdId] = 1
                            self.trade_forbidden_signal = True
                            res = self.Makeorder(odt_buylong)
                            if res == "1":
                                print("发单失败")
                                del self.order_record[odt_buylong.clOrdId]
                                self.trade_forbidden_signal = False
            # 当当前bar的close小于bolling down时，做空
                if self.bar_hour_list.df["Close_price"] < self.bolling_down[-1]:
                     # 如果没有仓位，开空
                    if len(self.position_record) == 0:
                        odt_buyshort.clOrdId = self.StrategyName + str(self.OrderNumber)
                        self.OrderNumber += 1
                        self.order_record[odt_buyshort.clOrdId] = 1
                        self.trade_forbidden_signal = True
                        res = self.Makeorder(odt_buyshort)
                        if res == "1":
                            print("发单失败")
                            del self.order_record[odt_buyshort.clOrdId]
                            self.trade_forbidden_signal = False
                    # 如果持有空仓，平多，开空
                    elif self.position_record[0]["posSide"] == "long":
                        odt_selllong.clOrdId = self.StrategyName + str(self.OrderNumber)
                        self.OrderNumber += 1
                        self.order_record[odt_selllong.clOrdId] = 1
                        self.trade_forbidden_signal = True
                        res = self.Makeorder(odt_selllong)
                        if res == "1":
                            print("发单失败")
                            del self.order_record[odt_selllong.clOrdId]
                            self.trade_forbidden_signal = False
                        else:
                            odt_buyshort.clOrdId = self.StrategyName + str(self.OrderNumber)
                            self.OrderNumber += 1
                            self.order_record[odt_buyshort.clOrdId] = 1
                            self.trade_forbidden_signal = True
                            res = self.Makeorder(odt_buyshort)
                            if res == "1":
                                print("发单失败")
                                del self.order_record[odt_buyshort.clOrdId]
                                self.trade_forbidden_signal = False
            
        # print(self.bar_hour_list.df)
        
    def UpdateAccount(self,account_info):
        format_info = json.loads(account_info)
        print(format_info)
        if "arg" in format_info:
            if format_info["arg"]["channel"] == "account":
                print("infogather account called",format_info["data"][0]["details"][0])
            if format_info["arg"]["channel"] == "positions":
                print("infogather position called")
            if format_info["arg"]["channel"] == "orders":
                # 处理订单
                for order_respon in format_info["data"]:
                    if order_respon["side"] == "buy":
                        # 开仓成功记录持仓
                        if order_respon["clOrdId"] in self.order_record:
                            temp_position = {"clOrdId":order_respon["clOrdId"],"posSide":order_respon["posSide"],"avgPx":order_respon["avgPx"],"fillTime":order_respon["fillTime"],"fee":order_respon["fee"]}
                            self.position_record.append(temp_position)
                            del self.order_record[order_respon["clOrdId"]]
                    else:
                        # 平仓成功删除持仓
                        if order_respon["clOrdId"] in self.order_record:
                            del temp_position[0]
                            del self.order_record[order_respon["clOrdId"]]
                if len(self.order_record) == 0:
                    self.trade_forbidden_signal = False
        
    def Makeorder(self,order_info:tool.ordertemplate):
        response = self.stub_order.OrerRReceiver(order_info.genOrder())
        res = json.loads(response.response_me)["code"]
        return res
        
    def Start(self):
        self.trade_forbidden_signal = False
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
    my_conf.subtype = "account"
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
    
    # datagather.Start()
    # time.sleep(10)
    # {"instId":"ETH-USDT-SWAP","posSide":"long","tdMode":"cross","side":"buy","ordType":"market","sz":"1"}
    # datagather.Makeorder(deliver_pb2.Order(insId="ETH-USDT-SWAP",posSide="long",tdMode="cross",side="buy",ordType="market",sz="1"))
    pyserver.NeverStop()