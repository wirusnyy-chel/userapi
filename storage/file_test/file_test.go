package file_test

// import (
// 	"testing"

// 	"refactoring/storage/file"
// )

// type (
// 	SomeStr1 struct {
// 		text string
// 	}
// 	SomeStr2 struct {
// 		text string
// 	}
// 	command int
// 	test    struct {
// 		c   command
// 		str interface{}
// 		err bool
// 		res interface{}
// 	}
// )

// const (
// 	read command = iota
// 	write
// )

// var tests = []test{
// 	{c: read, str: SomeStr1{}, err: true},
// 	{c: write, str: SomeStr1{text: "text"}, err: false},
// 	{c: read, str: SomeStr1{}, err: false, res: SomeStr1{text: "text"}},
// 	{c: read, str: SomeStr2{}, err: true, res: SomeStr2{text: "text"}},
// }

// func Test(t *testing.T) {
// 	for _, v := range tests {
// 		if v.c == read {
// 			res := v.str
// 			err := file.Read(&res)
// 			if err != nil {
// 				if !v.err {
// 					t.Errorf("test:%v,\n result: v.err:%v,res:%v,err:%v", v, v.err, res, err)
// 				}
// 				continue
// 			} else if res != v.res {
// 				t.Errorf("test:%v,\n result: res:%v,err:%v", v, res, err)
// 			}
// 		} else {
// 			res := v.str
// 			err := file.Write(&res)
// 			if err != nil {
// 				if !v.err {
// 					t.Errorf("test:%v,\n result: err:%v", v, err)
// 				}
// 			}
// 		}
// 	}
// }
