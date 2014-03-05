// State Machine 

package elevdriver
import "fmt"
import "time"


// States
type State int
const (
    INVALID State = iota
    MOVING
    STANDSTILL
    EMG_STOPPED
    N_STATES
)

// Events
type Event int
const(
	NO_EVENT = iota
	MOVE
	HALT
	EMGSTOP
	N_EVENTS
)

// Create Channels for use 
buttonEventChan         := make(chan elevdriver.Button)
floorEventChan          := make(chan int)

// Variabel to save new state
var NewState State


// State Machine, switch-case 
switch State {
	case INVALID:
		// Initialize elevator
		elevdriver.Init(buttonEventChan, floorEventChan, nil, nil)
		NewState = STANDSTILL

	case MOVING:
		switch Event {
			case MOVE: 
				// Do Nothing, keep moving
			case HALT: 
				// STOOOOOOPP
				SetDoorOpenLight(elevdriver.ON)
			case EMG_STOP: 
				// NOt implemented 
				
	case STANDSTILL:
		switch Event {
			case MOVE:
				// Start moving in the right direction
				NewState = MOVING
			case HALT: 
				// Open Door
				SetDoorOpenLight(elevdriver.ON)
				time.Sleep(3*time.Second)
				SetDoorOpenLight(elevdriver.OFF) 
				 
			case EMG_STOP:
				// Not implemented

	case EMG_STOPPED:
		// Not implemented
		
		
}























