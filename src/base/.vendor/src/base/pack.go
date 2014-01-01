// Pack
package base

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Pack struct {
	message []byte
	length  int
	index   int
}

func NewPack(in []byte) *Pack { return &Pack{in, len(in), 0} }

func NewPackEmpty() *Pack {
	in := make([]byte, 0)
	return &Pack{in, 0, 0}
}

/*-------------------------------------------读取---------------------------------------------*/
/*读取 int  8位（1字节）*/
func (b *Pack) ReadInt8() int {
	in := b.message[b.index : b.index+1]
	b.index = b.index + 1
	return int(in[0])
}

/*读取 int  16位（2字节）*/
func (b *Pack) ReadInt16() int {
	in := b.message[b.index : b.index+2]
	result := int(in[1]) | (int(in[0]) << 8)
	b.index = b.index + 2
	return result
}

/*读取 int  32位（4字节）*/
func (b *Pack) ReadInt32() int {
	in := b.message[b.index : b.index+4]
	result := int(in[3]) | (int(in[2]) << 8) | (int(in[1]) << 16) | (int(in[0]) << 24)
	b.index = b.index + 4
	return result
}

/*读取 float  32位（4字节）*/
func (b *Pack) ReadFloat32() float32 {
	in := b.message[b.index : b.index+4]
	var result float32
	buf := bytes.NewBuffer(in)
	err := binary.Read(buf, binary.BigEndian, &result)
	if err != nil {
		fmt.Println("float解析失败", err)
	}
	b.index = b.index + 4
	return result
}

func TraceBytes(b []byte) {
	fmt.Print("[ ")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%08b ", b[i])
	}
	fmt.Print("]")
}

/*读取 double  64位（8字节）*/
func (b *Pack) ReadDouble64() float64 {
	in := b.message[b.index : b.index+8]
	var result float64
	buf := bytes.NewBuffer(in)
	err := binary.Read(buf, binary.BigEndian, &result)
	if err != nil {
		fmt.Println("double解析失败", err)
	}
	b.index = b.index + 8
	return result
}
func (b *Pack) ReadDouble642() uint64 {
	in := b.message[b.index : b.index+8]
	TraceBytes(in)
	var result uint64
	buf := bytes.NewBuffer(in)
	err := binary.Read(buf, binary.BigEndian, &result)
	if err != nil {
		fmt.Println("double解析失败", err)
	}
	b.index = b.index + 8
	return result
}

/*读取 string（string的前面嵌入32位的长度）*/
func (b *Pack) ReadString() string {
	length := b.ReadInt16()
	in := b.message[b.index : b.index+length]
	b.index = b.index + length
	return string(in)
}

/*-------------------------------------------写入---------------------------------------------*/
/*写入int  8位（1字节）*/
func (b *Pack) WriteInt8(value int) {
	by := byte(value)
	b.message = append(b.message, by)
	b.length = len(b.message)
}

/*写入int  16位（2字节）*/
func (b *Pack) WriteInt16(value int) {
	by := make([]byte, 2)
	by[1] = byte(value >> 8)
	by[0] = byte(value)
	b.message = append(b.message, by...)
	b.length = len(b.message)
}

/*写入int  32位（4字节）*/
func (b *Pack) WriteInt32(value int) {
	by := make([]byte, 4)
	by[3] = byte(value >> 24)
	by[2] = byte(value >> 16)
	by[1] = byte(value >> 8)
	by[0] = byte(value)
	b.message = append(b.message, by...)
	b.length = len(b.message)
}

/*写入 float 32位（4字节）*/
func (b *Pack) WriteFloat32(value float32) {
	by := make([]byte, 0)
	buf := bytes.NewBuffer(by)
	err := binary.Write(buf, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("float写入失败", err)
	}
	b.message = append(b.message, buf.Bytes()...)
	b.length = len(b.message)
}

/*写入 double 64位（8字节）*/
func (b *Pack) WriteDouble64(value float64) {
	by := make([]byte, 0)
	buf := bytes.NewBuffer(by)
	err := binary.Write(buf, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("double写入失败", err)
	}
	b.message = append(b.message, buf.Bytes()...)
	b.length = len(b.message)
}

/*写入 String（string的前面嵌入32位的长度）*/
func (b *Pack) WriteString(value string) {
	by := []byte(value)
	b.WriteInt16(len(by))
	b.message = append(b.message, by...)
	b.length = len(b.message)
}

func (b *Pack) WriteBytes(value []byte) {
	b.message = append(b.message, value...)
	b.length = len(b.message)
}

/*-------------------------------------------其它---------------------------------------------*/
/* 获取对象的 byteArray 值 */
func (b *Pack) Bytes() []byte {
	return b.message
}

func (b *Pack) Clear() {
	b.message = make([]byte, 0)
	b.length = 0
	b.index = 0
}
func (b *Pack) Index() int {
	return b.index
}
func (b *Pack) Reset() { b.index = 0 }

/* 从当前位置读取剩下的全部 */
func (b *Pack) ReadBytes() []byte {
	result := b.message[b.index:b.length]
	b.index = b.length
	return result
}
