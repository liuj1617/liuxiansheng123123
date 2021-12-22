package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

//文件块结构体
type filePart struct {
	partSize int64 //文件块字节大小
	offset   int64 //文件块偏移量
}

const mb = 1024 * 1024     //1MB
const BufferSize = 13 * mb //13是11为手机号+换行符 ，大概100M一个缓冲区

/**
  将大文件切割成多个文件块，采用readAt读偏移量的方式将文件切割，并发读取文件块并判断是否存在
*/
func readBuf(filePath, phoneNo string) bool {
	file, err := os.Open(filePath) //打开文件
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	defer file.Close() //程序执行完毕后关闭文件

	fileInfo, err := file.Stat() //获取文件信息（名称、大小、修改时间等）主要获取文件字节大小
	if err != nil {
		fmt.Println(err)
		return false
	}
	fileSize := fileInfo.Size()            // 文件字节大小
	partNum := fileSize / BufferSize       // 大文件字节大小/文件块大小=文件块数量
	fileParts := make([]filePart, partNum) // 创建文件块数组

	//循环赋值给结构体数组
	for i := 0; i < int(partNum); i++ {
		fileParts[i].partSize = BufferSize
		fileParts[i].offset = BufferSize * int64(i)
	}

	//如果文件字节大小 % 文件块字节大小不整除
	if fileSize%BufferSize != 0 {
		c := filePart{
			partSize: fileSize % BufferSize, //取余数得出最后文件块大小
			offset:   int64(BufferSize * partNum),
		}
		partNum++ //文件块数加1
		fileParts = append(fileParts, c)
	}

	isExist := false
	var wg sync.WaitGroup
	wg.Add(int(partNum))
	for i := 0; i < int(partNum); i++ {
		go func(fileParts []filePart, i int, phoneNo string) bool {
			defer wg.Done()
			filePart := fileParts[i]
			buffer := make([]byte, filePart.partSize)
			_, err := file.ReadAt(buffer, filePart.offset)
			if err != nil {
				println(err)
				panic(err)
			}
			//切小块
			smallPartNum := filePart.partSize / 13
			for i := 0; i < int(smallPartNum); i++ {
				start := i * 13
				end := (i + 1) * 13
				phoneNoTmp := string(buffer[start:end])
				if strings.TrimSpace(string(phoneNoTmp)) == phoneNo { //判断此行数据是否和控制太输入相等
					isExist = true
					break
				}
			}
			return isExist
		}(fileParts, i, phoneNo)
	}
	wg.Wait()
	return isExist
}

//按行读文件，传入文件路径和控制台输入字符串，返回是否存在
func readLine(filePath string, phoneNo string) bool {
	f, err := os.Open(filePath) //打开文件
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	defer f.Close() //程序执行完毕后关闭文件

	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	buf := bufio.NewReader(f)
	isExist := false
	for {
		line, _, err := buf.ReadLine()                  //按行读取
		if strings.TrimSpace(string(line)) == phoneNo { //判断此行数据是否和控制太输入相等
			isExist = true
			break
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error %s\n", err)
			break
		}
	}
	return isExist
}

/*
	大文件的内容排重
	大文件里面是都是号码，在控制台输入一个号码，判断这个文件里是否存在
	注：大文件可能超过本机内存
*/
func main() {
	filePath := "C:\\Users\\办公室\\Desktop\\phone_all.txt" //文件路径
	scanner := bufio.NewScanner(os.Stdin)                //标准输入文件
	isExist := false                                     //手机号码是否存在
	readType := 1                                        //读文件方式 0-按行读取，1-并发读取文件块
	fmt.Print("请输入需要查询的手机号码：")
	for scanner.Scan() {
		phoneNo := scanner.Text() //读控制台输入文本
		switch readType {
		case 0:
			start := time.Now()
			isExist = readLine(filePath, phoneNo) //按行读文件内容并判断
			cost := time.Since(start)
			fmt.Printf("程序运行时间：cost=[%s]\n", cost)
			if isExist {
				fmt.Println("该手机号码在文件中存在")
				break
			} else {
				fmt.Println("该手机号码在文件中不存在")
				break
			}
		case 1:
			start := time.Now()
			isExist = readBuf(filePath, phoneNo) //按文件块读文件内容并判断
			cost := time.Since(start)
			fmt.Printf("程序运行时间：cost=[%s]", cost)
			if isExist {
				fmt.Println("该手机号码在文件中存在")
				break
			} else {
				fmt.Println("该手机号码在文件中不存在")
				break
			}
		}
		break
	}
}
