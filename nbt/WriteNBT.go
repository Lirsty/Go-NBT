package nbt

import "fmt"

func (c *TagCompound) WriteNbt() *[]byte {
	var (
		result []byte
	)
	fmt.Print(c.Values[0].ToBytes())
	return &result
}
