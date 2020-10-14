package nbt

//TagLong 4
type TagLong struct {
	Tag
	Value int64
}

//NewTagLong return a new TagLong
func NewTagLong(key string, value int64) *TagLong {
	t := new(TagLong)
	t.SetKey(key)
	t.Value = value
	return t
}
