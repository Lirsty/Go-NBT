package nbt

type TagI interface {
	SetKey(string)
	ToBytes() *[]byte
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

func (t *Tag) ToBytes() *[]byte {
	return &[]byte{}
}

//0~12
const (
	TypeEnd = iota
	TypeByte
	TypeShort
	TypeInt
)
