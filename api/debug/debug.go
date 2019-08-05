package debug

import "fmt"

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
