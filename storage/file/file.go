//File package implements Read() and Write() functions to save data to .json file
package file

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"reflect"
)

//Read structure from corresponding .json file
func Read(str interface{}) error {
	f, err := ioutil.ReadFile(getFileName(str))
	if err != nil {
		return err
	}
	err = json.Unmarshal(f, str)
	return err
}

//Write structure to corresponding .json file
func Write(str interface{}) error {
	b, err := json.Marshal(str)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(getFileName(str), b, fs.ModePerm)
	return err
}

//Specifies name of .json file to given structure
func getFileName(str interface{}) string {
	name := reflect.TypeOf(str)
	return name.Elem().Name() + ".json"
}
