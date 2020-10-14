package nbt

//TagInt 3
type TagInt struct {
	Tag
	Value int32
}

//NewTagInt return a new TagInt
func NewTagInt(key string, value int32) *TagInt {
	t := new(TagInt)
	t.SetKey(key)
	t.Value = value
	return t
}
