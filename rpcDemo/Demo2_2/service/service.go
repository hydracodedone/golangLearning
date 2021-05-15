package service

import (
	"fmt"
	"io"
	"log"
)

type QueryService struct {
}

//QueryWithServerStream Server Stream Means Receive The Whole Client Information And Response In Stream
func (q *QueryService) QueryWithServerStream(requestIdList *RequestIdList, streamQueryServer QueryService_QueryWithServerStreamServer) error {
	responseIdList := new(ResponseIdList)
	for key := range requestIdList.GetRequestIdList() {
		key++
		responseIdList.ResponseIdList = append(responseIdList.ResponseIdList, &ResponseId{ResponseId: int32(key)})
		if key%2 == 0 && key > 0 {
			fmt.Printf("The Send Message is %v\n", responseIdList)
			// the key fragments,we send the fragments of the response to the client
			err := streamQueryServer.Send(responseIdList)
			if err != nil {
				log.Fatalf("The Stream is error:<%s>", err.Error())
			}
			responseIdList = new(ResponseIdList)
		}
	}
	// final send
	if len(responseIdList.ResponseIdList) > 0 {
		fmt.Printf("The Send Message is %v\n", responseIdList)
		err := streamQueryServer.Send(responseIdList)
		if err != nil {
			log.Fatalf("The Stream is error:<%s>", err.Error())
		}
	}
	return nil
}

// QueryWithClientStream Means Receive The Client Information In Stream Way And Response With ClientStream Server
func (q *QueryService) QueryWithClientStream(clientStreamServer QueryService_QueryWithClientStreamServer) error {
	responseInstance := ResponseIdList{
		ResponseIdList: []*ResponseId{},
	}
	for {
		requestInstance, err := clientStreamServer.Recv()
		if err != nil {
			if err == io.EOF {
				err := clientStreamServer.SendAndClose(&responseInstance)
				responseInstance = ResponseIdList{
					ResponseIdList: []*ResponseId{},
				}
				if err != nil {
					log.Fatalf("Client Stream Send And Close Fail:<%s>", err.Error())
				}
				break
			} else {
				log.Fatalf("The Response Stream Is Error:<%s>\n", err.Error())
			}
		} else {
			// We Can Use Goroutine To Handle The Receive Information
			for _, value := range requestInstance.RequestIdList {
				responseInstance.ResponseIdList = append(responseInstance.ResponseIdList, &ResponseId{
					ResponseId: value.RequestId,
				})
			}
		}
	}
	return nil
}
