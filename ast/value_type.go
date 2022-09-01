package ast

type ValueType int

const (
	_ ValueType = iota
	Void
	Float
	Int
	String
	Bool
	List
	Dict

	Named // ユーザーが名前をつけたもの、解析しないと何かわからない
)
