package handle

import (
	. "base"
	"fmt"
)

type C10000Up struct {
	uname string //用户名
	pwd   string //密码
	ctime int32  //创建时间
}

func f10000Up(c uint16, p *Pack) interface{} {
	d := new(C10000Up)
	d.uname = p.ReadString()
	d.pwd = p.ReadString()
	d.ctime = p.ReadInt32()
	fmt.Println(d) //需删除，否则影响性能
	return nil     //需修改，返回 []byte 类型的数据，否则客户端无法收到返回数据
}
