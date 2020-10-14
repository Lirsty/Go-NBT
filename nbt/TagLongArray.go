package nbt

//TagLongArray 12
type TagLongArray struct {
	Tag
	Value []int64
}

//NewTagLongArray return a new TagLongArray
func NewTagLongArray(key string, value []int64) *TagLongArray {
	t := new(TagLongArray)
	t.SetKey(key)
	t.Value = value
	return t
}
