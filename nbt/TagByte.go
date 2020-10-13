package nbt

//TagByte 1
type TagByte struct {
	Tag
	Value byte
}

//NewTagByte return a new TagByte
func NewTagByte(key string, value byte) *TagByte {
	t := new(TagByte)
	t.SetKey(key)
	t.Value = value
	return t
}
