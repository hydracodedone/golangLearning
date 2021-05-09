package main

import "fmt"

type readable interface {
	read()
}
type writeable interface {
	write()
}
type readAndWritable interface {
	readable
	writeable
}
type reader struct {
}
type writer struct {
}
type file struct {
}

func (readableInstance *reader) read() {

}
func (writeableInstance *writer) write() {

}
func (fileInstance *file) read() {

}
func (fileInstance *file) write() {

}

func main() {
	readerInstance := new(reader)
	writerInstance := new(writer)
	fileInstance := new(file)
	var readerAble readable
	var writerAble writeable
	var readerAndWriterable readAndWritable
	readerAble = readerInstance
	writerAble = writerInstance
	readerAndWriterable = fileInstance
	readerAble = fileInstance
	writerAble = fileInstance
	readerAble.read()
	writerAble.write()
	readerAndWriterable.read()
	readerAndWriterable.write()
	newReaderInstance := new(readable)
	fmt.Printf("The newReaderInstance type is %T\n", newReaderInstance)
	*newReaderInstance = fileInstance
	fmt.Printf("The newReaderInstance type is %T\n", newReaderInstance)
	(*newReaderInstance).read()
}
