package codec

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"testing"
)

func TestIntToBytes(t *testing.T) {
	docCount := 5
	buf, err := BinaryWrite(docCount)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf)
}

func TestBinarySize(t *testing.T) {
	v := []int{1, 2, 3}
	a := binary.Size(v)
	fmt.Println(a)
}

type Person struct {
	Name string
	Age  uint8
}

func TestBinaryByteRead(t *testing.T) {
	// 创建一个字节流，表示一个 Person 结构体的二进制数据
	buf := bytes.NewBuffer([]byte{
		0x41, 0x6c, 0x69, 0x63, 0x65, // Name: "Alice"
		0x1E, // Age: 30
	})

	var p Person

	// 使用 binary.Read 解析二进制数据到结构体 p 中
	err := binary.Read(buf, binary.BigEndian, &p)
	if err != nil {
		fmt.Println("解析错误:", err)
		return
	}

	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func TestBinaryByteWrite(t *testing.T) {
	// 创建一个字节流，表示一个 Person 结构体的二进制数据
	buf := bytes.NewBuffer([]byte{})

	var p = Person{
		Name: "FanOne",
		Age:  22,
	}

	err := binary.Write(buf, binary.LittleEndian, p)
	if err != nil {
		fmt.Println("解析错误:", err)
		return
	}
	var p2 = Person{}
	err = binary.Write(buf, binary.LittleEndian, p2)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("Name: %s, Age: %d\n", p2.Name, p2.Age)
}

func TestGobEncoding(t *testing.T) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	s := &Person{
		Name: "FanOne",
		Age:  22,
	}
	if err := enc.Encode(s); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", buf.Bytes())

	dec := gob.NewDecoder(buf)
	var s2 *Person
	if err := dec.Decode(&s2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(s2)
}