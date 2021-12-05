package Generat_dict

import (
	"bufio"
	"fmt"
	"os"
)

func Generate() {
	var (
		i uint64
	)
	for i = 0; i < 1000000000000000; i++ {
		a := fmt.Sprintf("%016d", i)
		WritePasswd(a)
	}
	fmt.Println("写入第一段")

	for i = 1000000000000000; i < 2000000000000000; i++ {
		a := fmt.Sprintf("%016d", i)
		WritePasswd(a)

	}
	fmt.Println("写入第二段")

	for i = 2000000000000000; i < 3000000000000000; i++ {
		a := fmt.Sprintf("%016d", i)
		WritePasswd(a)
	}
	fmt.Println("写入第三段")
	for i = 3000000000000000; i < 4000000000000000; i++ {
		a := fmt.Sprintf("%016d", i)
		WritePasswd(a)
	}
	for i = 4000000000000000; i < 5000000000000000; i++ {
		a := fmt.Sprintf("%016d", i)
		WritePasswd(a)
	}
	fmt.Println("写入第五段")

	for i = 5000000000000000; i < 6000000000000000; i++ {
		a := fmt.Sprintf("%016d", i)
		WritePasswd(a)
	}
	fmt.Println("写入第六段")
	for i = 6000000000000000; i < 7000000000000000; i++ {
		a := fmt.Sprintf("%016d", i)
		WritePasswd(a)
	}
	fmt.Println("写入第七段")

	for i = 7000000000000000; i < 8000000000000000; i++ {
		a := fmt.Sprintf("%016d", i)
		WritePasswd(a)
	}
	fmt.Println("写入第八段")

	for i = 8000000000000000; i < 9000000000000000; i++ {
		a := fmt.Sprintf("%016d", i)
		WritePasswd(a)
	}
	fmt.Println("写入第九段")

	for i = 9000000000000000; i < 1000000000000000; i++ {
		a := fmt.Sprintf("%016d", i)
		WritePasswd(a)
	}
	fmt.Println("写入第十段")

}

func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
func WritePasswd(num string) {
	filePath := "passwd.dict"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(num)
	write.WriteString("\n")
	write.Flush()
	//for _, v := range nums {
	//	write.WriteString(v)
	//	write.WriteString("\n")
	//	write.Flush()
	//}
	fmt.Printf("条目%v构建完成\n", num)
}
