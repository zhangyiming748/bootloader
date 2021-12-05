package main

import (
	"bootloader/Execute"
	"bootloader/Generat_dict"
	"bootloader/useDict"
	"fmt"
)

func init() {
	if !Generat_dict.CheckFileIsExist("passwd.dict") {
		fmt.Println("文件不存在,重新生成字典")
		Generat_dict.Generate()
		//Generat_dict.WritePasswd(passwd)
	}
}

func main() {
	codes := useDict.UseDict("passwd.dict")
	for _, v := range codes {
		Execute.Command(v)
	}
}
