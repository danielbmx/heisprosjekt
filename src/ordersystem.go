// Ordersystem
package main
import (
	"fmt"
	//"net"
	"time"
	"./elevdriver"
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

/*
func DistributeOrder() {

}

func ResetOrder() {

}
*/

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













