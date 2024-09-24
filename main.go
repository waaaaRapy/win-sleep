package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
)

func main() {
	var (
		hybernate = flag.Bool("h", false, "Hybernate")
		sleep     = flag.Bool("s", false, "Sleep")
		disable   = flag.Bool("d", false, "Disable wake events")
	)
	flag.Parse()

	if *hybernate && *sleep || !*hybernate && !*sleep {
		fmt.Println("You must specify either -h or -s")
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	if *sleep {
		setSuspendState(0, 0, boolToUintptr(*disable))
	}

	if *hybernate {
		setSuspendState(1, 0, boolToUintptr(*disable))
	}
}

func boolToUintptr(b bool) uintptr {
	if b {
		return 1
	} else {
		return 0
	}
}

/*
Call `SetSuspendState` procedure in the `powrprof.dll`

See https://learn.microsoft.com/en-us/windows/win32/api/powrprof/nf-powrprof-setsuspendstate

Parameters:
- hybernate: hibernate (1) or sleep (0).
- forceCritical: this parameter has no effect.
- disableWakeEvent: whether to disable wake events (1) or not (0).
*/
func setSuspendState(bHybernate uintptr, bForce uintptr, bWakeupEventsDisabled uintptr) {
	powrprof := syscall.NewLazyDLL("powrprof.dll")

	setSuspendState := powrprof.NewProc("SetSuspendState")

	setSuspendState.Call(bHybernate, bForce, bWakeupEventsDisabled)
}
