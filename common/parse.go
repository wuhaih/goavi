package common

import (
	"github.com/zhangweilun/goxmlpath"
	"unsafe"
)

/**
*
* @author willian
* @created 2017-01-23 17:15
* @email 18702515157@163.com
**/

func PutData(iter *goxmlpath.Path, page *goxmlpath.Node) []string {
	var result []string
	items := iter.Iter(page)
	for items.Next() {
		item := items.Node()
		result = append(result, item.String())
	}
	return result
}

//bytes to string
func ByteString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
