package main

import (//"./networkmodule"
        "./elevdriver"
        "fmt")

func main(){
   
   fmt.Println("main")
   
   /*
   // Declare channels
    ButtonEventChan      := make(chan elevdriver.Button)
    FloorEventChan       := make(chan int)
   
   
    go elevdriver.Poller(buttonEventChan, floorEventChan)
    
    
    go HandleOrder()
   
    for{
         select{
            case ...
            case ...
         
         
         
        }
    }
   */
	//message_chan := make(chan elevdriver.Button) 
	
	knapp := elevdriver.Button{}
	fmt.Println(knapp.Floor)
	
	/*
	go UdpButtonReciver(message_chan)
	fmt.Println("før knapp")
	knapp = <- message_chan
	fmt.Println("recieved")
	fmt.Println(knapp)
	
	elevdriver.SetButtonLight(knapp.Floor, knapp.Dir, elevdriver.OFF)
	fmt.Println("ferdi")
	*/
}
