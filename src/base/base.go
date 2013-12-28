/**
 * Created by Administrator on 13-12-9.
 */
package base

import (
	"fmt"
	"runtime"
)

func Defer() {
	if r := recover(); r != nil {
		fmt.Println("【Recovered】", r)
	}
	return
}

func SetCPU() {
	cpu := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("CPU num is %d \n", cpu)
}

func SayHello(str string) {
	fmt.Println(str)
}

func CheckErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
