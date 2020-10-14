package nbt

//TagIntArray 11
type TagIntArray struct {
	Tag
	Value []int32
}

//NewTagIntArray return a new TagByteArray
func NewTagIntArray(key string, value []int32) *TagIntArray {
	t := new(TagIntArray)
	t.SetKey(key)
	t.Value = value
	return t
}
