package main

import (
	"fmt"
	"package012/mypackage"
	//"root_folder(go.mod packagename)/pkg_folder"
)

func main() {
	result := mypackage.Add(3, 5)
	//result := pkg_folder's func
	fmt.Println(result)
}
