package deliver_server

import (
	"global"
	"log"
	"time"
)

type StrategyUnit struct {
	// Public
	BarDeliver      *Bar_deliver
	TickDeliver     *Tick_deliver
	PingPongChan    chan bool
	Submitinfo      global.SubmitInfo
	AccountInfoChan chan []byte
	// Private
	strategyName   string
	timeout_second int
}

func GenStrategyUnit(strategy_name string, timeout_second int, subinfo global.SubmitInfo, PingPongChan chan bool) *StrategyUnit {
	su := &StrategyUnit{}
	su.PingPongChan = PingPongChan
	su.strategyName = strategy_name
	su.timeout_second = timeout_second
	su.Submitinfo = subinfo
	su.AccountInfoChan = make(chan []byte, 100)
	return su
}

func (S *StrategyUnit) Start() {
	S.initDeliver()
	if S.Submitinfo.Bar.Judge {
		go S.BarDeliver.DeliverBar()
	}
	if S.Submitinfo.Tick.Judge {
		go S.TickDeliver.DeliverTick()
	}
	S.pingPong()
}

func (S *StrategyUnit) Close() {
	if S.Submitinfo.Bar.Judge {
		S.BarDeliver.Signal = true
	}
	if S.Submitinfo.Tick.Judge {
		S.TickDeliver.Signal = true
	}
}

func (S *StrategyUnit) initDeliver() {
	if S.Submitinfo.Bar.Judge {
		S.BarDeliver = GenBarDeliver(S.Submitinfo.Bar.InsList, S.Submitinfo.Bar.Port, S.Submitinfo.Bar.Custom_type)
	}
	if S.Submitinfo.Tick.Judge {
		S.TickDeliver = GenTickDeliver(S.Submitinfo.Tick.InsList, S.Submitinfo.Tick.Port)
	}
}

func (S *StrategyUnit) pingPong() {
	temp_signal := true
	for temp_signal {
		select {
		case <-S.PingPongChan:
		case <-time.After(time.Second * time.Duration(S.timeout_second)):
			S.Close()
			log.Println("strategy:", S.strategyName, " pingpong timeout! disconnected")
			temp_signal = false
		}
	}
}
