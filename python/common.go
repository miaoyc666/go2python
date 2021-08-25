package python

import "reflect"

/*
File name    : common.go
Author       : miaoyc
Create date  : 2021/8/25 5:01 下午
Description  :
*/

func In(slice interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(slice)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}
