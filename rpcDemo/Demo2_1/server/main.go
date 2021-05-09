package main

func main() {
	temp := make(chan int)
	go GateWithMutualCredentialRpcServer()
	go GateWithMutualCredentialHttpServer()
	<-temp
}
