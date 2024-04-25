package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello World")

}

// 计算信息熵
func culEnt(data [][]string) float64 {
	num := len(data)

	labelMap := make(map[string]int)

	for _, temp := range data {
		curLabel := temp[len(temp)-1] // 取得当前的标签
		if _, ok := labelMap[curLabel]; !ok {
			labelMap[curLabel] = 0
		}
		labelMap[curLabel] += 1
	}

	ent := float64(0)

	for _, v := range labelMap {
		prob := float64(v) / float64(num)
		ent -= math.Log2(prob) * prob
	}
	return ent
}

func loadDataSet(filename string, trainPercentage int) ([][]string, [][]string, []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil, nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading header:", err)
		return nil, nil, nil
	}

	features := header[:len(header)-1] // 除了最后一列，其他都是特征

	// 读取所有数据
	var allData [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading record:", err)
			continue
		}
		allData = append(allData, record)
	}

	// 打乱数据顺序
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allData), func(i, j int) {
		allData[i], allData[j] = allData[j], allData[i]
	})

	// 分割数据为训练集和测试集
	totalRecords := len(allData)
	trainRecords := (trainPercentage * totalRecords) / 100
	trainDataSet := allData[:trainRecords]
	testDataSet := allData[trainRecords:]

	return trainDataSet, testDataSet, features
}

// 辅助函数来打印数据集
func printDataSet(dataSet [][]string) {
	for _, record := range dataSet {
		for _, field := range record {
			fmt.Printf("%s ", field)
		}
		fmt.Println()
	}
}
