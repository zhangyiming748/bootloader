package Execute

import (
	"fmt"
	"log"
	"os/exec"
)

//fastboot oem unlock 1234567812345678
func Command(code string) {
	defer func() {
		if errorwarning := recover(); errorwarning != nil {
			fmt.Println(errorwarning)
		}
	}()
	cmd := exec.Command("fastboot", "oem", "unlock", code)
	log.Printf("生成的命令是:%s\n", cmd)
	// 命令的错误输出和标准输出都连接到同一个管道
	defer func() {
		if err := recover(); err != nil {
			//恢复错误
			fmt.Printf("fastboot命令执行有错误产生,比如解锁码错误%v\n", err)
		}
	}()
	stdout, err1 := cmd.StdoutPipe()

	cmd.Stderr = cmd.Stdout
	if err1 != nil {
		log.Printf("cmd.StdoutPipe产生的错误:%v\n", err1)
	}
	err2 := cmd.Start()
	if err2 != nil {
		log.Printf("cmd.Run产生的错误:%v\n", err2)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		log.Println(string(tmp))
		if err != nil {
			break
		}
	}
	err3 := cmd.Wait()
	if err3 != nil {
		log.Printf("命令执行中有错误产生:%v\n", err3)
	}

}
