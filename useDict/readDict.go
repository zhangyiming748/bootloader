package useDict

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func UseDict(src string) []string {
	ret := []string{}
	filepath := src
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return []string{}
	}
	defer file.Close()
	//stat, err := file.Stat()
	//if err != nil {
	//	panic(err)
	//}
	//var size = stat.Size()
	//fmt.Println("file size=", size)
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		ret = append(ret, line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return []string{}
			}
		}
	}
	return ret
}
