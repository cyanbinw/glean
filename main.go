package main

import (
	"fmt"
	"glean/asynchronousGlean"
	"glean/genericityGlean"
	"glean/messageQueue"
	"glean/typeGlean"
	"strconv"
)

type IGleanWork interface {
	Run() error
	Stop() error
	Close() error
}

type Action func(gleanWork IGleanWork) error

func main() {

	var i = selected()
	if i == nil {
		return
	}
	run(work, i)
}

func selected() IGleanWork {
	fmt.Println(" ------------------ Please  Select ------------------ ")
	fmt.Println(" Please enter the number:")
	fmt.Println(" 1.TypeFunc ")
	fmt.Println(" 2.GenericityDemo ")
	fmt.Println(" 3.AsynchronousDemo ")
	fmt.Println(" 4.RabbitMQDemo ")
	fmt.Println(" 0.Exit ")

	var str string
	fmt.Scan(&str)
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(" Error, please re-enter")
		return selected()
	}
	var i IGleanWork
	switch num {
	case 1:
		i = &typeGlean.TypeFuncClass{}
		break
	case 2:
		i = &genericityGlean.GenericityDemo{}
		break
	case 3:
		i = &asynchronousGlean.AsynchronousDemo{}
		break
	case 4:
		i = &messageQueue.RabbitMQDemo{}
	case 5:
	case 0:
		return nil

	default:
		fmt.Println(" Error, please re-enter")
		return selected()
	}

	return i
}

// run function
func run(action Action, gleanWork IGleanWork) {
	fmt.Println(" -------------------- Work Start -------------------- ")
	err := action(gleanWork)
	if err != nil {
		fmt.Println(" -------------------- Work Error -------------------- ")
		fmt.Println(err)
		fmt.Println(" -------------------- Work Error -------------------- ")
	}
	fmt.Println(" --------------------- Work End --------------------- ")
}

func work(gleanWork IGleanWork) error {
	fmt.Println(" ------------------- Work Running ------------------- ")
	return gleanWork.Run()
}
