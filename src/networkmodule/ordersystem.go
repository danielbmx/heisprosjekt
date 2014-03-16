// Ordersystem
package networkmodule

import (
	"fmt"
	//"net"
	"time"
	elevdriver "../elevdriver"
) 

var OrderChannel = make(chan [4][3]int, 1)

func InitOrderMatrix(orderchan chan [4][3]int) {
	var ordermatrix[4][3]int
	
	for i := 0; i<4; i++ {
		for j := 0; j<3; j++{
			ordermatrix[i][j] = 0;
		}
	}
	orderchan <- ordermatrix
}

func SaveOrder(buttonEventChan chan elevdriver.Button, orderchan chan [4][3]int){
	for{
		button := <-buttonEventChan
		buttonEventChan <- button
		ordermatrix := <- orderchan
		ordermatrix[button.Floor - 1][button.Dir] = 1
		orderchan <- ordermatrix
		time.Sleep(50*time.Millisecond)
	}
}


func CalculateCost(buttonEventChan chan elevdriver.Button, floorEventChan chan int, floorDirectionChan chan elevdriver.Direction) int{
	dir := <- floorDirectionChan
	floorDirectionChan <- dir
	floor := <- floorEventChan
	floorEventChan <- floor
	button := <- buttonEventChan
	buttonEventChan <- button
	score := 0
	if dir != button.Dir {
		score += 1
	}
	if floor != button.Floor {
		score += 1
	}
	if button.Dir == elevdriver.DOWN {
		if floor >= button.Floor{
			score += (floor - button.Floor)
		}else{
			score += 4
		}
	}
	if button.Dir == elevdriver.UP {
		if floor <= button.Floor{
			score += (button.Floor - floor)
		}else{
			score += 4
		}
	}
	return score
}

func ResetOrder(elevator int, orderchan chan[4][3]int) {
	ordermatrix := <- orderchan
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if ordermatrix[i][j] == elevator{
				ordermatrix[i][j] = 1
			}
		}
	}
}

func DeleteOrder(button elevdriver.Button, orderchan chan[4][3]int){
   ordermatrix := <- orderchan
   ordermatrix[button.Floor][button.Dir] = 0
   orderchan <- ordermatrix
}

func HandleOrder(buttonEventChan         chan elevdriver.Button,
				  orderchan 			 chan [4][3]int){
   /*
   Provide neccesary order-handling based on information from elevdriver via channels.
   communication? 
   */
		for{
			fmt.Println("In handleOrder")
			time.Sleep(25*time.Millisecond)
			
			button := <- buttonEventChan
			buttonEventChan <- button
			elevdriver.SetButtonLight(button.Floor, button.Dir, elevdriver.ON)
			ordermatrix := <- orderchan
			ordermatrix[button.Floor - 1][button.Dir] = 1
			fmt.Println(ordermatrix)
			orderchan <- ordermatrix
			toemptybuttonChan := <- buttonEventChan
			fmt.Println("trykket:  ", toemptybuttonChan)
		}
			//UdpButtonSender(button, con_udp)
			//UdpButtonReciver(buttonEventChan)
	
}












