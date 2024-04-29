package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// 打开CSV文件
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 创建一个新的CSV reader
	r := csv.NewReader(f)

	// 读取CSV文件的所有行
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 检查是否有足够的行
	if len(records) < 4 {
		log.Fatal("Not enough rows in the CSV file")
	}

	// 获取第4行的所有列数据
	row := records[3] // 因为索引从0开始，所以第4行是索引3

	// 对比第4行的所有列数据，获取最大值和最小值
	var maxF, minF float64
	var maxI, minI int
	for index, v := range row {
		// 转换为浮点数
		numF, e := strconv.ParseFloat(v, 64)
		if e != nil {
			fmt.Println("解析失败，非浮点数", v)
			return
		}
		if index == 0 {
			maxF = numF
			minF = numF
		}
		if numF > maxF {
			maxF = numF
			maxI = index
		}
		if numF < minF {
			minF = numF
			minI = index
		}
	}
	// 打印第4行的所有列数据
	fmt.Println(row)
	fmt.Println("获取最大列和最小列索引:", maxI, minI)
	// 构建一个2维数组，用于存储指定列的数据
	var data [][]string
	// 获取指定列的数据
	for _, v := range records {
		data = append(data, []string{v[maxI], v[minI]})
	}
	// 保存到新文件中如果存在则覆盖，不存在则创建
	f, err = os.Create("new_" + f.Name())
	if err != nil {
		panic("文件创建失败:" + err.Error())
	}
	err = csv.NewWriter(f).WriteAll(data)
	if err != nil {
		panic("写入CSV文件失败:" + err.Error())
	}
}
