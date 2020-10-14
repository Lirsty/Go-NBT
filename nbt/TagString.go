package nbt

//TagString 8
type TagString struct {
	Tag
	Value string
}

//NewTagString return a new TagByteArray
func NewTagString(key string, value string) *TagString {
	t := new(TagString)
	t.SetKey(key)
	t.Value = value
	return t
}
