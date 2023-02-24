# encoding: UTF-8

"""
NET交易策略
"""

from datetime import time, datetime
import vnpy.trader.app
from vnpy.trader.app.ctaStrategy.ctaTemplate import CtaTemplate, BarGenerator, ArrayManager
from mailAlert import send_gmail
from vnpy.trader.language.chinese.constant import *
import talib
import math
import time
from pymongo import MongoClient
import numpy as np
import csv

import sys

reload(sys)
sys.setdefaultencoding("utf-8")


########################################################################
class StrategyChannelPort(CtaTemplate):
    """DualThrust交易策略"""
    className = 'StrategyChannelPort'
    author = u'Leonhardt'

    # 策略参数

    status = False

    # 仓位变量
    position = 0.0
    amount = 0.0
    orderTradedDict = {}

    openPrice = None

    longEntered = False
    shortEntered = False

    # 参数列表，保存了参数的名称
    paramList = ['name',
                 'className',
                 'author',
                 'vtSymbol']

    # 变量列表，保存了变量的名称
    varList = ['inited',
               'trading',
               'position',
               'max_amount',
               'openPrice',
               'maxProfit'
               ]

    # 同步列表，保存了需要保存到数据库的变量名称
    syncList = ['position', 'OrderPosition', 'openPrice', 'maxProfit']

    # ----------------------------------------------------------------------
    def __init__(self, ctaEngine, setting):
        """Constructor"""
        super(StrategyChannelPort, self).__init__(ctaEngine, setting)
        self.bg = BarGenerator(self.onBar, 60, self.onHourBar)
        self.record = setting["record"]
        self.period = setting["period"]
        self.stdup = setting["stdup"]
        self.stddown = setting["stddown"]
        self.length = setting["length"]
        self.lengthS = setting["lengthS"]
        datalength = setting["dataLength"]
        self.initDays = setting["initDays"]
        self.am = ArrayManager(datalength, 1)
        self.Threshold = setting["Threshold"]
        self.stopLoss = setting["stopLoss"]
        self.stopProfit = setting["stopProfit"]
        self.ratio = setting["ratio"]

        self.max_amount = setting["max_amount"]
        self.num = setting["num"]
        self.OrderLimit = setting["OrderLimit"]
        self.OrderAmountLimit = setting["OrderAmountLimit"]
        self.Ordertips = 0
        self.OrderAmountTips = 0
        self.alert = True

        self.upper = np.array([])
        self.lower = np.array([])
        self.ma = []
        self.tp = []

        self.orderList = []
        self.orderMonitor = []
        self.orderHour = []
        self.stopOrder = []

        self.orderTime = 0
        self.maxProfit = None
        self.OrderPosition = None
        self.action = False
        # 创建K线合成器对象

    # ----------------------------------------------------------------------
    def onInit(self):
        """初始化策略（必须由用户继承实现）"""

        # 载入历史数据，并采用回放计算的方式初始化策略数值

        # self.updateFlag()
        self.writeCtaLog(u'%s策略初始化' % self.name)
        initData = self.loadHourBar(self.initDays)
        for bar in initData:
            self.onHourBar(bar)

        self.putEvent()

    # ----------------------------------------------------------------------
    def onStart(self):
        """启动策略（必须由用户继承实现）"""
        self.writeCtaLog(u'%s策略启动' % self.name)
        self.trading = True
        self.putEvent()

    # ----------------------------------------------------------------------
    def onStop(self):
        """停止策略（必须由用户继承实现）"""
        self.writeCtaLog(u'%s策略停止' % self.name)
        self.trading = False
        self.putEvent()

    # ----------------------------------------------------------------------
    def onAccount(self, account):
        pass

    # ----------------------------------------------------------------------
    def onTick(self, tick):
        """收到行情TICK推送（必须由用户继承实现）"""

        if not self.trading:
            return

        if (self.Ordertips > self.OrderLimit or self.OrderAmountTips > self.OrderAmountLimit) and self.alert:
            self.bg.updateTick(tick)
            text = self.name + "_ORDERISSUE"
            send_gmail("15940402405@163.com", "Emergency", text)
            self.alert = False
            for orderID in self.orderList:
                if orderID != []:
                    self.cancelOrder(orderID[0])
            self.orderList = []
            return

        if not self.alert:
            self.bg.updateTick(tick)
            return

        if time.time() - self.orderTime > 5:
            for orderID in self.orderList:
                if orderID != []:
                    self.cancelOrder(orderID[0])
            self.orderList = []

            self.action = False
            if self.stopOrder:
                buyprice = tick.lastPrice * 1.005
                sellprice = tick.lastPrice * 0.995
                for order_Info in self.stopOrder:
                    if order_Info[0] == "buy":
                        if tick.lastPrice > order_Info[1]:
                            self.OrderPosition = self.max_amount
                            print("buy", tick.lastPrice, self.position)
                            if self.position == 0:
                                self.maxProfit = 0
                                self.status = False
                                res = self.buy(buyprice, self.max_amount)
                                if res:
                                    print('open buy', buyprice, self.max_amount)
                                    self.orderList.append(res)
                                    self.orderTime = time.time()
                                    self.action = True
                                self.Ordertips += 1
                                self.OrderAmountTips += self.max_amount

                            elif self.position < 0:
                                self.maxProfit = 0
                                self.status = False
                                res = self.cover(buyprice, abs(self.position))
                                if res:
                                    print('close buy', buyprice, abs(self.position))
                                    self.orderList.append(res)
                                    self.orderTime = time.time()
                                    self.action = True
                                res = self.buy(buyprice, self.max_amount)
                                if res:
                                    print('open buy', buyprice, self.max_amount)
                                    self.orderList.append(res)
                                    self.orderTime = time.time()
                                    self.action = True
                                self.Ordertips += 2
                                self.OrderAmountTips += self.max_amount + abs(self.position)

                    if order_Info[0] == "short":
                        if tick.lastPrice < order_Info[1]:
                            print("short", tick.lastPrice, self.position)
                            self.OrderPosition = -self.max_amount
                            if self.position == 0:
                                self.maxProfit = 0
                                self.status = False
                                res = self.short(sellprice, self.max_amount)
                                if res:
                                    print('open short', sellprice, self.max_amount)
                                    self.orderList.append(res)
                                    self.orderTime = time.time()
                                    self.action = True
                                self.Ordertips += 1
                                self.OrderAmountTips += self.max_amount
                            elif self.position > 0:
                                self.maxProfit = 0
                                self.status = False
                                res = self.sell(sellprice, self.position)
                                if res:
                                    print('close short', sellprice, self.position)
                                    self.orderList.append(res)
                                    self.orderTime = time.time()
                                    self.action = True
                                res = self.short(sellprice, self.max_amount)
                                if res:
                                    print('open short', sellprice, self.max_amount)
                                    self.orderList.append(res)
                                    self.orderTime = time.time()
                                    self.action = True
                                self.Ordertips += 2
                                self.OrderAmountTips += self.max_amount + abs(self.position)

            if self.openPrice and not self.action:
                buyprice = tick.lastPrice * 1.005
                sellprice = tick.lastPrice * 0.995
                if self.position > 0:
                    profit = (tick.lastPrice - self.openPrice) / self.openPrice
                    if not self.maxProfit:
                        self.maxProfit = profit
                    else:
                        self.maxProfit = max(self.maxProfit, profit)

                    if self.maxProfit > self.stopProfit:
                        self.status = True

                    if profit < self.stopLoss:
                        res = self.sell(sellprice, self.position)
                        print('stoploss short', sellprice, self.position)
                        self.Ordertips += 1
                        self.OrderAmountTips += abs(self.position)
                        self.status = False
                        self.maxProfit = 0
                        self.OrderPosition = 0
                        self.orderTime = time.time()
                        if res:
                            self.orderList.append(res)
                    else:
                        if self.status and profit < self.maxProfit * self.ratio:
                            self.OrderPosition = 0
                            self.status = False
                            self.maxProfit = 0
                            res = self.sell(sellprice, self.position)
                            print('stopprofit short', sellprice, self.position)
                            self.orderTime = time.time()
                            if res:
                                self.orderList.append(res)
                            self.Ordertips += 1
                            self.OrderAmountTips += abs(self.position)

                elif self.position < 0:
                    profit = (-tick.lastPrice + self.openPrice) / self.openPrice
                    if not self.maxProfit:
                        self.maxProfit = profit
                    else:
                        self.maxProfit = max(self.maxProfit, profit)

                    if self.maxProfit > self.stopProfit:
                        self.status = True

                    if profit < self.stopLoss:
                        res = self.cover(buyprice, -self.position)
                        print('stoploss buy', buyprice, -self.position)
                        self.OrderPosition = 0
                        self.orderTime = time.time()
                        self.status = False
                        self.maxProfit = 0
                        if res:
                            self.orderList.append(res)
                        self.Ordertips += 1
                        self.OrderAmountTips += abs(self.position)
                    else:
                        if self.status and profit < self.maxProfit * self.ratio:
                            self.OrderPosition = 0
                            self.status = False
                            self.maxProfit = 0
                            res = self.cover(buyprice, -self.position)
                            print('stopprofit buy', buyprice, -self.position)
                            self.orderTime = time.time()
                            if res:
                                self.orderList.append(res)
                            self.Ordertips += 1
                            self.OrderAmountTips += abs(self.position)

            if self.position == 0:
                self.maxProfit = None

        self.bg.updateTick(tick)

    # -----------------------------------------------------------------------
    def onBar(self, bar):

        if not self.trading:
            return

        if self.position == 0:
            self.maxProfit = None

        for orderID in self.orderHour:
            if orderID != []:
                self.cancelOrder(orderID[0])
        self.orderHour = []

        if self.Ordertips > self.OrderLimit or self.OrderAmountTips > self.OrderAmountLimit:
            self.bg.updateBar(bar)
            return

        if abs(self.position) > self.max_amount:
            overLoad = abs(self.position) - self.max_amount
            if self.position > 0:
                res = self.sell(bar.close * 0.99, overLoad)
                if res:
                    self.orderHour.append(res)
                self.Ordertips += 1
                self.OrderAmountTips += abs(self.position)

            if self.position < 0:
                res = self.cover(bar.close * 1.01, overLoad)
                if res:
                    self.orderHour.append(res)
                self.Ordertips += 1
                self.OrderAmountTips += abs(self.position)

        if time.time() - self.orderTime > 2:
            if self.OrderPosition != None:
                buyPrice = bar.close * 1.01
                sellPrice = bar.close * 0.99
                if self.position != self.OrderPosition:
                    if self.OrderPosition > 0:
                        if self.position >= 0 and self.OrderPosition - self.position > 0:
                            res = self.buy(buyPrice, self.OrderPosition - self.position)
                            if res:
                                self.orderHour.append(res)
                            self.Ordertips += 1
                            self.OrderAmountTips += abs(self.OrderPosition - self.position)
                        elif self.position < 0:
                            res = self.cover(buyPrice, -self.position)
                            if res:
                                self.orderHour.append(res)
                            res = self.buy(buyPrice, self.OrderPosition)
                            if res:
                                self.orderHour.append(res)
                            self.Ordertips += 2
                            self.OrderAmountTips += abs(self.position) + abs(self.OrderPosition)

                    elif self.OrderPosition < 0:
                        if self.position <= 0 and -self.OrderPosition + self.position > 0:
                            res = self.short(sellPrice, -self.OrderPosition + self.position)
                            if res:
                                self.orderHour.append(res)
                            self.Ordertips += 1
                            self.OrderAmountTips += abs(self.OrderPosition - self.position)
                        elif self.position > 0:
                            res = self.sell(sellPrice, self.position)
                            if res:
                                self.orderHour.append(res)
                            res = self.short(sellPrice, -self.OrderPosition)
                            if res:
                                self.orderHour.append(res)
                            self.Ordertips += 2
                            self.OrderAmountTips += abs(self.position) + abs(self.OrderPosition)

                    elif self.OrderPosition == 0:
                        if self.position < 0:
                            res = self.cover(buyPrice, -self.position)
                            if res:
                                self.orderHour.append(res)
                            self.Ordertips += 1
                            self.OrderAmountTips += abs(self.position)
                        elif self.position > 0:
                            res = self.sell(sellPrice, self.position)
                            if res:
                                self.orderHour.append(res)
                            self.Ordertips += 1
                            self.OrderAmountTips += abs(self.position)

        self.bg.updateBar(bar)

    # -----------------------------------------------------------------------
    def onHourBar(self, bar):

        self.am.updateBar(bar)

        if not self.am.inited:
            return

        for orderID in self.orderHour:
            if orderID != []:
                self.cancelOrder(orderID[0])
        self.orderHour = []
        self.stopOrder = []

        # reset
        print(self.name, self.Ordertips, self.OrderAmountTips, self.alert)
        self.Ordertips = 0
        self.OrderAmountTips = 0
        self.alert = True

        self.upper, self.mid, self.lower = talib.BBANDS(self.am.close,
                                                        timeperiod=self.period,
                                                        nbdevup=self.stdup,
                                                        nbdevdn=self.stddown,
                                                        matype=0)

        close = self.am.close
        ma = np.mean(close[-self.period:])
        self.ma.append(ma)
        if len(self.ma) > self.length:
            tp = (ma / self.ma[0] - 1) * 1000
            self.tp.append(tp)
            self.ma.pop(0)
        else:
            return

        if len(self.tp) == self.lengthS:
            tpma = sum(self.tp) / self.lengthS
            Point = abs(tpma - tp)
            self.tp.pop(0)
        else:
            return

        print([bar.datetime, float(self.am.close[-1]), float(self.upper[-1]), float(self.mid[-1]), float(self.lower[-1])
            , float(Point)])

        if not self.trading:
            return

        condition1 = float(close[-1]) < float(self.lower[-1])
        condition2 = float(close[-1]) > float(self.upper[-1])
        condition3 = float(Point) < float(self.Threshold)
        condition4 = float(Point) > float(self.Threshold)

        buyP = float(close[-1]) * 1.01
        sellP = float(close[-1]) * 0.99

        if condition1 and condition3:
            if self.position <= 0:
                self.stopOrder.append(("buy", float(self.lower[-1])))
                print("buyS", float(self.lower[-1]))

        elif condition2 and condition3:
            if self.position >= 0:
                self.stopOrder.append(("short", float(self.upper[-1])))
                print("shortS", float(self.upper[-1]))

        elif condition2 and condition4:
            self.OrderPosition = self.max_amount
            if self.position == 0:
                self.maxProfit = 0
                self.status = False
                res = self.buy(buyP, self.max_amount)
                if res:
                    self.orderTime = time.time()
                    self.orderHour.append(res)
                self.Ordertips += 1
                self.OrderAmountTips += abs(self.max_amount)

            elif self.position < 0:
                self.maxProfit = 0
                self.status = False
                res = self.cover(buyP, abs(self.position))
                if res:
                    self.orderHour.append(res)
                    self.orderTime = time.time()
                res = self.buy(buyP, self.max_amount)
                if res:
                    self.orderHour.append(res)
                    self.orderTime = time.time()
                self.Ordertips += 2
                self.OrderAmountTips += abs(self.max_amount) + abs(self.position)

        elif condition1 and condition4:
            self.OrderPosition = -self.max_amount
            if self.position == 0:
                self.maxProfit = 0
                self.status = False
                res = self.short(sellP, self.max_amount)
                if res:
                    self.orderHour.append(res)
                self.Ordertips += 1
                self.OrderAmountTips += abs(self.max_amount)

            elif self.position > 0:
                self.maxProfit = 0
                self.status = False
                res = self.sell(sellP, self.position)
                if res:
                    self.orderHour.append(res)
                    self.orderTime = time.time()
                res = self.short(sellP, self.max_amount)
                if res:
                    self.orderHour.append(res)
                    self.orderTime = time.time()
                self.Ordertips += 2
                self.OrderAmountTips += abs(self.max_amount) + abs(self.position)

        try:
            myclient = MongoClient("mongodb://localhost:27017/")
            dbName = "Strategy_Var"
            db_min_bar = myclient[dbName]
            col = db_min_bar[self.className + "_" + str(self.vtSymbol) + "_" + str(self.num)]
            a = {}
            datetime_ = bar.datetime
            a["datetime"] = datetime_
            a["upper"] = self.upper[-1]
            a["lower"] = self.lower[-1]
            a["mid"] = self.mid[-1]
            a["tp"] = tp
            a["position"] = self.position
            a["maxProfit"] = self.maxProfit
            col.insert(a)
            myclient.close()
        except Exception as e:
            print("error insert", e)

        # 同步数据到数据库
        self.saveSyncData()

        # 发出状态更新事件
        self.putEvent()

    # -----------------------------------------------------------------------
    def onOrder(self, order):
        """收到委托变化推送（必须由用户继承实现）"""
        """委托更新"""

        vn_symbol = "_".join(order.vtSymbol.split("-"))
        if vn_symbol != self.vtSymbol:
            return

        vtOrderID = order.vtOrderID
        vtSymbol = order.vtSymbol
        newTradedVolume = order.tradedVolume
        lastTradedVolume = self.orderTradedDict.get(vtOrderID, 0)

        direction = order.direction
        offset = order.offset

        if newTradedVolume > lastTradedVolume:
            self.orderTradedDict[vtOrderID] = newTradedVolume
            volume = newTradedVolume - lastTradedVolume

            if direction == DIRECTION_LONG:
                if self.position >= 0:
                    self.openPrice = order.avgprice
                else:
                    self.openPrice = None
                self.position += volume
            else:
                if self.position <= 0:
                    self.openPrice = order.avgprice
                else:
                    self.openPrice = None
                self.position -= volume

        # 同步数据到数据库
        self.saveSyncData()

        # 发出状态更新事件
        self.putEvent()

    # ----------------------------------------------------------------------
    def onTrade(self, trade):
        # 发出状态更新事件
        self.putEvent()

    # ----------------------------------------------------------------------
    def onStopOrder(self, so):
        """停止单推送"""
        pass

    # ----------------------------------------------------------------------


def round_range(x, y):
    if y > x:
        return (y - x) // 40.0 * 40 + x
    else:
        return math.ceil((y - x) / 40.0) * 40 + x


def ma(x, y):
    if len(x) >= y:
        return round(sum(x[-y:]) / y, 3)
    else:
        return None


def trade_signal(a, b, ma_len):
    diff = b - a
    diff_ma = ma(diff, ma_len)
    if not diff_ma:
        return None





