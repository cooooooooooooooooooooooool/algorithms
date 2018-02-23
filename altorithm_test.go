package main

import (
	"fmt"
	"testing"
)

//list rotate

func reverse(data []int, start int, end int) {
	for start < end {
		data[start], data[end] = data[end], data[start]
		start++
		end--
	}
}

func rotate(data []int, k int) {
	l := len(data)
	reverse(data, 0, k-1)
	reverse(data, k, l-1)
	reverse(data, 0, l-1)
}

func TestRotate(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 3
	fmt.Printf("before data: %v, k: %d, ", data, k)
	rotate(data, 3)
	fmt.Printf("after data: %v\n", data)
}

//find the largest sum contiguous subarray

func maxSubArraySum(data []int) int {
	sum := 0
	max := 0
	for _, v := range data {
		sum = sum + v
		if sum < 0 {
			sum = 0
		}
		if max < sum {
			max = sum
		}
	}
	return max
}

func TestMaxSubArraySum(t *testing.T) {
	data := []int{1, -2, 3, 4, -4, 6, -14, 8, 2}
	fmt.Printf("data: %v, maxSubArraySum: %d\n", data, maxSubArraySum(data))
}

//tower of hanoi

type toh struct {
	data []int
}

func (t *toh) push(data int) {
	t.data = append(t.data, data)
}

func (t *toh) pop() int {
	l := len(t.data)
	r := t.data[l-1]
	t.data = t.data[0 : l-1]
	return r
}

func move(start *toh, target *toh) {
	target.push(start.pop())
}

func calc(start *toh, moveNum int, help *toh, target *toh) {
	if moveNum < 1 {
		return
	}
	calc(start, moveNum-1, target, help)
	move(start, target)
	showToh()
	calc(help, moveNum-1, start, target)
}

func showToh() {
	fmt.Printf("start: %v, help: %v, target: %v\n", start.data, help.data, target.data)
}

var start = &toh{
	data: make([]int, 0),
}
var help = &toh{
	data: make([]int, 0),
}
var target = &toh{
	data: make([]int, 0),
}

func TestHanoi(t *testing.T) {
	n := 5
	for i := n; i > 0; i-- {
		start.data = append(start.data, i)
	}
	showToh()
	calc(start, n, help, target)
}

//gcd
//too easy ..skip

//find the second largest number in the list
//skip

//all permutations of an integer list
func findAllPermutations(data []int,i int){
	l:=len(data)
	if i==l{
		fmt.Println(data)
		return
	}
	for j:=i;j<l;j++{
		data[i],data[j]=data[j],data[i]
		findAllPermutations(data,i+1)
		data[i],data[j]=data[j],data[i]
	}
}

func TestFindAllPermutations(t *testing.T){
	data:=[]int{1,2,3}
	findAllPermutations(data,0)
}

//distinct
func findAllPermutations_distinct(data []int,i int){
	l:=len(data)
	if i==l{
		fmt.Println(data)
		return
	}
	helpMap:=make(map[int]bool)
	for j:=i;j<l;j++{
		if _,ok:=helpMap[data[j]];ok{
			continue
		}
		data[i],data[j]=data[j],data[i]
		findAllPermutations_distinct(data,i+1)
		data[i],data[j]=data[j],data[i]
		helpMap[data[j]]=true
	}
}

func TestFindAllPermutations_distinct(t *testing.T){
	data:=[]int{1,1,2}
	findAllPermutations_distinct(data,0)
}






