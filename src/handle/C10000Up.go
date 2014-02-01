package handle

import (
	. "base"
	_ "fmt"
)

type C10000Up struct {
	v1 int32
	v2 int32
}

func f10000Up(c uint16, p *Pack) interface{} {
	return p.ReadBytesRemain()
}
