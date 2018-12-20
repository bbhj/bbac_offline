package main

import "github.com/astaxie/beego"

func PrintTest(strInput string) {
	beego.Error("string in print.so is:", strInput)
}

