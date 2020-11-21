package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

func main() {
	filePath :="G:\\abc.txt"
	json, _ := ioutil.ReadFile(filePath)
	re := gjson.Get(string(json), "result")
	fmt.Println(re.String())

	result := gjson.Get(string(json), "data")
	 nnn := len(result.Array())
	fmt.Println(nnn)
	result.ForEach(func(key, value gjson.Result) bool {
		//fmt.Println(key.String())

		//nCount := len(value.Raw)
		//nCount := len(value.Array())
		//fmt.Println(nCount)

		fmt.Println("=============start")
		value.ForEach(func(key1, value1 gjson.Result) bool {
			//fmt.Println(key1.String())
			//fmt.Println(value1.String())
			fmt.Println("=============center=============")
			return true
		})
		fmt.Println("=============end")
		return true // keep iterating
	})
}
