package main

import (
	"fmt"
	"runtime"
	"strconv"
)

type XandYaxis struct {
	X int
	Y int
}

func (xY XandYaxis) functionAlsMethode() {
	printFunctionName(getCurrentFrame())
	fmt.Println("Input", xY)
}

func (xY XandYaxis) functionAlsMethodeMitParametern(i int, i2 int) {
	printFunctionName(getCurrentFrame())
	fmt.Println("Input", xY, i, i2)
}

func (xY XandYaxis) functionAlsMethodeAufObjektMitParameternUndReturnWert(i int, i2 int) int {
	printFunctionName(getCurrentFrame())
	fmt.Println("Input", xY, i, i2)
	xY.X = 11
	fmt.Println(xY.X, "*", i, "+", xY.Y, "*", i2, "=")
	return xY.X*i + xY.Y*i2
}

func (xY *XandYaxis) functionAlsMethodeAufObjektPointerMitParameternUndReturnWert(i int, i2 int) int {
	printFunctionName(getCurrentFrame())
	fmt.Println("Input", xY, i, i2)
	xY.X = 11
	fmt.Println(xY.X, "*", i, "+", xY.Y, "*", i2, "=")
	return xY.X*i + xY.Y*i2
}

func main() {
	normalFunction()

	functionMitParameter(1, 2, 3)

	i := functionMitParameterUndReturnWert(1, 2, 3)
	fmt.Println("Resultat", i)

	i, j, k := functionMitParameterUndMehrerenReturnWerten(1, 2, 3)
	fmt.Println("Resultat", i, j, k)

	xy := XandYaxis{1, 2}
	functionMitStructAlsParameter(xy)
	fmt.Println("Resultat", xy)

	functionMitPointerAufStructAlsParameter(&xy)
	fmt.Println("Resultat", xy)

	xy = functionMitParameterUndStructAlsReturnWert(2, 3)
	fmt.Println("Resultat", xy)

	xy1, xy2, xy3 := functionMitParameterUndMehrerenStructsAlsReturnWert(2, 3)
	fmt.Println("Resultat", xy1, xy2, xy3)

	if1 := functionMitParameterUndInterfaceAlsReturnWert(2, 3)
	fmt.Println("Resultat", if1)

	if1, if2, if3 := functionMitParameterUndMehrerenInterfacesAlsReturnWert(2, 3)
	fmt.Println("Resultat", if1, if2, if3)

	if1, if2, if3 = functionMitMehrerenInterfacesAlsParameter(if1, if2, if3)
	fmt.Println("Resultat", if1, if2, if3)

	a := functionMitParameterUndPointerAlsReturnWert(2)
	fmt.Println("Resultat", *a)

	arr := functionMitParameterUndArrayAlsReturnWert(2, 3, 54)
	fmt.Println("Resultat", arr)

	xy = XandYaxis{2, 3}
	xy.functionAlsMethode()
	xy.functionAlsMethodeMitParametern(1, 2)
	fmt.Println("Input", xy)
	i = xy.functionAlsMethodeAufObjektMitParameternUndReturnWert(1, 2)
	fmt.Println("Resultat", i, xy)

	xyPointer := &XandYaxis{3, 4}
	fmt.Println("Input", xyPointer)
	i = xyPointer.functionAlsMethodeAufObjektMitParameternUndReturnWert(2, 7)
	fmt.Println("Resultat", i, xyPointer)

	xy = XandYaxis{2, 3}
	fmt.Println("Input", xy)
	i = xy.functionAlsMethodeAufObjektPointerMitParameternUndReturnWert(1, 2)
	fmt.Println("Resultat", i, xy)

	xyPointer = &XandYaxis{3, 4}
	fmt.Println("Input", xyPointer)
	i = xyPointer.functionAlsMethodeAufObjektPointerMitParameternUndReturnWert(2, 7)
	fmt.Println("Resultat", i, xyPointer)

	multipleCallsToPointerReturn()
}

func multipleCallsToPointerReturn() {
	printFunctionName(getCurrentFrame())
	a := getXyAxisWithPointerReturnInOne(1, 2)
	fmt.Println(a)
	b := getXyAxisWithPointerReturnInOne(11, 22)
	fmt.Println(a, b)

	a = getXyAxisWithPointerReturn(1, 2)
	fmt.Println(a)
	b = getXyAxisWithPointerReturn(11, 22)
	fmt.Println(a, b)
}

func getXyAxisWithPointerReturnInOne(i int, i2 int) *XandYaxis {
	return &XandYaxis{i, i2}
}

func getXyAxisWithPointerReturn(i int, i2 int) *XandYaxis {
	retVal := new(XandYaxis)
	retVal.X = i
	retVal.Y = i2
	return retVal
}

func functionMitParameterUndArrayAlsReturnWert(i int, i2 int, i3 int) []int {
	printFunctionName(getCurrentFrame())
	return []int{i, i2, i3}
}

func functionMitParameterUndPointerAlsReturnWert(i1 int) *int {
	//iiihh... dont do that, but it works :-)
	printFunctionName(getCurrentFrame())
	lv := i1 * i1
	return &lv
}

func functionMitMehrerenInterfacesAlsParameter(if1 interface{}, if2 interface{}, if3 interface{}) (interface{}, interface{}, interface{}) {
	printFunctionName(getCurrentFrame())
	//Cast to correct type
	if1Local, ok := if1.(XandYaxis)
	if ok {
		if1Local.X = 42
	}
	return if1Local, if2, if3
}

func functionMitParameterUndMehrerenInterfacesAlsReturnWert(i int, i2 int) (interface{}, interface{}, interface{}) {
	printFunctionName(getCurrentFrame())
	return XandYaxis{i, i2}, XandYaxis{i * i, i * i2}, XandYaxis{i2 * i2, 0}
}

func functionMitParameterUndInterfaceAlsReturnWert(i int, i2 int) interface{} {
	printFunctionName(getCurrentFrame())
	return XandYaxis{i, i2}
}

func functionMitParameterUndMehrerenStructsAlsReturnWert(i int, i2 int) (XandYaxis, XandYaxis, XandYaxis) {
	printFunctionName(getCurrentFrame())
	return XandYaxis{i, i2}, XandYaxis{i * i, i * i2}, XandYaxis{i2 * i2, 0}
}

func functionMitParameterUndStructAlsReturnWert(i int, i2 int) XandYaxis {
	printFunctionName(getCurrentFrame())
	return XandYaxis{i, i2}
}

func functionMitPointerAufStructAlsParameter(xy *XandYaxis) {
	printFunctionName(getCurrentFrame())
	fmt.Println("Input", xy)
	xy.X += 1
	fmt.Println("Modified", xy)
}

func functionMitStructAlsParameter(xy XandYaxis) {
	printFunctionName(getCurrentFrame())
	fmt.Println("Input", xy)
	xy.X += 1
	fmt.Println("Modified", xy)
}

func functionMitParameterUndMehrerenReturnWerten(i int, i2 int, i3 int) (int, int, int) {
	printFunctionName(getCurrentFrame())
	return i, i2, i3
}

func functionMitParameterUndReturnWert(i int, i2 int, i3 int) int {
	printFunctionName(getCurrentFrame())
	return i + i2 + i3
}

func functionMitParameter(i int, i2 int, i3 int) {
	printFunctionName(getCurrentFrame())
	fmt.Println("Resultat", i+i2+i3)
}

func printFunctionName(frame runtime.Frame) {
	fmt.Println("Wir sind in", frame.Function+"() ["+frame.File+":"+strconv.Itoa(frame.Line)+"]")
}

func normalFunction() {
	dasIstMeineNormaleFunktionOhneIrgendwas()
}

func dasIstMeineNormaleFunktionOhneIrgendwas() {
	printFunctionName(getCurrentFrame())
}

func getCurrentFrame() runtime.Frame {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame
}
