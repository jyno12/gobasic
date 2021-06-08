package sort

import (
	"container/list"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"sort"
	"testing"
)

type Node struct {
	Val  int
	Next *Node
}

func BuildLink(vals []int) *Node {
	var r, p *Node
	for _, val := range vals {
		if r == nil {
			r = &Node{Val: val}
			p = r
		} else {
			p.Next = &Node{Val: val}
			p = p.Next
		}
	}
	return r
}

func Travse(r *Node) {
	for r != nil {
		fmt.Printf("%v ", r.Val)
		r = r.Next
	}
	fmt.Println("")
}

func Sort(r *Node, rlen int) {
	if r == nil || rlen == 0 || rlen == 1 {
		return
	}

	Sort(r, rlen/2)
	p := r
	for i := 0; i < rlen/2; i++ {
		p = p.Next
	}
	Sort(p, rlen-rlen/2)
	Merge(r, p, rlen-rlen/2)
}

func swap(r, l *Node) {
	tmp := r.Val
	r.Val = l.Val
	l.Val = tmp
}

func Merge(l, r *Node, lr int) {
	if r == nil || l == nil {
		return
	}

	pl, pr, i := l, r, 0
	for pl != pr && i < lr {
		if pl.Val > pr.Val {
			swap(pl, pr)
		}
		pl = pl.Next
	}

	for pr != nil && pr.Next != nil {
		if pr.Val > pr.Next.Val {
			swap(pr, pr.Next)
		}
		pr = pr.Next
	}

	// Travse(l)
	return
}

func LenR(r *Node) int {
	var i int
	for r != nil {
		i = i + 1
		r = r.Next
	}
	return i
}

func Retravse(r *Node) []int {
	var res []int
	for r != nil {
		res = append(res, r.Val)
		r = r.Next
	}
	return res
}

func TestLinkSort(t *testing.T) {
	for _, tc := range []struct {
		Input []int
		Exp   []int
	}{
		// {[]int{1}, []int{1}},
		// {[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		// {[]int{4, 1, 2, 3}, []int{1, 2, 3, 4}},
		// {[]int{1, 2, 3, 5, 4}, []int{1, 2, 3, 4, 5}},
		// {[]int{1, 2, 3, 9, 101}, []int{1, 2, 3, 9, 101}},
	} {
		// r := BuildLink(tc.Input)
		// Sort(r, LenR(r))
		// assert.Equal(t, tc.Exp, Retravse(r))

		r1 := BuildLink(tc.Input)
		QSort(r1, Last(r1))
		assert.Equal(t, tc.Exp, Retravse(r1))
	}
}

func Last(r *Node) *Node {
	if r == nil {
		panic("")
	}
	cur := r
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}

func QSort(l, r *Node) {
	Travse(l)

	if r == nil || l == nil {
		return
	}

	sh, bh := &Node{}, &Node{}
	smaller, bigger := sh, bh
	base := l
	p := l.Next
	for p != nil && p != r.Next {
		if p.Val > base.Val {
			bigger.Next = p
			bigger = bigger.Next
		} else {
			smaller.Next = p
			smaller = smaller.Next
		}
		p = p.Next
	}
	smaller.Next = base
	base.Next = bh.Next
	l = sh.Next
	if smaller != sh {
		QSort(l, smaller)
	}
	if bigger != bh {
		QSort(base.Next, bigger)
	}
	return
}

func maxProfit(inventory []int, orders int) int {
	const mod = 1_000_000_000 + 7

	sort.Ints(inventory)
	arr := make([]int, 0, len(inventory)+1)
	for i := len(inventory) - 1; i >= 0; i-- {
		arr = append(arr, inventory[i])
	}
	arr = append(arr, 0)

	var i int
	var sum int
	var min int
	var traverseEnd bool
	for i < len(inventory) && !traverseEnd {
		for i < len(inventory) && arr[i] != arr[i+1] {
			i++
		}

		var sameColorBallCount int
		if i < len(inventory) {
			sameColorBallCount = (arr[i] - arr[i+1]) * (i + 1)
		}
		if orders > sameColorBallCount && i < len(inventory) {
			orders = orders - sameColorBallCount
			i++
		} else {
			n := int(math.Floor(float64(orders) / float64(i+1)))
			min = arr[i] - n
			orders -= n * (i + 1)
			traverseEnd = true
		}
	}

	compute := func(num, orders int) int {
		return ((num + num - orders + 1) * orders / 2) % mod
	}

	for i >= 0 {
		sum = sum + compute(arr[i], arr[i]-min)
		i--
	}

	return (sum + (orders*min)%mod) % mod
}

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 14, maxProfit([]int{2, 5}, 4))
	// assert.Equal(t, 19,maxProfit([]int{3,5},6))
	// assert.Equal(t, 110,maxProfit([]int{2,8,4,10,6},20))
	// assert.Equal(t, 21,maxProfit([]int{1000000000},1000000000))
}

func maxSumTwoNoOverlap(A []int, L int, M int) int {
	max := mf(A, L, M)
	reverseMax := mf(A, M, L)
	if reverseMax > max {
		max = reverseMax
	}
	return max
}

func mf(inputArr []int, l, m int) int {
	f := func(inputArr []int, l int) []int {
		arr := make([]int, len(inputArr))
		sum := 0
		currentMax := 0
		for i, v := range inputArr {
			sum = sum + v
			if i >= l {
				sum = sum - inputArr[i-l]
			}
			if currentMax < sum {
				currentMax = sum
			}
			arr[i] = currentMax
		}
		return arr
	}

	reverseArr := make([]int, len(inputArr))
	for j, i := 0, len(inputArr)-1; i >= 0; j, i = j+1, i-1 {
		reverseArr[j] = inputArr[i]
	}
	arrl := f(inputArr, l)
	arrm := f(reverseArr, m)
	reverseArr = make([]int, len(inputArr))
	for j, i := 0, len(inputArr)-1; i >= 0; j, i = j+1, i-1 {
		reverseArr[j] = arrm[i]
	}
	var max int
	for i := l - 1; i < len(inputArr)-m; i++ {
		current := arrl[i] + reverseArr[i+1]
		if current > max {
			max = current
		}
	}

	return max
}

func TestMaxSumTwoNoOverlap(t *testing.T) {
	assert.Equal(t, 20, maxSumTwoNoOverlap([]int{0, 6, 5, 2, 2, 5, 1, 9, 4}, 1, 2))
	assert.Equal(t, 29, maxSumTwoNoOverlap([]int{3, 8, 1, 3, 2, 1, 8, 9, 0}, 3, 2))
	assert.Equal(t, 31, maxSumTwoNoOverlap([]int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}, 4, 3))
}

func circularArrayLoop1(nums []int) bool {
	GetNextIndex := func(i, num, mod int) int {
		v := i + num
		if v < 0 {
			v = v + mod
		}
		return v % mod
	}
	visitedMap := make(map[int]int)
	for i := range nums {
		if _, visited := visitedMap[i]; visited {
			continue
		}

		for true {
			if _, visited := visitedMap[i]; visited {
				break
			}

			index := GetNextIndex(i, nums[i], len(nums))
			if nums[i]*nums[index] < 0 {
				break
			}

			nums[i] = index
			if index == i {
				break
			}
			i = index
		}

		if visitedMap[i] == i {
			continue
		}
	}
	return false
}

func arrayLoop(arr []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			continue
		}
		if _, ok := m[i]; ok {
			continue
		}

		loopArr, ok := gen(arr, i, m)
		if !ok {
			continue
		}
		if len(loopArr) > 1 {
			return true
		}
	}
	return false
}

func gen(arr []int, i int, m map[int]int) ([]int, bool) {
	for true {
		if _, visited := m[i]; visited {
			break
		}

		index := arr[i] + i
		if index < 0 {
			index = index + len(arr)
		}
		index = index % len(arr)
		m[i] = index
		i = index
	}

	l := make([]int, 0, len(arr))
	for true {
		next, ok := m[i]
		if !ok {
			return nil, false
		}
		l = append(l, i)
		if i == next {
			return nil, false
		}
		i = next
	}
	return l, len(l) > 1
}

func TestCircularArrayLoop(t *testing.T) {
	// assert.Equal(t, true,circularArrayLoop([]int{2,-1,1,2,2}))
	// assert.Equal(t, false,circularArrayLoop([]int{-1,2}))
	// assert.Equal(t, false,circularArrayLoop([]int{-2,1,-1,-2,-2}))
	// assert.Equal(t, false,circularArrayLoop([]int{-1,-2,-3,-4,-5}))
	// assert.Equal(t, true,circularArrayLoop([]int{3,1,2}))
	// assert.Equal(t, true,circularArrayLoop([]int{1,1,2}))

	assert.Equal(t, true, arrayLoop([]int{2, -1, 1, 2, 2}))
	assert.Equal(t, false, arrayLoop([]int{-1, 2}))
	assert.Equal(t, false, arrayLoop([]int{-2, 1, -1, -2, -2}))
	assert.Equal(t, false, arrayLoop([]int{-1, -2, -3, -4, -5}))
	assert.Equal(t, true, arrayLoop([]int{3, 1, 2}))
	assert.Equal(t, true, arrayLoop([]int{1, 1, 2}))
	assert.Equal(t, true, arrayLoop([]int{2, 2, 2, 2, 2, 4, 7}))
}

func merge(intervals [][]int) [][]int {
	return nil
}

func robot(command string, obstacles [][]int, x int, y int) bool {
	commandMap, commandX, commandY := func(cmd string) (m map[int]map[int]bool, cmdx int, cmdy int) {
		m = make(map[int]map[int]bool)
		m[0] = map[int]bool{
			0: true,
		}
		for _, c := range cmd {
			if c == 'U' {
				cmdy++
			} else {
				cmdx++
			}
			mm, ok := m[cmdx]
			if !ok {
				mm = map[int]bool{
					cmdy: true,
				}
			}
			mm[cmdy] = true
			m[cmdx] = mm
		}
		return m, cmdx, cmdy
	}(command)

	collision := func(m map[int]map[int]bool, x, y int) bool {
		if x > commandX || y > commandY {
			s := x / commandX
			x = x % commandX
			y = y - s*commandY
		}
		if x == 0 && y == 0 {
			return true
		}

		_, ok := m[x][y]
		return ok
	}

	for _, obstacle := range obstacles {
		obstacleX, obstacleY := obstacle[0], obstacle[1]
		if obstacleX > x || obstacleY > y {
			continue
		}
		if collision(commandMap, obstacleX, obstacleY) {
			return false
		}
	}

	return collision(commandMap, x, y)
}

func TestRobot(t *testing.T) {
	assert.Equal(t, true, robot("URR", [][]int{}, 3, 2))
	assert.Equal(t, false, robot("URR", [][]int{{2, 2}}, 3, 2))
	assert.Equal(t, true, robot("URR", [][]int{{4, 2}}, 3, 2))
	assert.Equal(t, false, robot("RRRUUU", [][]int{{3, 0}}, 3, 3))
	assert.Equal(t, true, robot("URRURRR", [][]int{
		{7, 7}, {0, 5}, {2, 7}, {8, 6}, {8, 7}, {6, 5}, {4, 4}, {0, 3}, {3, 6},
	}, 4915, 1966))
}

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		c := m[v]
		m[v] = c + 1
	}

	mm := make(map[int]int)
	for _, v := range m {
		_, ok := mm[v]
		if !ok {
			mm[v] = 1
		} else {
			return false
		}
	}
	return true
}

func TestList(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	for e := l.Front(); e != nil; e = e.Next() {
		l.PushBack(3)
		fmt.Println(e.Value)
		e = e.Next()
	}
}

func maximumGain(s string, x int, y int) int {
	src := []byte(s)

	if x < y {
		x, y = y, x
		for i, v := range src {
			if v == 'b' {
				src[i] = 'a'
			} else if v == 'a' {
				src[i] = 'b'
			}
		}
	}

	var ret int
	for i := 0; i < len(src); {
		for i < len(src) && src[i] != 'a' && src[i] != 'b' {
			i++
		}
		var ca, cb int
		for i < len(src) && (src[i] == 'a' || src[i] == 'b') {
			if src[i] == 'a' {
				ca++
			} else {
				if ca > 0 {
					ca--
					ret += x
				} else {
					cb++
				}
			}
			i++
		}

		minc := ca
		if cb < minc {
			minc = cb
		}
		ret += minc * y
	}
	return ret
}

func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0, len(words))
	for _, word := range words {
		patternMap := make(map[byte]int32)
		mapped := make(map[int32]bool)
		var isPattern bool = true
		for i, v := range word {
			p, ok := patternMap[pattern[i]]
			if !ok {
				if _, ismap := mapped[v]; ismap {
					isPattern = false
					break
				}
				patternMap[pattern[i]] = v
				mapped[v] = true
				p = v
			}
			if p != v {
				isPattern = false
				break
			}
		}
		if isPattern {
			res = append(res, word)
		}
	}
	return res
}

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	type Node struct {
		x, y int
		dir  int
	}

	var x, y, cur int
	path := make([]Node, len(obstacleGrid)+len(obstacleGrid[0]))

	cur = -1
	if obstacleGrid[0][0] == 0 {
		cur++
		path[cur] = Node{x: 0, y: 0, dir: 0}
	}
	for cur >= 0 {
		node := path[cur]

		if node.x == len(obstacleGrid)-1 && node.y == len(obstacleGrid[0])-1 {
			break
		}

		switch node.dir {
		case 0:
			x, y = node.x+1, node.y
			path[cur].dir = 1
		case 1:
			x, y = node.x, node.y+1
			path[cur].dir = 2
		default:
			cur--
			continue
		}

		if x < 0 || x >= len(obstacleGrid) || y < 0 || y >= len(obstacleGrid[0]) {
			continue
		}

		if obstacleGrid[x][y] == 1 {
			continue
		}

		cur++
		path[cur] = Node{x: x, y: y, dir: 0}
		obstacleGrid[x][y] = 1
	}

	res := make([][]int, 0, len(obstacleGrid)+len(obstacleGrid[0]))
	for index, node := range path {
		if index <= cur {
			res = append(res, []int{node.x, node.y})
		}
	}
	return res
}

func TestPathWithObstacle(t *testing.T) {
	pathWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})

	pathWithObstacles([][]int{
		{0},
	})

	pathWithObstacles([][]int{
		{1},
	})

	pathWithObstacles([][]int{
		{0, 0},
		{1, 0},
	})
}

func circularArrayLoop(nums []int) bool {
	cur := 1050
	var k bool
	getK := func(v int) bool {
		return v > 0
	}

	next := func(i, v int) (int, bool) {
		next := v + i
		for v < 0 {
			v += len(nums)
		}
		next = v % len(nums)
		if v > 1000 {
			return next, true
		}
		return next, true
	}

	for i, v := range nums {
		if v < -1000 || v > 1000 {
			continue
		}
		cur = cur + 1
		k = getK(v)
		j, ok := next(i, v)
		if !ok {
			continue
		}

		if i == j {
			continue
		}
		for {
			if nums[j] == cur {
				return true
			}
			if k && getK(nums[j]) {
				break
			}
			nums[j] = cur
			if j, ok = next(i, v); !ok {
				break
			}
		}
	}
	return false
}

func TestCircularArrayLoop1(t *testing.T) {
	assert.Equal(t, true, circularArrayLoop([]int{2, -1, 1, 2, 2}))
	assert.Equal(t, false, circularArrayLoop([]int{-1, 2}))
	assert.Equal(t, false, circularArrayLoop([]int{-2, 1, -1, -2, -2}))
	assert.Equal(t, false, circularArrayLoop([]int{-1, -2, -3, -4, -5}))
	assert.Equal(t, true, circularArrayLoop([]int{3, 1, 2}))
	assert.Equal(t, true, circularArrayLoop([]int{1, 1, 2}))
	assert.Equal(t, false, circularArrayLoop([]int{2, 2, 2, 2, 2, 4, 7}))
	assert.Equal(t, false, circularArrayLoop([]int{-2, -3, -9}))
}

func minSetSize(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		count := m[v]
		count++
		m[v] = count
	}
	countArr := make([]int, 0, len(m))
	for _, v := range m {
		countArr = append(countArr, v)
	}
	sort.Ints(countArr)
	total := len(arr)
	sub := 0
	count := 0
	for i := len(countArr) - 1; i >= 0; i-- {
		sub = sub + countArr[i]
		count++
		if sub >= total/2 {
			break
		}
	}
	return count
}

func TestMinSetSize(t *testing.T) {
	assert.Equal(t, 2, minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
	assert.Equal(t, 1, minSetSize([]int{7, 7, 7, 7, 7, 7}))
	assert.Equal(t, 1, minSetSize([]int{1, 9}))
	assert.Equal(t, 1, minSetSize([]int{1000, 1000, 3, 7}))
	assert.Equal(t, 5, minSetSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var nodeCount int
	prev, cur := head, head.Next
	for cur != nil {
		if prev == head {
			prev.Next = nil
		}
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
		nodeCount++
	}
	res := make([]int, 0, nodeCount+1)
	for prev != nil {
		res = append(res, prev.Val)
		prev = prev.Next
	}
	return res
}

func TestReversePrint(t *testing.T) {
	assert.Equal(t, []int{2, 3, 1}, reversePrint(&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2,
		Next: nil}}}))
}

type WORDS []string

func (ws WORDS) Len() int {
	return len(ws)
}

func (ws WORDS) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

func (ws WORDS) Less(i, j int) bool {
	return ws[i] < ws[j]
}

func findLongestWord(s string, dictionary []string) string {
	match := func(s, word string) bool {
		var i, j int
		for i, j = 0, 0; i < len(s) && j < len(word); {
			if s[i] == word[j] {
				i++
				j++
			} else {
				i++
			}
		}
		if j == len(word) {
			return true
		}
		return false
	}
	var ans string
	sort.Sort(WORDS(dictionary))
	for _, word := range dictionary {
		if len(word) <= len(ans) || !match(s, word) {
			continue
		}
		if len(ans) < len(word) {
			ans = word
		}
	}
	return ans
}

func minElements(nums []int, limit int, goal int) int {
	var currentSum int
	for _, v := range nums {
		currentSum = currentSum + v
	}
	goal -= currentSum
	if goal < 0 {
		goal = -goal
	}
	var res int
	res = goal / limit
	if goal%limit != 0 {
		res = res + 1
	}
	return res
}
