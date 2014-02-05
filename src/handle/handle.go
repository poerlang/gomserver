package handle

import (
	. "base"
	_ "fmt"
	"reflect"
)

type ACMD struct {
	Code uint16                              //协议号
	Func func(uint16, *Pack, *Player) []byte //协议号对应函数
}

var DIC map[uint16]ACMD = map[uint16]ACMD{} //以字典形式存在的协议
var CMD CmdStuct                            //以结构形式存在的协议

type CmdStuct struct {
	//moeditor struct start
	C10000up ACMD
	C10001up ACMD
	C12000up ACMD
	C12001up ACMD
	//moeditor struct end
}

func init() {
	//moeditor init start
	CMD.C10000up = ACMD{10000, f10000Up}
	CMD.C10001up = ACMD{10001, f10001Up}
	CMD.C12000up = ACMD{12000, f12000Up}
	CMD.C12001up = ACMD{12001, f12001Up}
	//moeditor init end

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
