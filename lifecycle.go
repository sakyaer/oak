//+build !js

package oak

import (
	"image"
	"sync"

	"github.com/oakmound/oak/dlog"
	"github.com/oakmound/oak/event"

	"golang.org/x/exp/shiny/screen"
)

var (
	winBuffer     screen.Buffer
	screenControl screen.Screen
	windowControl screen.Window

	initControl = sync.Mutex{}

	lifecycleInit bool
)

func lifecycleLoop(s screen.Screen) {
	initControl.Lock()
	if lifecycleInit {
		dlog.Error("Started lifecycle twice, aborting second call")
		initControl.Unlock()
		return
	}
	lifecycleInit = true
	initControl.Unlock()
	dlog.Info("Init Lifecycle")

	screenControl = s
	var err error

	// The window buffer represents the subsection of the world which is available to
	// be shown in a window.
	dlog.Info("Creating window buffer")
	winBuffer, err = screenControl.NewBuffer(image.Point{ScreenWidth, ScreenHeight})
	if err != nil {
		dlog.Error(err)
		return
	}

	// The window controller handles incoming hardware or platform events and
	// publishes image data to the screen.\
	dlog.Info("Creating window controller")
	changeWindow(ScreenWidth*conf.Screen.Scale, ScreenHeight*conf.Screen.Scale)

	dlog.Info("Getting event bus")
	eb = event.GetBus()

	dlog.Info("Starting draw loop")
	go drawLoop()
	dlog.Info("Starting input loop")
	go inputLoop()

	dlog.Info("Starting event handler")
	go event.ResolvePending()
	// The quit channel represents a signal
	// for the engine to stop.
	// Lifecycle goroutine reaches here.
	<-quitCh
}
