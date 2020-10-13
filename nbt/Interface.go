package nbt

type TagI interface {
	SetKey(string)
}

type Tag struct {
	key string
}

func NewTag(key string) *Tag {
	t := new(Tag)
	t.SetKey(key)
	return t
}

func (t *Tag) SetKey(key string) {
	t.key = key
}
