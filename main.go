package main

import (
	"fmt"
	"glean/typeGlean"
)

type IGleanWork interface {
	Run() error
}

type Action func(gleanWork IGleanWork) error

func main() {
	var i IGleanWork = &typeGlean.TypeClass{}

	run(work, i)
}

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
