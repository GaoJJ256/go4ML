package main

import (
	"fmt"
	"math"
)

var K = 5                   // 设定knn的K
var DataPath = "./Data.csv" // 训练数据路径
var arr, arr_test []Element // 训练集和验证集
var Array []Element         // 样本的全部数据
var rate = float64(0.7)     // 训练数据占全部数据比例

// 自定义训练数据的结构
type Element struct {
	ID     int
	Action float64
	Kiss   float64

	Label string // 标签

	Distance float64 // 距离
}

// 自定义数据的距离计算方法
func GetDistance(A, B *Element) float64 {
	sum := float64(0)
	sum += (A.Action - B.Action) * (A.Action - B.Action)
	sum += (A.Kiss - B.Kiss) * (A.Kiss * B.Kiss)
	return math.Sqrt(sum)
}

// 程序入口
func main() {
	fmt.Println("Hello world")
	// 读取数据，放入Array

	// 分割数据
	SplitData(Array)

	// 对arr_test中的每个数据执行KNN方法
	cnt, i := 0, 0
	for ; i < len(arr_test); i++ {
		cnt += KNN(arr_test[i])
	}
	if i == 0 {
		fmt.Errorf("数据量有问题")
	}
	fmt.Println("--------")
	fmt.Println("正确率：", float64(float64(cnt)/float64(i)))
}

// knn主体
func KNN(A Element) int {
	for i := 0; i < len(arr); i++ {
		arr[i].Distance = GetDistance(&arr[i], &A)
	}

	QuickSort(arr, 0, len(arr)-1)

	mp := make(map[string]int)
	var weight float64 // 权重

	maxV := arr[0].Label
	for i := 0; i < K; {
		mp[arr[i].Label] += 1
		if mp[arr[i].Label] > mp[maxV] {
			maxV = arr[i].Label
		}
	}
	fmt.Printf("ID: %d", A.ID)
	fmt.Println(maxV)   // 标签
	fmt.Println(weight) // 权重

	if maxV == A.Label {
		return 1
	}
	return 0
}

// 分割，按照rate * len(arr) 的数据进行训练，剩余的数据进行验证
func SplitData(Array []Element) {
	func(A []Element) {
		// 将数据顺序随机打乱，这样保证数据集有较好的性质
		//  ...
	}(Array)
	length := int(rate * float64(len(Array)))
	arr = Array[0:length]
	arr_test = Array[length:len(Array)]
}

// 排序
func QuickSort(q []Element, l, r int) {
	if l >= r {
		return
	}
	x := q[l+(r-l)/2].Distance
	i, j := l-1, r+1
	for i < j {
		for {
			i++
			if q[i].Distance >= x {
				break
			}
		}
		for {
			j--
			if q[j].Distance <= x {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	QuickSort(q, l, j)
	QuickSort(q, j+1, r)
}
