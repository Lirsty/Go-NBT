package nbt

//TagShort 2
type TagShort struct {
	Tag
	Value int16
}

//NewTagShort return a new TagShort
func NewTagShort(key string, value int16) *TagShort {
	t := new(TagShort)
	t.SetKey(key)
	t.Value = value
	return t
}
