// Network module
package main
import (
	"fmt"
	"net"
	"time"
	"encoding/json"
	"./elevdriver"
) 
/*
// Confirm elevator order taken to other elevators in the network
func UdpConfirmOrder() {
	UdpSender()
}

// Broadcast order recieved to all elevators in the network
func PassOrder() {

}
*/

// main function for testing JSON package sending:
func main() {
	
	
	testbutton := elevdriver.Button{
		Floor : 2,
		Dir : elevdriver.DOWN,
	} 
	connection := UdpConnect()
	
	UdpSender(testbutton, connection)
	
	knapp := UdpReciver(connection)
	fmt.Println(knapp)
	
}

// Recieve message via UDP
func UdpReciver(con_udp *net.UDPConn) elevdriver.Button{
    save := elevdriver.Button{} 
    buffer := make([]byte,1024)
	//connection, err := net.ListenUDP("udp", UDP_addr)
	//PrintError(err)
	
	for {
        _, _,_ = con_udp.ReadFromUDP(buffer)
        //PrintError(err)
        
        err1 := json.Unmarshal(buffer,&save)
        PrintError(err1)
    }
    return save
    
}

// Create UDP connection
func UdpConnect() *net.UDPConn{
	serverAddr_udp, err := net.ResolveUDPAddr("udp", "129.241.187.255:20020")
	PrintError(err)

    con_udp, err := net.DialUDP("udp", nil, serverAddr_udp)
    PrintError(err)
    
    return con_udp
}




// Broadcast message via UDP using Json
func UdpSender(parameter elevdriver.Button, con_udp *net.UDPConn) {
    message, err := json.Marshal(parameter) 
    PrintError(err)
	
	for {
		time.Sleep(1000 * time.Millisecond)
		_, err2 := con_udp.Write(message)
		PrintError(err2)
	}
}


func PrintError(err error) {
	if err != nil{
        fmt.Println(err)
	}
}



























