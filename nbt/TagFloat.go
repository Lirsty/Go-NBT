package nbt

//TagFloat 5
type TagFloat struct {
	Tag
	Value float32
}

//NewTagFloat return a new TagFloat
func NewTagFloat(key string, value float32) *TagFloat {
	t := new(TagFloat)
	t.SetKey(key)
	t.Value = value
	return t
}
