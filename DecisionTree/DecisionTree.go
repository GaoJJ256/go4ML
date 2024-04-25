package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello World")

	// 70%的数据作为训练集，30%的数据作为测试集
	trainDataSet, testDataSet, features := loadDataSet("Data.csv", 70)

	// 打印训练数据集
	fmt.Println("Train Data Set:")
	printDataSet(trainDataSet)

	// 打印测试数据集
	fmt.Println("Test Data Set:")
	printDataSet(testDataSet)

	// 打印特征
	fmt.Println("Features:")
	for _, feature := range features {
		fmt.Printf("%s ", feature)
	}
	fmt.Println()
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
