package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

/**
统计样本文件中（包含两种VoLTE和移网语音话单格式类型，参见话单格式说明文档）
呼叫类型为主叫的话单
出现最多的被叫号码及其出现次数
*/
func main() {
	//1.打开文件
	filePath := "C:\\Users\\办公室\\Desktop\\phone1.txt" //文件路径
	file, err := os.Open(filePath)                    //打开文件
	if err != nil {
		fmt.Printf("Error %s\n", err)
		panic(err)
	}
	defer func() {
		err = file.Close() //程序执行完毕后关闭文件
		if err != nil {
			fmt.Printf("Error %s\n", err)
			panic(err)
		}
	}()
	buf := bufio.NewReader(file)

	//2.定义map key为被叫号码，value为出现次数
	tempmap := make(map[string]int)

	//3.循环判断是否存在
	for {
		line, _, err := buf.ReadLine() //按行读取
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error %s\n", err)
			panic(err)
		}
		linestring := strings.TrimSpace(string(line))
		linesplit := strings.Split(linestring, ",")
		bztype := linesplit[0]       //业务类型：1.语音业务
		callednumber := linesplit[3] //原始被叫号码

		if bztype == "1" {
			if _, ok := tempmap[callednumber]; ok {
				tempmap[callednumber]++
			} else {
				tempmap[callednumber] = 1
			}
		}
	}

	//比较得出map中value最高的
	start := time.Now().UnixNano()
	var maxPhoneKey string
	var maxCountVal int = 0
	for key, val := range tempmap {
		if val > maxCountVal {
			maxPhoneKey = key
			maxCountVal = val
		}
	}
	end := time.Now().UnixNano()
	fmt.Printf("出现最多的被叫号码：%s ,出现次数：%d ,程序运行时间：cost=[%d]ms\n", maxPhoneKey, maxCountVal, end-start)

	//使用sort.stable方法对map的value出现次数进行排序
	start1 := time.Now().UnixNano()
	mapSorter := NewMapSorter(tempmap)
	sort.Stable(mapSorter)
	end1 := time.Now().UnixNano()
	fmt.Printf("出现最多的被叫号码：%s ,出现次数：%d ,程序运行时间：cost=[%d]ms\n", mapSorter[0].Key, mapSorter[0].Value, end1-start1)

}

type MapSorter []MapItem

func NewMapSorter(m map[string]int) MapSorter {
	ms := make(MapSorter, 0, len(m))
	for k, v := range m {
		ms = append(ms, MapItem{Key: k, Value: v})
	}
	return ms
}

type MapItem struct {
	Key   string
	Value int
}

func (ms MapSorter) Len() int {
	return len(ms)
}

func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

// Less 按键排序
func (ms MapSorter) Less(i, j int) bool {
	return ms[i].Value > ms[j].Value
}
