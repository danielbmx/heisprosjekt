// State Machine 

package "./elevdriver"
import "fmt"
import "time"


// States
type State int
const (
    INVALID State = iota
    MOVING
    STANDSTILL
    //EMG_STOPPED
    N_STATES
)

// Events
type Event int
const(
	INITIALIZE = iota
	NOEVENT
	MOVE
	HALT
	//EMGSTOP
	N_EVENTS
)

// Create Channels for use 
buttonEventChan         := make(chan elevdriver.Button)
floorEventChan          := make(chan int)
dir
// Private varialbles:
var last_floor = make(chan int)

var direction = make(chan elevdriver.Direction)
var timer_start int
var current_state State
var event Event


func UpdateState() State{
    event := G

    switch State {
        
        case 
        
        case 
        
        
        case 
      
    }

}

func GetNextEvent(state State, floor int) Event{
    switch(state){
    
        case INITIALIZE:
        
        case MOVING:
        
        case STANDSTILL:
         
    }

}

// Initialize system and drive car down to closest floor
func Init(  buttonEventChan         chan elevdriver.Button,
            floorEventChan          chan int,
            stopButtonEventChan     chan bool,
            obstructionEventChan    chan bool
            dirEventChan            chan elevdriver.Direction){

    val := IoInit()
    if !val {
        fmt.Printf("Driver initiated\n")
    } else {
        fmt.Printf("Driver not initiated\n")
    }
    
    elevdriver.SetMotorDir(elevdriver.NONE)

    elevdriver.Poller(  buttonEventChan,
                floorEventChan,
                stopButtonEventChan,
                obstructionEventChan)
                

    // Drive down to nearest floor and stop
    if <-floorEventChan != -1 {
    	return
    }
    
    for <-floorEventChan == -1{
    	time.Sleep(25*time.Millisecond)
    	SetMotorDir(elevdriver.DOWN)  
    	fmt.Println("Hit2")  	
    }
    
    elevdriver.ElevatorStop(elevdriver.DOWN)
 
    // Initialize FSM variables:
    last_floor:=<-floorEventChan
    dirEventChan:=<- elevdriver.NONE
    current_state := STANDSTILL
    event := NOEVENT
    
}




// Variabel to save new state
var NewState State

/*
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

*/





















