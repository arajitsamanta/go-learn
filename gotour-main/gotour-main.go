package main

import (
	"fmt"
	"go-learn/basic"
	"go-learn/concurrency"
	"go-learn/flowcontrol"
	"go-learn/interfaces"
	"go-learn/method"
	"go-learn/types"
	"math"
)

func main() {

	fmt.Printf("============== START - Running gotour Basic codes =============\n")
	basic.TestPackages()
	basic.TestImport()
	basic.Add(2, 5)
	basic.AddOmitTypeDeclaration(6464666646, 8794984945)
	basic.Swap("6", "87")
	basic.Split(17)
	basic.VaribalesDeclaration()
	basic.VaribalesWithInitializers()
	basic.ShortVaribaleDeclaration()
	basic.DataTypes()
	basic.ZeroValueVariable()
	basic.ConvertTypes()
	basic.TypeInterference()
	basic.Constants()
	basic.NumericConstant()
	fmt.Printf("\n=============== END - Basic =============\n")

	fmt.Printf("\n============== START - Running gotour FlowControl codes =============\n")
	flowcontrol.SimpleForLoop()
	flowcontrol.ForOptionalInitAndPost()
	//flowcontrol.ForeverLoop()
	fmt.Println(flowcontrol.Sqrt(-4), flowcontrol.Sqrt(2))
	fmt.Println(flowcontrol.IfWithShortStatement(3, 2, 10))
	fmt.Println(flowcontrol.IfWithElse(3, 2, 20))
	//flowcontrol.SqrtMatching(1)
	flowcontrol.SqrtMatching(25000)
	fmt.Println("actual sqrt", math.Sqrt(25000))
	flowcontrol.SwitchExample()
	flowcontrol.SwitchInOrder()
	flowcontrol.SwitchNoCondition()
	flowcontrol.DeferExample()
	fmt.Printf("\n\n=============== END - FlowControl =============\n")

	fmt.Printf("\n============== START - Running gotour Types(structs, slice, maps, pointers etc) codes =============\n")
	//Pointers
	types.HelloPointers()

	//Structures
	types.HelloStructures()
	types.StructureFieldAccess()
	types.StructureFieldAccessUsingPointers()
	types.StructLiterals()

	//Arrays and Slices
	types.HelloArrays()
	types.HelloSlices()
	types.SlicePointers()
	types.SliceLiterals()
	types.SliceDefaults()
	types.SliceLengthAndCapacity()
	types.NilSlices()
	types.SliceWithMake()
	types.SlicesOfSlice()
	types.SliceAppend()
	types.RangeOverSlice()
	types.SliceExcercise()

	//Maps
	types.HelloMaps()
	types.MapLiterals()
	types.MapLiteralsWithNoTypeName()
	types.MapMutation()
	types.MapExcercise()

	//Function as values
	types.FunctionAsValues()
	types.FunctionClosures()
	types.FibonacciWithFunctionClosure()

	fmt.Printf("\n\n=============== END - Types =============\n")

	fmt.Printf("\n============== START - Running gotour methods and interfaces codes =============\n")
	//Go methods
	method.HelloGoMethods()
	method.GoMethodsAreFunctions()
	method.MethodsOnNonStructTypes()
	method.MethodsWithPointerReceiver()
	method.MethodsWithPointerReceiver2()
	method.MethodsAndPointerIndirection()
	method.MethodsAndPointerIndirection2()
	method.MethodsWithValueOrPointerReceiver()

	//Go interfacs
	interfaces.HelloInterfaces()
	interfaces.ImplicitInterfaceImplementations()
	interfaces.InterfaceValues()
	interfaces.InterfaceWithNilValues()
	//interfaces.NilInterfaceValues() - No interface implementation. Go does not know which mehtod to call
	interfaces.EmptyInterfaces()
	interfaces.TypeAssertions()
	interfaces.TypeSwitch()
	interfaces.HelloStringers()
	interfaces.StringersExcercise()
	interfaces.GoErrors()
	interfaces.ErrorsExcercise()
	interfaces.GoReaders()
	interfaces.ReaderExcercise()
	interfaces.ROT13Reader()
	interfaces.GoImages()
	interfaces.ImageExcercise()
	fmt.Printf("\n\n=============== END - Methods & Interfaces =============\n")

	fmt.Printf("\n============== START - Running Go concurrency codes =============\n")
	concurrency.GoRoutines()
	concurrency.GoChannels()
	concurrency.BufferedChannel()
	concurrency.RangeAndClose()
	concurrency.GoSelect()
	concurrency.DefaultSelect()
	concurrency.EquivalentBinaryTree()
	concurrency.MutexExample()
	concurrency.WebCrawler()
	fmt.Printf("\n\n==========)===== END -  Go concurrency=============\n")

}
