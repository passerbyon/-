package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func cmd(fileDir string, file string){
	filename := file[0 : len(file)-4]	//去掉后缀
	a := []string{"-i", fileDir + filename + ".mp4", filename + ".mp3"}
	fmt.Println(a)
	Cmdexec := exec.Command("ffmpeg", a...)	//执行cmd命令
	fmt.Printf("正在转换 %v.mp4 >>> %v.mp3\n", filename, filename )
	err := Cmdexec.Run()
	if err != nil{
		fmt.Println("失败",err)
	}else{
		fmt.Println("------转换完成------")
	}
}

func main(){
	var FileDir string
	fmt.Print("文件路径：")
	fmt.Scan(&FileDir)
	//遍历目录，获取文件名
	files, _ := ioutil.ReadDir(FileDir)
	for _, f := range files{
		cmd(FileDir, f.Name())
	}
}