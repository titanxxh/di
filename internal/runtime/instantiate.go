package runtime

import (
	"fmt"
	"reflect"
	"unsafe"
)

//go:linkname typesByString reflect.typesByString
func typesByString(s string) []*_type

func searchForType(pkgPath, typeName string) (*_type, error) {
	tbs := typesByString("*" + typeName)
	for _, tb := range tbs {
		p := (*ptrtype)(unsafe.Pointer(tb))
		path := p.elem.pkgpath()
		if path == pkgPath {
			return &p.typ, nil
		}
	}
	return nil, fmt.Errorf("%s, %s not found in executable image", pkgPath, typeName)
}

func Instantiate(pkgPath, typeString string, isPtr bool) (reflect.Value, error) {
	res, err := searchForType(pkgPath, typeString)
	if err != nil {
		return reflect.Value{}, err
	}
	var emptyFace interface{}
	ptr := (*eface)(unsafe.Pointer(&emptyFace))
	ptr._type = res
	ty := reflect.ValueOf(emptyFace).Type().Elem()
	if isPtr {
		e := reflect.New(ty)
		return e, nil
	}
	e := reflect.New(ty).Elem()
	return e, nil
}
