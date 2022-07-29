package typeGlean

import (
	"fmt"
	"strings"
)

type TypeFuncClass struct {
}

type testFunc func(str string) error

func (t *TypeFuncClass) Run() error {
	var i []string
	i = append(i, "I")
	i = append(i, "want")
	i = append(i, "to")
	i = append(i, "talk")
	i = append(i, "about")
	i = append(i, "some")
	i = append(i, "Type")
	i = append(i, "issues.")
	use(work, i)
	return nil
}

func (t *TypeFuncClass) Stop() error {
	return nil
}

func (t *TypeFuncClass) Close() error {
	return nil
}

func use(test testFunc, str []string) {
	fmt.Println("What are you do?")

	var value string
	for _, j := range str {
		value += j + " "
	}
	value = strings.Trim(value, " ")

	test(value)
}

func work(str string) error {
	fmt.Println(str)
	return nil
}
