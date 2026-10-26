//go:build i_tdisplays3

package input

import (
	"machine"
	"sync/atomic"
	"time"
	_ "unsafe"

	"github.com/axellse/cyclos/cyclos/lib/libinput"
	"github.com/axellse/cyclos/cyclos/shared"
)

type TDisplayS3Input struct {}

func Init() TDisplayS3Input {
	return TDisplayS3Input{}
}

const (
    evUp      = 1
    evDown    = 2
    evConfirm = 3

    debounceNs = 20 * 1000 * 1000 // 20ms
    ringSize   = 16
)

var (
    ring       [ringSize]uint32
    ringHead   uint32
    ringTail   uint32
    upSeen     bool
    downSeen   bool
    lastEdgeUp int64
    lastEdgeDn int64
    UpButton = machine.GPIO14
    DownButton = machine.GPIO0
)

//go:linkname nanotime runtime.nanotime
func nanotime() int64

func interrupt(p machine.Pin) {
    now := nanotime()
    if p == UpButton {
        if now-lastEdgeUp < debounceNs {
            return
        }
        lastEdgeUp = now
    } else {
        if now-lastEdgeDn < debounceNs {
            return
        }
        lastEdgeDn = now
    }

    up := !UpButton.Get()
    down := !DownButton.Get()
    if up {
        upSeen = true
    }
    if down {
        downSeen = true
    }

    if !up && !down && (upSeen || downSeen) {
        var ev uint32
        switch {
        case upSeen && downSeen:
            ev = evConfirm
        case upSeen:
            ev = evUp
        default:
            ev = evDown
        }
        upSeen, downSeen = false, false

        head := atomic.LoadUint32(&ringHead)
        if head - atomic.LoadUint32(&ringTail) < ringSize { //drop if full
            ring[head%ringSize] = ev
            atomic.StoreUint32(&ringHead, head+1)
        }
    }
}

func (TDisplayS3Input) Channel(LatestState func() shared.InputState) chan shared.InputState {
    channel := make(chan shared.InputState)

    for _, b := range []machine.Pin{UpButton, DownButton} {
        b.Configure(machine.PinConfig{
			Mode: machine.PinInputPullup,
		})
        b.SetInterrupt(machine.PinRising|machine.PinFalling, interrupt)
    }

    go func() {
        for {
            tail := atomic.LoadUint32(&ringTail)
            if tail == atomic.LoadUint32(&ringHead) {
                time.Sleep(5 * time.Millisecond)
                continue
            }
            ev := ring[tail%ringSize]
            atomic.StoreUint32(&ringTail, tail+1)

            state := LatestState()
            switch ev {
            case evConfirm:
                state = libinput.Mutate_ThreeButtonInput(state, libinput.EnumString("confirm"))
            case evUp:
                state = libinput.Mutate_ThreeButtonInput(state, libinput.EnumString("up"))
            case evDown:
                state = libinput.Mutate_ThreeButtonInput(state, libinput.EnumString("down"))
            }
            channel <- state 
        }
    }()

    return channel
}