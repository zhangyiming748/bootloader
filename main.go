package main

import (
	"bootloader/Generate"
	"fmt"
	"sync"
)



func main() {
	ch:=make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go Generate.Generate(ch,&wg)
	for v:=range ch{
		fmt.Println(v)
	}
	wg.Wait()

	//codes := useDict.UseDict("passwd.dict")
	//for _, v := range codes {
	//	Execute.Command(v)
	//}
}
