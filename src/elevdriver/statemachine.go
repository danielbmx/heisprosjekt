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


// Private variables:
var last_floor = make(chan int)
var timeStart time.Time
var currentState State
var event Event

// channels for listening:
var ButtonEventChan      = make(chan Button)
var FloorEventChan       = make(chan int)
var DirEventChan         = make(chan Direction)


func UpdateState() { // State{
    event := GetNextEvent(currentState, FloorEventChan)
    switch currentState {
        case INVALID:
             switch event {
                case INITIALIZE:
                    Init(ButtonEventChan, FloorEventChan, DirEventChan)
                case NOEVENT:
                    break 
                case MOVE:
                    break 
                case HALT:
                    break 
                    }
        
        case MOVING:
            switch event {
                case INITIALIZE:
                    break
                case NOEVENT:
                    break 
                case MOVE:
                    break 
                case HALT:
                    /* 
                    Stop car
                    Setmot
                    Open door for 3 sec
                    Delete this order from queue HAPPENDS IN ORDERSYSTEM!
                    current state = STANDSTILL
                    */ 
                    SetMotorDir(NONE)
                    SetDoorOpenLight(ON)
                    timeStart = time.Now()
                    currentState = STANDSTILL
                    }
                    
        case STANDSTILL:
            switch event {
                case INITIALIZE:
                    break
                case NOEVENT:
                    /*
                    Close door if door has been open for more than 3 sec
                    */
                    closeTime := time.Now()
                    if closeTime.Sub(timeStart) > 3 {
                        SetDoorOpenLight(OFF)
                        }
                case MOVE:
                    /*
                    move car in right direction
                    current state = MOVING
                    */
                    SetMotorDir(<-DirEventChan)
                    currentState = MOVE
                
                case HALT:
                    /*
                    Open door
                    */
                    SetDoorOpenLight(ON)
           }
}
}


func GetNextEvent(state State, floorEventChan chan Button.Floor) { //Event {
    switch currentState {
    
        case INVALID:
        	// return INITIALIZE
        
        case MOVING:
        	// if should stop
        	// return HALT 
        	
        	// if shoud continue moving
        	// return NOEVENT    
        
        case STANDSTILL:
        	// if order in queue
        	// return MOVE
        	
        	// if no order in queue
        	// return NOEVENT
    }

}


// Initialize system and drive car down to closest floor
func Init(  buttonEventChan         chan Button,
            floorEventChan          chan int,
            dirEventChan            chan Direction){

    // Check if hardware can be initialized:
    val := IoInit()
    if !val {
        fmt.Printf("Driver initiated\n")
    } else {
        fmt.Printf("Driver not initiated\n")
    }
    
    
    SetMotorDir(NONE)

      
   Poller(buttonEventChan, floorEventChan)
                

    // Drive down to nearest floor and stop
    if <-floorEventChan != -1 {
    	return
    }
    
    for <-floorEventChan == -1{
    	time.Sleep(25*time.Millisecond)
    	SetMotorDir(DOWN)  
    	fmt.Println("Hit2")  	
    }
    
    ElevatorStop(DOWN)
 
    // Initialize FSM variables:
    //last_floor:=<-floorEventChan
    dirEventChan <- NONE
    currentState := STANDSTILL
    event := NOEVENT
    
}























