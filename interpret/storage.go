package interpret

import "fmt"

type Storage map[string]*Object

func NewStorage() *Storage {
	return &Storage{}
}

func Store(dest *Storage, key string, value *Object) error {
	// 型を導入したら、ここで型チェック
	(*dest)[key] = value
	return nil
}

func Load(src *Storage, key string) (*Object, error) {
	// 指定されたスコープ内を最優先でサーチ
	o, ok := (*src)[key]
	if ok {
		return o, nil
	}

	// 指定されたところになかった場合、グローバルをサーチ
	gO, gOk := (*globalStorage)[key]
	if gOk {
		return gO, nil
	}

	// なかったわ
	return nil, fmt.Errorf("undefined: %s", key)
}
