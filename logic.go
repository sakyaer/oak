package oak

import (
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/timing"
)

var (
	logicHandler event.Handler = event.DefaultBus
)

// SetLogicHandler swaps the logic system of the engine with some other
// implementation. If this is never called, it will use event.DefaultBus
func SetLogicHandler(h event.Handler) {
	logicHandler = h
}

func logicTickerInit() *timing.DynamicTicker {
	LogicTicker = timing.NewDynamicTicker()
	LogicTicker.SetTick(timing.FPSToDuration(FrameRate))
	return LogicTicker
}

func logicLoopSingle(LogicTicker *timing.DynamicTicker) {
	select {
	case <-LogicTicker.C:
		<-eb.TriggerBack("EnterFrame", framesElapsed)
		framesElapsed++
	}
}
