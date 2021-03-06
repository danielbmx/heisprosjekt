// Network module
package networkmodule

import (
	"fmt"
	"net"
	"time"
	"encoding/json"
	elevdriver "../elevdriver"
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


// Recieve message via UDP
func UdpButtonReciver(message_channel chan elevdriver.Button) {
    
    serverAddr_udp, err := net.ResolveUDPAddr("udp", ":20005")
	PrintError(err)

    con_udp, err := net.ListenUDP("udp", serverAddr_udp)
    PrintError(err)
    save := elevdriver.Button{} 
    buffer := make([]byte,1024)
	//connection, err := net.ListenUDP("udp", UDP_addr)
	//PrintError(err)
	
	for {
        n, _,err := con_udp.ReadFromUDP(buffer)
        PrintError(err)
        
        err1 := json.Unmarshal(buffer[0:n],&save)
        PrintError(err1)
        message_channel <- save
    }
    
}


// Create UDP connection
func UdpConnect(address string) *net.UDPConn{
	serverAddr_udp, err := net.ResolveUDPAddr("udp", address)
	PrintError(err)

    con_udp, err := net.DialUDP("udp", nil, serverAddr_udp)
    PrintError(err)
    
    return con_udp
}





// Broadcast message via UDP using Json
func UdpButtonSender(parameter elevdriver.Button, con_udp *net.UDPConn) {
    fmt.Println("in udpSender")
    message, err := json.Marshal(parameter) 
    PrintError(err)
	
	for {
	fmt.Println("for in udpSender")
		time.Sleep(1000 * time.Millisecond)
		_, err2 := con_udp.Write(message)
		PrintError(err2)
	}
}

func UdpAliveReciver(message_alive chan int) {
    
    serverAddr_udp, err := net.ResolveUDPAddr("udp", ":20006")
	PrintError(err)

    con_udp, err := net.ListenUDP("udp", serverAddr_udp)
    PrintError(err)
    save := 0
    buffer := make([]byte,1024)
	//connection, err := net.ListenUDP("udp", UDP_addr)
	//PrintError(err)
	
	for {
        n, _,_ := con_udp.ReadFromUDP(buffer)
        //PrintError(err)
        
        err1 := json.Unmarshal(buffer[0:n],&save)
        PrintError(err1)
        message_alive <- save
    }
}


func UdpAliveSender(parameter int, con_udp *net.UDPConn) {
    message, err := json.Marshal(parameter) 
    PrintError(err)
	
	for {
	fmt.Println("for in udpSender")
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
/*
func OrderDecide(messageAlive chan int){
	UdpAliveReciever(messageAlive)
	cost1 := <-messageAlive
	cost2 := <-messageAlive
	if cost1 >= cost2 
	
}
*/
























