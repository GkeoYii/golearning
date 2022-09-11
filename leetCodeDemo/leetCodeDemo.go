package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//words := []string{"leetcode", "et", "code"}
	//words := []string{"mass", "as", "hero", "superhero"}
	//fmt.Println(stringMatching(words))

	//i := [][]int{{5, 1}, {4, 2}, {3, 3}, {2, 4}, {1, 5}}
	//j := [][]int{{7, 1}, {6, 2}, {5, 3}, {4, 4}}
	//
	//
	//fmt.Println(mergeSimilarItems(i, j))
	//nums := "-x+5-3+x=6+x-2"
	//eq := switcher(nums)
	//fmt.Println(solveEquation(eq))

	//s := "a0b1c2"
	//fmt.Println(reformat(s))

	//groupSizes := []int{3, 3, 3, 3, 3, 1, 3}
	//fmt.Println(groupThePeople(groupSizes))

	//s := "011101"
	//fmt.Println("max := ", maxScore(s))
	//nums := []int{2, 2, 1, 1, 5, 3, 3, 5}
	//fmt.Println(shuffle(nums, 4))
	//nums := []int{10, 1, 1, 6}
	//fmt.Println(finalPrices(nums))

	//nums := [][]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 1, 1}}
	//fmt.Println(numSpecial(nums))
	//nums := [][]int{{1, 2}, {6, 7}, {2, 3}, {4, 5}}
	//fmt.Println(findLongestChain(nums))

	//text := "this"
	//fmt.Println(reorderSpaces(text))
	//quality = [10,20,5], wage = [70,50,30]

	quality := []int{3, 1, 10, 10, 1}
	wage := []int{4, 8, 2, 2, 7}

	//quality := []int{10, 20, 5}
	//wage := []int{70, 50, 30}

	fmt.Println(mincostToHireWorkers(quality, wage, 3))

}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {

	nums := 1
	for _, value := range wage {
		nums *= value
	}

	type temp struct {
		index int
		value int
	}
	var ts []temp
	for i := 0; i < len(quality); i++ {
		ts = append(ts, temp{i, (nums / wage[i]) * quality[i]})
	}

	sort.Slice(ts,
		func(i, j int) bool {
			return ts[i].value < ts[j].value
		})

	fmt.Println("ts ", ts)
	res := ts[:k]
	sum := 0.0
	sum = float64(wage[res[0].index])
	fmt.Println("sum ", sum)
	x := quality[res[0].index]
	y := wage[res[0].index]
	for i := 1; i < k; i++ {
		sum += float64(y*quality[res[i].index]) / float64(x)
	}
	return float64(sum)
}

func minOperations(logs []string) int {
	res := 0
	for _, v := range logs {
		switch v {
		case "../":
			if res > 0 {
				res--
			}
		case "./":
			continue
		default:
			if res >= 0 {
				res++
			}
		}
	}
	return res
}

func reorderSpaces(text string) string {
	space_counts := strings.Count(text, " ")
	words := strings.Fields(text)

	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", space_counts)
	}
	sl := len(words) - 1
	res := strings.Join(words, strings.Repeat(" ", space_counts/sl)) + strings.Repeat(" ", space_counts%sl)
	return res
}

func reorderSpaces2(s string) (ans string) {
	words := strings.Fields(s)
	space := strings.Count(s, " ")
	lw := len(words) - 1
	if lw == 0 {
		return words[0] + strings.Repeat(" ", space)
	}
	return strings.Join(words, strings.Repeat(" ", space/lw)) + strings.Repeat(" ", space%lw)
}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs,
		func(i, j int) bool {
			fmt.Println(pairs)
			return pairs[i][0] < pairs[j][0]
		})

	fmt.Println(pairs)
	arr := []int{}
	for _, p := range pairs {
		i := sort.SearchInts(arr, p[0])
		if i < len(arr) {
			arr[i] = min(arr[i], p[1])
		} else {
			arr = append(arr, p[1])
		}
	}
	return len(arr)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func numSpecial(mat [][]int) int {

	res_r := make([]int, len(mat))
	res_c := make([]int, len(mat[0]))
	for i, m1 := range mat {
		for j, m2 := range m1 {
			res_r[i] += m2
			res_c[j] += m2
		}
	}
	res := 0
	for i, row := range mat {
		for j, x := range row {
			if x == 1 && res_r[i] == 1 && res_c[j] == 1 {
				res++
			}
		}
	}

	return res
}

func finalPrices(prices []int) []int {

	res := make([]int, len(prices))
	for i, p := range prices {
		res[i] = p
		for j := i + 1; j < len(prices); j++ {
			if p >= prices[j] {
				res[i] = p - prices[j]
				break
			}
		}
	}

	return res
}
func shuffle(nums []int, n int) []int {
	n1 := nums[:n]
	n2 := nums[n:]
	res := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		if i%2 == 0 {
			res[i] = n1[i/2]
		} else {
			res[i] = n2[i/2]
		}
	}
	return res
}

func maxProduct(nums []int) int {

	sort.Ints(nums)
	len := len(nums)
	return (nums[len-1] - 1) * (nums[len-2] - 1)
}

func maxEqualFreq(nums []int) (ans int) {
	freq := map[int]int{}
	count := map[int]int{}
	maxFreq := 0
	for i, num := range nums {
		if count[num] > 0 {
			freq[count[num]]--
		}
		count[num]++
		maxFreq = max(maxFreq, count[num])
		freq[count[num]]++
		if maxFreq == 1 ||
			freq[maxFreq]*maxFreq+freq[maxFreq-1]*(maxFreq-1) == i+1 && freq[maxFreq] == 1 ||
			freq[maxFreq]*maxFreq+1 == i+1 && freq[1] == 1 {
			ans = max(ans, i+1)
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) (sum int) {
	maxLevel := -1
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > maxLevel {
			maxLevel = level
			sum = node.Val
		} else if level == maxLevel {
			sum += node.Val
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	return
}

type OrderedStream struct {
	steam []string
	ptr   int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{make([]string, n+1), 1}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.steam[idKey] = value
	start := this.ptr
	for this.ptr < len(this.steam) && this.steam[this.ptr] != "" {
		this.ptr++
	}
	return this.steam[start:this.ptr]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */

func maxScore(s string) int {
	max := 0
	for i := 1; i < len(s); i++ {
		s0 := s[:i]
		s1 := s[i:]
		temp := strings.Count(s0, "0") + strings.Count(s1, "1")
		if temp > max {
			max = temp
		}
	}
	return max
}

func groupThePeople(groupSizes []int) [][]int {
	m := map[int][]int{}
	for i, v := range groupSizes {
		m[v] = append(m[v], i)

	}
	var res [][]int
	for k, v := range m {
		for i := 0; i < len(v); i += k {
			res = append(res, v[i:i+k])
		}
	}
	return res
}

func groupThePeople2(groupSizes []int) [][]int {
	m := map[int][]int{}
	for i, v := range groupSizes {
		m[v] = append(m[v], i)

	}
	n := 0
	mm := map[int][]int{}
	for k, v := range m {
		for i, j := range v {
			mm[n] = append(mm[n], j)
			if (i+1)%k == 0 {
				n++
			}
		}
	}
	res := make([][]int, len(mm))
	for i, j := range mm {
		res[i] = j
	}
	return res
}

func reformat(s string) string {
	s1 := ""
	s2 := ""
	for _, v := range s {
		if v >= 'a' {
			s1 = s1 + string(v)
		} else {
			s2 = s2 + string(v)
		}
	}
	if math.Abs(float64(len(s1)-len(s2))) > 1 {
		return ""
	}
	res := ""
	if len(s1) > len(s2) {
		for i := 0; i < len(s1)-1; i++ {
			res += string(s1[i]) + string(s2[i])
		}
		res += string(s1[len(s1)-1])
	} else if len(s1) < len(s2) {
		for i := 0; i < len(s2)-1; i++ {
			res += string(s2[i]) + string(s1[i])
		}
		res += string(s2[len(s2)-1])
	} else {
		for i := 0; i < len(s2); i++ {
			res += string(s2[i]) + string(s1[i])
		}
	}
	return res
}

func switcher(equation string) string {
	s := ""
	flag := 0
	for _, j := range equation {
		if j == '=' {
			flag = 1
			s += "-"
			continue
		}
		if flag == 1 {
			if j == '+' {
				j = '-'
			} else if j == '-' {
				j = '+'
			}
		}
		s += string(j)
	}
	s += "+"
	return s
}

func solveEquation(equation string) string {
	equation = switcher(equation)
	//"x+5-3+x-6-x+2"
	x := ""
	y := ""
	sum := ""
	sign := '+'
	for i, v := range equation {
		if v == 'x' {
			if sum == "" {
				sum = "1"
			}
			x = subForString(x, sum, string(sign))
			sum = ""
			continue
		} else if v == '+' || v == '-' {
			if i-1 >= 0 && equation[i-1] != 'x' {
				y = subForString(y, sum, string(sign))
				sum = ""
			}
			sign = v
			continue
		}
		sum += string(v)
	}
	if x == "0" || x == "" {
		if y == "0" || y == "" {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	} else {
		if y == "0" || y == "" {
			return "0"
		}
	}
	x1, _ := strconv.Atoi(x)
	y1, _ := strconv.Atoi(y)
	res := -y1 / x1
	return "x=" + strconv.Itoa(res)
}

func subForString(s1, s2 string, sign string) string {
	x, _ := strconv.Atoi(s1)
	y, _ := strconv.Atoi(s2)
	if sign == "-" {
		return strconv.Itoa(x - y)
	}
	return strconv.Itoa(x + y)
}

func minStartValue2(nums []int) int {
	//value=1-sum
	sum, min := 0, 0
	for _, v := range nums {
		sum += v
		min = minNumber(sum, min)
	}
	return 1 - min
}
func minNumber(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minStartValue(nums []int) int {
	x := 1
	for {
		sum := x
		fmt.Println("x:=", x)
		for i, v := range nums {
			sum += v
			fmt.Println("sum:=", sum)
			if sum <= 0 {
				x++
				break
			}
			if i == len(nums)-1 {
				return x
			}
		}
	}
	return x
}

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}

	subs := sort.StringSlice{}
	cnt, left := 0, 0

	for i, ch := range s {
		if ch == '1' {
			cnt++
		} else if cnt--; cnt == 0 {
			subs = append(subs, "1"+makeLargestSpecial(s[left+1:i])+"0")
			left = i + 1
		}
	}
	sort.Sort(sort.Reverse(subs))
	return strings.Join(subs, "")
}

func tempx(s string) string {
	if len(s) <= 2 {
		return s
	}
	count := 0
	left := 0
	res := sort.StringSlice{}
	for i, v := range s {
		if v == '1' {
			count++
		} else {
			count--
			if count == 0 {
				res = append(res, "1"+makeLargestSpecial(s[left+1:i])+"0")
				left = i + 1
			}
		}
	}
	fmt.Println(res)
	sort.Sort(sort.Reverse(res))
	fmt.Println(res)

	return strings.Join(res, "")
}

func countBadPairs(nums []int) int64 {
	r := 0
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		r += i - m[nums[i]-i]
		m[nums[i]-i]++
	}
	return int64(r)
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {

	m := map[int]int{}
	for i := 0; i < len(items1); i++ {
		m[items1[i][0]] = items1[i][1]
	}
	for j := 0; j < len(items2); j++ {
		m[items2[j][0]] = m[items2[j][0]] + items2[j][1]
	}

	i := 0
	temp := make([]int, len(m))
	for k, _ := range m {
		temp[i] = k
		i++
	}
	sort.Ints(temp)
	i = 0
	res := make([][]int, len(m))
	for _, v := range temp {
		res[i] = []int{v, m[v]}
		i++
	}
	return res
}

func stringMatching(words []string) []string {

	m := map[string]string{}
	for _, v1 := range words {
		for _, v2 := range words {
			if len(v2) > len(v1) && strings.Contains(v2, v1) {
				m[v1] = v1
			}
		}
	}
	s := []string{}
	for k := range m {
		s = append(s, k)
	}
	return s
}

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

func sortedSquares(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	sort.Ints(nums)
	return nums
}

func maximumGroups(grades []int) int {
	m := sum(grades)
	sum := make(map[int]int)
	num := [3][]int{}
	for k, _ := range m {
		count := 0
		for i := 0; i < len(grades); i++ {

			fmt.Println(count, k)
			if count >= k {
				break
			}

			num[k] = append(num[k], grades[i])
			sum[k] = sum[k] + grades[i]

			if sum[k] < sum[k-1] || k != 1 {
				continue
			}
			count++
		}
		fmt.Println("sum: ", sum)
		fmt.Println("num: ", num)
	}

	return 5
}
func sum(grades []int) map[int]int {

	sum := 0
	len := len(grades)
	m := make(map[int]int)
	for i := 1; i < len; i++ {
		m[i] = i
		sum = sum + m[i]
		if sum >= len {
			break
		}
	}
	return m
}

//func minimumOperations(nums []int) int {
//	count := 0
//	if check(nums) {
//		return count
//	}
//	for i := 0; i < len(nums); i++ {
//		temp := min(nums)
//		fmt.Println("temp ", temp)
//		if check(temp) {
//			count = i + 1
//			break
//		}
//	}
//
//	return count
//}
//
//func check(nums []int) bool {
//	for _, i := range nums {
//		if i != 0 {
//			return false
//		}
//	}
//	return true
//}
//
//func min(nums []int) []int {
//	x := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == 0 {
//			continue
//		}
//
//		if x == 0 {
//			x = nums[i]
//		}
//
//		if nums[i] < x {
//			x = nums[i]
//		}
//	}
//
//	for i, _ := range nums {
//		if nums[i] > 0 {
//			nums[i] = nums[i] - x
//		}
//	}
//	return nums
//}
