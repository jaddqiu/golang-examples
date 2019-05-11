package main

import (
	"log"
	"os"
)

/*
selection-sort
input: a list of intergers
output: the souted list of interger, ascending
algorithm:
    每轮 选择 剩余最小的值和index， 与当前未排序的第一个值进行交换
    重复n轮之后， 队列排序完成
时间复杂度： O(n*n)
空间复杂度： O(1)
*/
func main() {
	var (
		logger = log.New(os.Stdout, "logger: ", log.Ldate|log.Ltime|log.Lshortfile)
	)
	input := []int{3, 11, 1, 9, 4, 6, 7}
	logger.Printf("Original input is:%v\n", input)
	var minIndex int
	for i := range input {
		logger.Printf("start %d sort\n", i+1)
		minIndex = 0
		for j, value := range input[i:] {
			if value < input[i:][minIndex] {
				minIndex = j
			}
		}
		value := input[i]
		input[i] = input[i+minIndex]
		input[i+minIndex] = value
		logger.Printf("minIndex for %d round is %d, min value is %d\n", i+1, i+minIndex, input[i])
		logger.Printf("after %d sort, input is %v \n", i+1, input)
	}
}
