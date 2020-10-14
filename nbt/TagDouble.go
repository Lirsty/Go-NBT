package nbt

//TagDouble 6
type TagDouble struct {
	Tag
	Value float64
}

//NewTagDouble return a new TagDouble
func NewTagDouble(key string, value float64) *TagDouble {
	t := new(TagDouble)
	t.SetKey(key)
	t.Value = value
	return t
}
