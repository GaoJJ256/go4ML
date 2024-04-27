package main

import (
	"fmt"
)


func main() {
	fmt.Println("Hello World")

	postingList, classVec := loadDataSet()
    fmt.Println("Posting List:", postingList)
    fmt.Println("Class Vector:", classVec)

}
// 创建并加载数据集
func loadDataSet() ([][]string, []int) {
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

