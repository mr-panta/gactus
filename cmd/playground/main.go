package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, 100000)
	fmt.Println(bs)
	fmt.Println(binary.LittleEndian.Uint32(bs))
}
