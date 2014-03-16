// State Machine 

package statemachine

import ("time"
		"fmt"
		networkmodule "../networkmodule"
		elevdriver "../elevdriver"
		)



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
var last_floor = make(chan int, 1)
var timeStart time.Time
var currentState State = INVALID
var direction elevdriver.Direction
var event Event


// channels for listening:
var ButtonEventChan      = make(chan elevdriver.Button, 1)
var FloorEventChan       = make(chan int, 1)
var DirEventChan         = make(chan elevdriver.Direction, 1)



func UpdateState() { 
	for {
/**/
		time.Sleep(10*time.Millisecond)
/**/	
		event := GetNextEvent(currentState, DirEventChan, FloorEventChan, networkmodule.OrderChannel)
		//fmt.Println("State:", currentState, "  Event:", event)
		switch currentState {
		    
		    case INVALID:
		         switch event {
		            case INITIALIZE:
		                elevdriver.Init(ButtonEventChan, FloorEventChan, DirEventChan)
		                currentState = STANDSTILL
		                fmt.Println("US: INVALID-INITIALIZE")
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
		                
		                //Stop car
		                //Setmot
		                //Open door for 3 sec
		                //Delete this order from queue HAPPENS IN ORDERSYSTEM!
		                //current state = STANDSTILL
		                fmt.Println("US: In HALT")
		                elevdriver.ElevatorStop(elevdriver.UP)
		                elevdriver.SetDoorOpenLight(elevdriver.ON)
		                timeStart = time.Now()
		                currentState = STANDSTILL
		                }
		                
		    case STANDSTILL:
		        switch event {
		            
		            case INITIALIZE:
		                break
		            
		            case NOEVENT:
		                //Close door if door has been open for more than 3 sec
		                closeTime := time.Now()
		                if closeTime.Sub(timeStart) > 3 {
		                    elevdriver.SetDoorOpenLight(elevdriver.OFF)
		                    }
		                    
		            case MOVE:
		                //move car in right direction
		                networkmodule.GetNextDirection(DirEventChan, FloorEventChan, networkmodule.OrderChannel)
		                nextDir := <- DirEventChan
		                DirEventChan <- nextDir
		                elevdriver.SetMotorDir(nextDir)
		                currentState = MOVING
		            
		            case HALT:
		                //Open door
		                timeStart = time.Now()
		                elevdriver.SetDoorOpenLight(elevdriver.ON)
		       }		// Delete order
		}
		<-FloorEventChan
	}
}


func GetNextEvent(currentState State, dirEventChan chan elevdriver.Direction, floorEventChan chan int, orderChan chan[4][3] int) Event { 
    //fmt.Println("Inside GetNextEvent")
    switch currentState {
    
        case INVALID:
        	return INITIALIZE
        
        case MOVING:
        	//fmt.Println("case MOVING")
        	// if should stop
        	//return HALT

        	// "Compare"
        	if networkmodule.StopAtFloor(dirEventChan, floorEventChan, orderChan){
        		return HALT
        	}

        	// if shoud continue moving
        	// return NOEVENT 
        	return NOEVENT   
        
        case STANDSTILL:

        	fmt.Println("GNE: Inside Standstill")
        	//if order in queue
        	nextDir := <- DirEventChan
		    DirEventChan <- nextDir
		    fmt.Println("GNE: nextDir= ", nextDir)
		    
        	if nextDir != elevdriver.NONE {//&& time.Now().Sub(timeStart) > 3 
				fmt.Println("GNE: getting dir")
        		return MOVE
        	
        	}
        	
        	// if no order in queue
        	// return NOEVENT


	}
	fmt.Println("GNE: return NOEVENT")
    return NOEVENT

}























