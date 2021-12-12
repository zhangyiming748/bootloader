package Generat_dict

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Generate() {
	for a := 0; a < 10; a++ { //1
		sa:=strconv.Itoa(a)
		for b := 0; b < 10; b++ { //2
			sb:=strconv.Itoa(b)
			for c := 0; c < 10; c++ { //3
				sc:=strconv.Itoa(c)
				for d := 0; d < 10; d++ { //4
					sd:=strconv.Itoa(d)
					for e := 0; e < 10; e++ { //5
						se:=strconv.Itoa(e)
						for f := 0; f < 10; f++ { //6
							sf:=strconv.Itoa(f)
							for g := 0; g < 10; g++ { //7
								sg:=strconv.Itoa(g)
								for h := 0; h < 10; h++ { //8
									sh:=strconv.Itoa(h)
									for i := 0; i < 10; i++ { //9
										si:=strconv.Itoa(i)
										for j := 0; j < 10; j++ { //10
											sj:=strconv.Itoa(j)
											for k:=0;k<10;k++{//11
												sk:=strconv.Itoa(k)
												for l:=0;l<10;l++{//12
													sl:=strconv.Itoa(l)
													for m:=0;m<10;m++{//13
														sm:=strconv.Itoa(m)
														for n:=0;n<10;n++{//14
															sn:=strconv.Itoa(n)
															for o:=0;o<10;o++{//15
																so:=strconv.Itoa(o)
																for p:=0;p<10;p++{//16
																	sp:=strconv.Itoa(p)
																	ret:=strings.Join([]string{sa,sb,sc,sd,se,sf,sg,sh,si,sj,sk,sl,sm,sn,so,sp},"")
																	fmt.Printf("生成的数字是%s\n",ret)
																	WritePasswd(ret)
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}

			}
		}
	}
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
