package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	fmt.Println("Hello World")
	trainDataSet, testDataSet, features := loadDataSet("Data.csv", 14)

	var remainLabels []string

	tree := creatTree(trainDataSet, features, remainLabels)

	fmt.Println(tree)

	total := 0

	correctNum := 0

	for _, temp := range testDataSet {
		result := classify(tree, features, temp[:len(temp)-1])
		if strings.Compare(result, temp[len(temp)-1]) == 0 {
			correctNum++
		}

		total++
	}

	rate := float64(correctNum) / float64(total) * 100

	fmt.Println("测试集正确率：" + fmt.Sprintf("%.2f", rate) + "%")

}

// 使用tree进行分类
func classify(tree map[string]interface{}, features []string, testVec []string) string {
	// 获取决策树根节点
	var firstStr string
	for k, v := range tree {
		if v == nil {
			return k
		}

		firstStr = k
	}
	root := tree[firstStr].(map[string]interface{})

	featIdx := func(features []string, feature string) int {
		for i, f := range features {
			if f == feature {
				return i
			}
		}
		return -1 // 如果特征未找到，返回 -1
	}(features, firstStr)

	var classLabel string

	for k, v := range root {
		if strings.Compare(testVec[featIdx], k) == 0 {
			if v == nil {
				classLabel = k
			} else {
				classLabel = classify(root[k].(map[string]interface{}), features, testVec)
			}
		}
	}

	return classLabel
}

// 计算信息熵
func calcEnt(data [][]string) float64 {
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

// 获取某一列的所有特征值
func getFeatureList(dataSet [][]string, columnIndex int) []string {
	var featList []string
	for _, temp := range dataSet {
		featList = append(featList, temp[columnIndex])
	}
	return featList
}

// 取子集
func splitDataSet(data [][]string, columnIndex int, temp string) [][]string {
	var subDataSet [][]string
	for _, row := range data {
		if row[columnIndex] == temp {
			subDataSet = append(subDataSet, row)
		}
	}
	return subDataSet
}

// 计算信息增益
func calcInfoGain(dataSet [][]string, baseEntropy float64, columnIndex int) float64 {

	// 获取某一列的所有特征值
	featList := getFeatureList(dataSet, columnIndex)

	// 获取不同的特征值
	uniqueFeatureValues := distinct(featList)

	// 经验条件熵
	newEntropy := 0.0

	// 计算信息增益
	for _, temp := range uniqueFeatureValues {
		subDataSet := splitDataSet(dataSet, columnIndex, temp)
		// 计算子集的概率
		prob := float64(len(subDataSet)) / float64(len(subDataSet))
		// 计算经验条件熵
		newEntropy += prob * calcEnt(subDataSet)
	}

	// 信息增益
	return baseEntropy - newEntropy
}

// 选择最佳特征
func chooseBestFeature(dataSet [][]string) int {
	// 特征数量
	featureNum := len(dataSet[0]) - 1
	// 计算数据集的熵
	baseEntropy := calcEnt(dataSet)
	// 信息增益
	bestInfoGain := 0.0
	// 最优特征的索引值
	bestFeatureIdx := -1

	// 遍历所有特征
	for i := 0; i < featureNum; i++ {
		// 计算信息增益
		infoGain := calcInfoGain(dataSet, baseEntropy, i)
		// 更新信息增益，找到最大的信息增益
		if infoGain > bestInfoGain {
			bestInfoGain = infoGain
			// 记录信息增益最大的特征的索引
			bestFeatureIdx = i
		}
	}

	return bestFeatureIdx
}

// 构建决策树
func creatTree(data [][]string, labels []string, remainFeatures []string) map[string]interface{} {
	classList := getFeatureList(data, len(data[0])-1)

	// 如果当前类别中元素相同，就停止划分
	if len(classList) == func(a []string, b string) int {
		ans := 0
		for _, t := range a {
			if t == b {
				ans++
			}
		}
		return ans
	}(classList, classList[0]) {
		return map[string]interface{}{classList[0]: nil}
	}

	// 只有一个特征，无法构建
	if len(data[0]) == 1 {
		return map[string]interface{}{func(classList []string) string {
			counts := make(map[string]int)
			for _, class := range classList {
				counts[class]++
			}

			var maxCount int
			var maxClass string
			for class, count := range counts {
				if count > maxCount {
					maxCount = count
					maxClass = class
				}
			}
			return maxClass
		}(classList): nil}
	}

	// 选择最优的特征和标签
	bestFeatureIdx := chooseBestFeature(data)
	bestFeatureLabel := labels[bestFeatureIdx]
	remainFeatures = append(remainFeatures, bestFeatureLabel)

	//  根据最优特征的标签生成树
	tree := make(map[string]interface{})

	// 删除已经使用的特征标签
	tar := make([]string, len(labels))
	copy(tar, labels)
	labels = append(tar[:bestFeatureIdx], tar[bestFeatureIdx+1:]...)

	// 获取最优特征中的属性值
	var featValues []string
	for _, temp := range data {
		featValues = append(featValues, temp[bestFeatureIdx])
	}

	// 去重
	uniqueValues := distinct(featValues)

	// 遍历特征创建决策树
	for _, temp := range uniqueValues {
		if _, ok := tree[bestFeatureLabel]; !ok {
			tree[bestFeatureLabel] = make(map[string]interface{})
		}
		tree[bestFeatureLabel].(map[string]interface{})[temp] = creatTree(splitDataSet(data, bestFeatureIdx, temp), labels, remainFeatures)
	}

	return tree
}

// ---- ----- ----- ----- ---- ------ ----- ----- ------ ----- -----
// 加载数据
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

// 去重函数
func distinct(slice []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, v := range slice {
		if encountered[v] == false {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}
