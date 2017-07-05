// Package oak is a game engine...
package oak

import (
	"fmt"
	"image"
	"sync"

	"bitbucket.org/oakmoundstudio/oak/dlog"
	"bitbucket.org/oakmoundstudio/oak/event"

	"golang.org/x/exp/shiny/screen"
)

var (
	winBuffer     screen.Buffer
	screenControl screen.Screen
	windowControl screen.Window

	windowRect     image.Rectangle
	windowUpdateCh = make(chan bool)

	osCh = make(chan func())

	initControl = sync.Mutex{}

	lifecycleInit bool
)

func lifecycleLoop(s screen.Screen) {
<<<<<<< HEAD
	screenControl = s
	var err error
	fmt.Println("Lifecycle enter")

	// The world buffer represents the total space that is conceptualized by the engine
	// and able to be drawn to. Space outside of this area will appear as smeared
	// white (on windows).
	worldBuffer, err = screenControl.NewBuffer(image.Point{WorldWidth, WorldHeight})
	if err != nil {
		dlog.Error(err)
=======
	initControl.Lock()
	if lifecycleInit {
		dlog.Error("Started lifecycle twice, aborting second call")
>>>>>>> master
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
	changeWindow(ScreenWidth, ScreenHeight)
<<<<<<< HEAD
	defer windowControl.Release()

	eb = event.GetEventBus()

	//go KeyHoldLoop()
	//go InputLoop()

	// Initiate the first scene
	//initCh <- true

	if conf.ShowFPS {
		fmt.Println("Starting draw")
		go DrawLoopFPS()
	} else {
		fmt.Println("Starting draw")
		go DrawLoopNoFPS()
	}

	event.ResolvePending()
}

// do runs f on the osLocked thread.
func osLockedFunc(f func()) {
	done := make(chan bool, 1)
	osCh <- func() {
		f()
		done <- true
	}
	<-done
}

func LogicLoop() chan bool {
	// The logical loop.
	// In order, it waits on receiving a signal to begin a logical frame.
	// It then runs any functions bound to when a frame begins.
	// It then allows a scene to perform it's loop operation.
	ch := make(chan bool)
	go func(doneCh chan bool) {
		ticker := time.NewTicker(time.Second / time.Duration(int64(FrameRate)))
		for {
			select {
			case <-ticker.C:
				<-eb.TriggerBack("EnterFrame", nil)
				sceneCh <- true
			case <-doneCh:
				ticker.Stop()
				break
			}
		}
	}(ch)
	return ch
}

func GetScreen() *image.RGBA {
	return winBuffer.RGBA()
}

func GetWorld() *image.RGBA {
	return worldBuffer.RGBA()
}

func SetWorldSize(x, y int) {
	worldBuffer, _ = screenControl.NewBuffer(image.Point{x, y})
=======

	dlog.Info("Getting event bus")
	eb = event.GetBus()

	dlog.Info("Starting draw loop")
	go drawLoop()
	dlog.Info("Starting key hold loop")
	go keyHoldLoop()
	dlog.Info("Starting input loop")
	go inputLoop()

	dlog.Info("Starting event handler")
	go event.ResolvePending()
	// The quit channel represents a signal
	// for the engine to stop.
	<-quitCh
	return
>>>>>>> master
}

func changeWindow(width, height int) {
	// The window controller handles incoming hardware or platform events and
	// publishes image data to the screen.
	wC, err := windowController(screenControl, width, height)
	if err != nil {
		dlog.Error(err)
		panic(err)
	}
	windowControl = wC
	windowRect = image.Rect(0, 0, width, height)
}

// ChangeWindow sets the width and height of the game window. But it doesn't.
func ChangeWindow(width, height int) {
	windowRect = image.Rect(0, 0, width, height)
}

// GetScreen returns the current screen as an rgba buffer
func GetScreen() *image.RGBA {
	return winBuffer.RGBA()
}
