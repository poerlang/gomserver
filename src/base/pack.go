// Pack
package base

import (
	"fmt"
	"math"
)

type Pack struct {
	data []byte
	pos  int
}

func NewPack(in []byte) *Pack { return &Pack{in, 0} }

func NewPackEmpty() *Pack {
	in := make([]byte, 0)
	return &Pack{in, 0}
}

/*-------------------------------------------读取---------------------------------------------*/
/*读取 无符号 8位整数（1字节）*/
func (p *Pack) ReadUInt8() uint8 {
	in := p.data[p.pos : p.pos+1]
	p.pos = p.pos + 1
	return uint8(in[0])
}

/*读取 无符号 16位整数（2字节）*/
func (p *Pack) ReadUInt16() uint16 {
	in := p.data[p.pos : p.pos+2]
	result := uint16(in[1]) | (uint16(in[0]) << 8)
	p.pos = p.pos + 2
	return result
}

/*读取 无符号 32位整数（4字节）*/
func (p *Pack) ReadUInt32() uint32 {
	in := p.data[p.pos : p.pos+4]
	result := uint32(in[3]) | (uint32(in[2]) << 8) | (uint32(in[1]) << 16) | (uint32(in[0]) << 24)
	p.pos = p.pos + 4
	return result
}

/*读取 有符号 8位整数（1字节）*/
func (p *Pack) ReadInt8() int8 {
	result := p.ReadUInt8()
	return int8(result)
}

/*读取 有符号 16位整数（2字节）*/
func (p *Pack) ReadInt16() int16 {
	result := p.ReadUInt16()
	return int16(result)
}

/*读取 有符号 32位整数（4字节）*/
func (p *Pack) ReadInt32() int32 {
	result := p.ReadUInt32()
	return int32(result)
}

/*读取 无符号 64位整数（8字节）*/
func (p *Pack) ReadUInt64() uint64 {
	in := p.data[p.pos : p.pos+8]
	result := uint64(in[7]) | uint64(in[6])<<8 | uint64(in[5])<<16 | uint64(in[4])<<24 |
		uint64(in[3])<<32 | uint64(in[2])<<40 | uint64(in[1])<<48 | uint64(in[0])<<56
	p.pos = p.pos + 8
	return result
}

/*读取 有符号 64位整数（8字节）*/
func (p *Pack) ReadInt64() int64 {
	result := p.ReadUInt64()
	return int64(result)
}

/*读取 双精度 浮点数 32位*/
func (p *Pack) ReadF32() float32 {
	result := uint32(p.ReadUInt32())
	return math.Float32frombits(result)
}

/*读取 双精度 浮点数 64位*/
func (p *Pack) ReadF64() float64 {
	result := uint64(p.ReadUInt64())
	return math.Float64frombits(result)
}

/*读取 string（string的前面嵌入32位的长度）*/
func (p *Pack) ReadString() string {
	strlen := p.ReadUInt16()
	in := p.data[p.pos : p.pos+int(strlen)]
	p.pos = p.pos + int(strlen)
	return string(in)
}

/*读取 二进制（二进制的前面包含16位的长度）*/
func (p *Pack) ReadBytes() []byte {
	blen := p.ReadUInt16()
	in := p.data[p.pos : p.pos+int(blen)]
	p.pos = p.pos + int(blen)
	return in
}

/*读取 剩余的二进制（从当前pos开始）*/
func (p *Pack) ReadBytesRemain() []byte {
	in := p.data[p.pos:]
	p.pos = p.Len()
	return in
}

/*-------------------------------------------写入---------------------------------------------*/
/*写入 无符号 8位整数（1字节）*/
func (p *Pack) WriteUInt8(v uint8) {
	by := byte(v)
	p.data = append(p.data, by)
}

/*写入 无符号 16位整数（2字节）*/
//func (p *Pack) WriteUInt16(v uint16) {
//	by := make([]byte, 2)
//	by[1] = byte(v >> 8)
//	by[0] = byte(v)
//	p.data = append(p.data, by...)
//}

/*写入 无符号 16位整数（2字节）*/
func (p *Pack) WriteUInt16(v uint16) {
	by := make([]byte, 2)
	by[0] = byte(v >> 8)
	by[1] = byte(v)
	p.data = append(p.data, by...)
}

/*写入 无符号 32位整数（4字节）*/
func (p *Pack) WriteUInt32(v uint32) {
	by := make([]byte, 4)
	by[0] = byte(v >> 24)
	by[1] = byte(v >> 16)
	by[2] = byte(v >> 8)
	by[3] = byte(v)
	p.data = append(p.data, by...)
}

/*写入 无符号 64位整数（8字节）*/
func (p *Pack) WriteUInt64(v uint64) {
	by := make([]byte, 8)
	by[7] = byte(v >> 56)
	by[6] = byte(v >> 48)
	by[5] = byte(v >> 40)
	by[4] = byte(v >> 32)
	by[3] = byte(v >> 24)
	by[2] = byte(v >> 16)
	by[1] = byte(v >> 8)
	by[0] = byte(v)
	p.data = append(p.data, by...)
}

/*写入 有符号 8位整数（1字节）*/
func (p *Pack) WriteInt8(v int8) {
	p.WriteUInt8(uint8(v))
}

/*写入 有符号 16位整数（2字节）*/
func (p *Pack) WriteInt16(v int16) {
	p.WriteUInt16(uint16(v))
}

/*写入 有符号 32位整数（4字节）*/
func (p *Pack) WriteInt32(v int32) {
	p.WriteUInt32(uint32(v))
}

/*写入 有符号 64位整数（8字节）*/
func (p *Pack) WriteInt64(v int64) {
	p.WriteUInt64(uint64(v))
}

/*写入单精度浮点数*/
func (p *Pack) WriteF32(f float32) {
	by := math.Float32bits(f)
	p.WriteUInt32(uint32(by))
}

/*写入双精度浮点数*/
func (p *Pack) WriteF64(f float64) {
	by := math.Float64bits(f)
	p.WriteUInt64(uint64(by))
}

/*写入 String（string的前面嵌入16位的长度）*/
func (p *Pack) WriteString(v string) {
	by := []byte(v)
	p.WriteUInt16(uint16(len(by)))
	p.data = append(p.data, by...)
}

/*写入 二进制（二进制的前面嵌入16位的长度）*/
func (p *Pack) WriteBytes(v []byte) {
	p.WriteUInt16(uint16(len(v)))
	p.data = append(p.data, v...)
}

/*写入 二进制（不写入长度）*/
func (p *Pack) WriteRawBytes(v []byte) {
	p.data = append(p.data, v...)
}

/*-------------------------------------------其它---------------------------------------------*/
/* 获取对象的 byteArray 值 */
func (p *Pack) Data() []byte {
	return p.data
}

/* 获取对象的 长度 值 */
func (p *Pack) Len() int {
	return len(p.data)
}
func (p *Pack) Clear() {
	p.data = make([]byte, 0)
	p.pos = 0
}
func (p *Pack) Pos() int {
	return p.pos
}
func (p *Pack) Reset() { p.pos = 0 }

/*显示 b的内部结构，以二进制的形式，如 00000000 00001111 00001000 */
func (p *Pack) TraceBytes() {
	b := p.Data()
	fmt.Print("[ ")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%08b ", b[i])
	}
	fmt.Print("]")
}
func TraceBytes2(b []byte) {
	fmt.Print("[ ")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%08b ", b[i])
	}
	fmt.Print("]")
}
