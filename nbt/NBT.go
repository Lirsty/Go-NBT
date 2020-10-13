package nbt

//NBT .
type NBT struct {
	NBT TagCompound
}

//New NBT
func New(name string) *NBT {
	n := new(NBT)
	n.NBT.SetKey(name)
	return n
}
