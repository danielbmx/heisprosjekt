// Ordersystem
package main
// Should be package networkmodule
import (
	"fmt"
	//"net"
	"time"
	elevdriver "../elevdriver"
) 

var OrderChannel = make(chan [4][3]int, 1)

func InitOrderMatrix( orderchan chan [4][3]int) {
	var ordermatrix[4][3]int
	
	for i := 0; i<4; i++ {
		for j := 0; j<3; j++{
			ordermatrix[i][j] = 0;
		}
	}
	orderchan <- ordermatrix
}

func SaveOrder(buttonc chan elevdriver.Button, orderchan chan [4][3]int){
	for{
	button:=<-buttonc
	matrix := <- orderchan
	matrix[button.Floor][button.Dir] = 1
	orderchan <- matrix
	time.Sleep(50*time.Millisecond)
	}
}


func CalculateCost(buttonEventChan chan elevdriver.Button, floorEventChan chan int, floorDirectionChan chan elevdriver.Direction) int{
	dir := <- floorDirectionChan
	floor := <- floorEventChan
	button := <- buttonEventChan
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

func HandleOrder(){
   /*
   Provide neccesary order-handling based on information from elevdriver via channels.
   communication? 
   */
   
   


}


func main() {

go InitOrderMatrix(OrderChannel)


button := make(chan elevdriver.Button,1000)

testbutton := elevdriver.Button{
		Floor : 3,
		Dir : elevdriver.DOWN,
}
button<-testbutton
testbutton2 := elevdriver.Button{
		Floor : 2,
		Dir : elevdriver.UP,
}
button<-testbutton2
testbutton3 := elevdriver.Button{
		Floor : 1,
		Dir : elevdriver.NONE,
}
button<-testbutton3

go SaveOrder(button, OrderChannel)


time.Sleep(1*time.Second)
fmt.Println(<-OrderChannel)

}













