package main

import (
	"fmt"
)


func main() {
	fmt.Println("Hello World")

	// postingList, classVec := loadDataSet()
    // fmt.Println("Posting List:")
	// for _, row := range postingList {
	// 	fmt.Println(row)
	// }

    // fmt.Println("Class Vector:", classVec)

	// fmt.Println("Data Vec:", createDataVec(postingList))

	// t := setOfWords2Vec(postingList, createDataVec(postingList))

	// fmt.Println("dataVec:")
	// for i, row := range t {
	// 	fmt.Println(row, "[", classVec[i], "]")
	// }

	// p0Vect, p1Vect, pAbusive := trainNB0(t, classVec)
	// fmt.Println("p0Vect:", p0Vect)
	// fmt.Println("p1Vect:", p1Vect)
	// fmt.Println("pAbusive:", pAbusive)

}

// 创建词汇向量表
func createDataVec(data [][]string) []string {
	var dataVec []string
	st := make(map[string]bool)

	for _, row := range data {
		for _, word := range row {
			if _, ok := st[word]; !ok {
				dataVec = append(dataVec, word)
				st[word] = true
				
			}
		}
	}
	return dataVec
}

// 
func contain(wordlist []string, word string) bool {
	for _, temp := range wordlist {
		if temp == word {
			return true
		}	
	}
	return false
}

// 转换
func setOfWords2Vec(data [][]string, vec []string) [][]int {
	var dataVec [][]int
	for _, row := range data {
		var temp []int
		for _, word := range vec {
			if contain(row, word) {
				temp = append(temp, 1)
			} else {
				temp = append(temp, 0)
			}
		}
		dataVec = append(dataVec, temp)
	}
	return dataVec
}

// 训练
func trainNB0(trainMatrix [][]int, trainCategory []int) ([]float64, []float64, float64) {
	numTrainDocs := len(trainMatrix)
	numWords := len(trainMatrix[0])
	pAbusive := float64(sum(trainCategory)) / float64(numTrainDocs)
	p0Num := make([]float64, numWords)  // 存放当 label 为 0 时，每个单词的个数
	p1Num := make([]float64, numWords)
	p0Denom := 0.0  // 存放当 label 为 0 时，所有词的个数
	p1Denom := 0.0

	for i := 0; i < numTrainDocs; i++ {
		if trainCategory[i] == 1 {
			for j := 0; j < numWords; j++ {
				p1Num[j] += float64(trainMatrix[i][j])
				p1Denom += float64(trainMatrix[i][j])
			}
		} else {
			for j := 0; j < numWords; j++ {
				p0Num[j] += float64(trainMatrix[i][j])
				p0Denom += float64(trainMatrix[i][j])
			}
		}
	}

	p1Vect := make([]float64, numWords)
	p0Vect := make([]float64, numWords)

	for i := 0; i < numWords; i++ {
		p1Vect[i] = p1Num[i] / p1Denom
		p0Vect[i] = p0Num[i] / p0Denom
	}

	return p0Vect, p1Vect, pAbusive
}

func sum(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}


// 创建并加载数据集
func loadDataSet() ([][]string, []int){
    postingList := [][]string{
        {"my", "dog", "has", "flea", "problems", "help", "please"},
        {"maybe", "not", "take", "him", "to", "dog", "park", "stupid"},
        {"my", "dalmation", "is", "so", "cute", "I", "love", "him"},
        {"stop", "posting", "stupid", "worthless", "garbage"},
        {"mr", "licks", "ate", "my", "steak", "how", "to", "stop", "him"},
        {"quit", "buying", "worthless", "dog", "food", "stupid"},
    }
    classVec := []int{0, 1, 0, 1, 0, 1}
    return postingList, classVec
}

