package analyze

type Comment struct {
	Value string
}

func NewComment(c string) *Comment {
	return &Comment{Value: c}
}
