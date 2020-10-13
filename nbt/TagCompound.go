package nbt

//TagCompound 10
type TagCompound struct {
	Tag
	Values []TagI
}

//NewCompound return a new TagCompound
func NewCompound(key string) *TagCompound {
	t := new(TagCompound)
	t.SetKey(key)
	return t
}

//SetValue .
func (c *TagCompound) SetValue(v TagI) {
	c.Values = append(c.Values, v)
}

//SetByte add a TagByte into TagCompound
func (c *TagCompound) SetByte(key string, value byte) {
	t := NewTagByte(key, value)
	c.SetValue(t)
}

//SetShort add a TagShort into TagCompound
func (c *TagCompound) SetShort(key string, value int16) {
	t := NewTagShort(key, value)
	c.SetValue(t)
}
