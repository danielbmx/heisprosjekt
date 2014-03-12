package main

import ("./networkmodule"
        "./elevdriver"
        "fmt")

func main(){
   
   fmt.Println("main")
   
   // Declare channels
    ButtonEventChan      := make(chan elevdriver.Button)
    FloorEventChan       := make(chan int)
   
   
    go elevdriver.Poller(buttonEventChan, floorEventChan)
    
    /*
    go HandleOrder()
   
    for{
         select{
            case ...
            case ...
         
         
         
        }
    }
   */


}
