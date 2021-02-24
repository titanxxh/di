package di

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/titanxxh/di/internal/runtime"
)

// Instantiate returns an instance of a given full-package-path type string.
func Instantiate(typ string) (interface{}, reflect.Value, error) {
	if len(typ) == 0 || (len(typ) == 1 && typ[0] == '*') {
		return nil, reflect.Value{}, fmt.Errorf("input format error: %s", typ)
	}
	var isPtr bool
	if typ[0] == '*' {
		isPtr = true
		typ = typ[1:]
	}
	var pkgPath, typeString, prefix, suffix string
	if indexSlash := strings.LastIndex(typ, "/"); indexSlash != -1 {
		prefix = typ[:indexSlash+1]
		suffix = typ[indexSlash+1:]
	} else {
		suffix = typ
	}
	indexPoint := strings.LastIndex(suffix, ".")
	indexLeftBrackets := strings.LastIndex(suffix, ".(")
	indexRightBrackets := strings.LastIndex(suffix, ").")
	if indexLeftBrackets != -1 && indexRightBrackets != -1 && indexRightBrackets > indexLeftBrackets+2 {
		// path name is not same as package name
		pkgPath = prefix + suffix[:indexLeftBrackets]
		typeString = suffix[indexLeftBrackets+2:indexRightBrackets] + suffix[indexPoint:]
	} else {
		// path name is same as package name
		typeString = suffix
		if indexPoint != -1 {
			pkgPath = prefix + suffix[:indexPoint]
		} else {
			// no package at all, primitives: int
			pkgPath = prefix
		}
	}
	v, err := runtime.Instantiate(pkgPath, typeString, isPtr)
	return v.Interface(), v, err
}
