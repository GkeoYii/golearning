package main

import (
	"sort"
	"strconv"
	"strings"
	"temp/temp/GinDemo"
)

func main() {
	//circularQueue := Constructor(3) // 设置长度为 3
	//fmt.Println(circularQueue)
	//k8sTest()
	//interfaceDemo.Println(interfaceDemo.Cook{})

	GinDemo.RequestTest()

	//GetPodByName("k3s", "kube-system", "traefik-df4ff85d6-l86vk")
}

func exclusiveTime(n int, logs []string) []int {
	type temp struct {
		id, times int
	}

	res := make([]int, n)
	st := []temp{}
	for _, log := range logs {
		s := strings.Split(log, ":")
		id, _ := strconv.Atoi(s[0])
		times, _ := strconv.Atoi(s[2])
		if s[1][0] == 's' {
			if len(st) > 0 {
				res[st[len(st)-1].id] += times - st[len(st)-1].times
				st[len(st)-1].times = times
			}
			st = append(st, temp{id, times})
		} else {
			p := st[len(st)-1]
			st = st[:len(st)-1]
			res[p.id] += times - p.times + 1
			if len(st) > 0 {
				st[len(st)-1].times = times + 1
			}
		}

	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	if depth == 2 {
		root.Left = &TreeNode{val, root.Left, nil}
		root.Right = &TreeNode{val, nil, root.Right}
	} else {
		root.Left = addOneRow(root.Left, val, depth-1)
		root.Right = addOneRow(root.Right, val, depth-1)
	}
	return root
}

func minSubsequence(nums []int) []int {
	sum := sumInt(nums)
	sort.Ints(nums)
	res := []int{}
	for _, v := range nums {
		res = append(res, v)
		if 2*sumInt(res) > sum {
			break
		}
	}

	return res
}

func sumInt(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	return sum
}

func orderlyQueue(s string, k int) string {
	if 1 == k {
		r := s
		for i := 1; i < len(s); i++ {
			s = s[1:] + s[:1]
			if s < r {
				r = s
			}
		}
		return r
	}

	t := []byte(s)
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	return string(t)
}

type MyCircularQueue struct {
	value []int
	used  int
}

func Constructor(k int) MyCircularQueue {
	v := make([]int, k)
	return MyCircularQueue{v, 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.used <= len(this.value)-1 {
		this.value[this.used] = value
		this.used++
		return true
	}
	return false
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.used > 0 {
		temp := make([]int, len(this.value))
		for i := 0; i < this.used-1; i++ {
			temp[i] = this.value[i+1]
		}
		this.value = temp
		this.used--
		return true
	}
	return false
}

func (this *MyCircularQueue) Front() int {
	if this == nil || this.used == 0 {
		return -1
	}
	return this.value[0]
}

func (this *MyCircularQueue) Rear() int {
	if this == nil || this.used == 0 {
		return -1
	}
	return this.value[this.used-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.used == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.used == len(this.value)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
