package handle

import (
	. "base"
	_ "fmt"
	"reflect"
)

type ACMD struct {
	Code uint16                     //协议号
	Func func(uint16, *Pack) []byte //协议号对应函数
}

var DIC map[uint16]ACMD = map[uint16]ACMD{} //以字典形式存在的协议
var CMD CmdStuct                            //以结构形式存在的协议

type CmdStuct struct {
	C1 ACMD
	C2 ACMD
}

func init() {
	//配置协议
	CMD.C1 = ACMD{1, f1}
	CMD.C2 = ACMD{2, f2}

	//利用reflect解析结构
	v := reflect.ValueOf(CMD)
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		switch t := value.Interface().(type) {
		case ACMD:
			DIC[t.Code] = t //将协议写到 map 中
		}
	}
}
func f1(c uint16, p *Pack) []byte {
	return []byte("协议处理函数f1")
}
func f2(c uint16, p *Pack) []byte {
	return []byte("协议处理函数f2")
}
