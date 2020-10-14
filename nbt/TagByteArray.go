package nbt

//TagByteArray 7
type TagByteArray struct {
	Tag
	Value []byte
}

//NewTagByteArray return a new TagByteArray
func NewTagByteArray(key string, value []byte) *TagByteArray {
	t := new(TagByteArray)
	t.SetKey(key)
	t.Value = value
	return t
}
