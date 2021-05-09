package main

type father struct {
}
type son struct {
	father
}

func (f *father) fatherP() {

}
func (f father) fatherT() {

}
func (s *son) sonP() {

}
func (s son) sonT() {

}

type fatherTinterface interface {
	fatherT()
}
type fatherTinterfaceP interface {
	fatherP()
}
type fatherInterfacePAndT interface {
	fatherT()
	fatherP()
}

func demo1() {
	fathert := father{}
	fatherp := &fathert
	sont := son{}
	sonp := &sont
	fathert.fatherP()
	fathert.fatherT()
	fatherp.fatherP()
	fatherp.fatherT()
	sonp.fatherP()
	sonp.fatherT()
	sont.fatherP()
	sont.fatherT()
}
func demo2() {
	fathert := father{}
	fatherp := &fathert
	sont := son{}
	sonp := &sont
	var fatherTinterfaceInstance fatherTinterface
	fatherTinterfaceInstance = fathert
	fatherTinterfaceInstance.fatherT()
	fatherTinterfaceInstance = fatherp
	fatherTinterfaceInstance.fatherT()
	fatherTinterfaceInstance = sont
	fatherTinterfaceInstance.fatherT()
	fatherTinterfaceInstance = sonp
	fatherTinterfaceInstance.fatherT()

	var fatherTinterfaceInstance2 fatherTinterfaceP
	fatherTinterfaceInstance2 = fatherp
	fatherTinterfaceInstance2.fatherP()
	fatherTinterfaceInstance2 = sonp
	fatherTinterfaceInstance2.fatherP()
}
func demo3() {
	fathert := father{}
	fatherp := &fathert
	var fatherInterfacePAndTInstance fatherInterfacePAndT
	// fatherInterfacePAndTInstance = fathert
	fatherInterfacePAndTInstance = fatherp
	fatherInterfacePAndTInstance.fatherP()
	fatherInterfacePAndTInstance.fatherT()
}

func main() {

}
