package nbt

//TagList 9
type TagList struct {
	Tag
	Type  int
	Value []TagI
}

//NewTagList return a new TagList
func NewTagList(key string, ValueType int) *TagList {
	t := new(TagList)
	t.SetKey(key)
	t.Type = ValueType
	return t
}
