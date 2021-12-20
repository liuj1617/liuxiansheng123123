package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
	大文件的内容排重
	大文件里面是都是号码，在控制台输入一个号码，判断这个文件里是否存在
	注：大文件可能超过本机内存
*/
func main() {
	//1.bitmap结构初始化
	bitMap := NewBitMap(100000000000)

	//2.打开文件
	filePath := "C:\\Users\\办公室\\Desktop\\phone.txt" //文件路径
	file, err := os.Open(filePath)                   //打开文件
	if err != nil {
		fmt.Printf("Error %s\n", err)
		panic(err)
	}
	defer file.Close() //程序执行完毕后关闭文件

	buf := bufio.NewReader(file)

	//3.将文件内存存入bitmap
	for {
		line, _, err := buf.ReadLine() //按行读取
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error %s\n", err)
			panic(err)
		}
		lineStr := strings.TrimSpace(string(line))
		bitMapInt, err := strconv.ParseUint(lineStr, 10, 64)
		bitMap.Add(uint(bitMapInt))
	}

	//4.判断输入手机号是否在文件中存在
	scanner := bufio.NewScanner(os.Stdin) //标准输入文件
	isExist := false                      //手机号码是否存在
	fmt.Println("请输入需要查重的手机号码：")
	for scanner.Scan() {
		phoneNo := scanner.Text() //读控制台输入文本
		phoneNo1, err := strconv.ParseUint(phoneNo, 10, 64)
		if err != nil {
			fmt.Printf("Error %s\n", err)
			panic(err)
		}
		isExist = bitMap.IsExist(uint(phoneNo1))
		if isExist {
			fmt.Println("该手机号码在文件中存在")
		} else {
			fmt.Println("该手机号码在文件中不存在")
		}
	}
}

//位图
type BitMap struct {
	bits []byte
	max  int
}

//初始化一个BitMap
//一个byte有8位,可代表8个数字,取余后加1为存放最大数所需的容量
func NewBitMap(max int) *BitMap {
	bits := make([]byte, (max>>3)+1)
	return &BitMap{bits: bits, max: max}
}

//添加一个数字到位图
//计算添加数字在数组中的索引index,一个索引可以存放8个数字
//计算存放到索引下的第几个位置,一共0-7个位置
//原索引下的内容与1左移到指定位置后做或运算
func (b *BitMap) Add(num uint) {
	index := num >> 3
	pos := num & 0x07
	b.bits[index] |= 1 << pos
}

//判断一个数字是否在位图
//找到数字所在的位置,然后做与运算
func (b *BitMap) IsExist(num uint) bool {
	index := num >> 3
	pos := num & 0x07
	return b.bits[index]&(1<<pos) != 0
}

//删除一个数字在位图
//找到数字所在的位置取反,然后与索引下的数字做与运算
func (b *BitMap) Remove(num uint) {
	index := num >> 3
	pos := num & 0x07
	b.bits[index] = b.bits[index] & ^(1 << pos)
}

//位图的最大数字
func (b *BitMap) Max() int {
	return b.max
}

func (b *BitMap) String() string {
	return fmt.Sprint(b.bits)
}
