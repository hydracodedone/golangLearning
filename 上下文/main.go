package main

import (
	"fmt"
)

/*
A Ture Simple Custom Context Demo Based On TreeStruct For Comprehension Of The Context
Ignore MultiProcessing
Ignore Cancel Signal
*/

type customContextStruct struct {
	ChildList  []*customContextStruct
	isCanceled bool
}

func customContextInit(fatherContext *customContextStruct) *customContextStruct {
	returnInfo := &customContextStruct{
		ChildList:  nil,
		isCanceled: false,
	}
	if fatherContext == nil {
		return returnInfo
	} else {
		fatherContext.ChildList = append(fatherContext.ChildList, returnInfo)
		return returnInfo
	}
}
func (customContextStructInstance *customContextStruct) customContextCancel() {
	customContextStructInstance.isCanceled = true
	for _, each := range customContextStructInstance.ChildList {
		each.isCanceled = true
		each.customContextCancel()
	}
}

func customContext() {
	backgroundContext := customContextInit(nil)
	firstSonContext := customContextInit(backgroundContext)
	secondSonContext := customContextInit(backgroundContext)
	firstGrandSonContext := customContextInit(firstSonContext)
	firstSonContext.customContextCancel()
	fmt.Printf("The BackgroundContext Cancel Status is %t\n", backgroundContext.isCanceled)
	fmt.Printf("The FirstSonContext Cancel Status is %t\n", firstSonContext.isCanceled)
	fmt.Printf("The SecondSoneContext Cancel Status is %t\n", secondSonContext.isCanceled)
	fmt.Printf("The FirstGrandSonContext Cancel Status is %t\n", firstGrandSonContext.isCanceled)
}

func main() {
	customContext()
}
