package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

type LRUCacheNode struct {
	key, value int
	prev, next *LRUCacheNode
}

type LRUCache struct {
	capacity             int
	mutex                *sync.Mutex
	elementMap           map[int]*LRUCacheNode
	linkHeader, linkTail *LRUCacheNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity:   capacity,
		elementMap: make(map[int]*LRUCacheNode),
		linkHeader: &LRUCacheNode{},
		linkTail:   &LRUCacheNode{},
		mutex:      new(sync.Mutex),
	}
	l.linkHeader.next = l.linkTail
	l.linkTail.prev = l.linkHeader
	return l
}

func (l *LRUCache) Get(key int) int {
	n, ok := l.elementMap[key]
	if !ok {
		return -1
	}
	l.move(n)
	return n.value
}

func (l *LRUCache) Put(key int, value int) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if len(l.elementMap) >= l.capacity {
		node := l.elementMap[key]
		l.del(node)
	}
	l.add(key, value)
}

func (l *LRUCache) del(n *LRUCacheNode) {
	node := n
	if node == nil {
		node = l.linkTail.prev
	}
	node.prev.next = l.linkTail
	l.linkTail.prev = node.prev
	delete(l.elementMap, node.key)
}

func (l *LRUCache) add(key int, value int) {
	if node, ok := l.elementMap[key]; ok {
		node.value = value
		l.elementMap[key] = node
		return
	}

	node := LRUCacheNode{key: key, value: value, next: nil}
	l.elementMap[key] = &node
	node.next = l.linkHeader.next
	node.prev = l.linkHeader
	node.next.prev = &node
	l.linkHeader.next = &node
}

func (l *LRUCache) move(n *LRUCacheNode) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = l.linkHeader.next
	l.linkHeader.next.prev = n
	l.linkHeader.next = n
}

func TestLRUCache(t *testing.T) {
	cache := Constructor(2 /* 缓存容量 */)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 返回
	cache.Put(3, 3)           // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回  3
	fmt.Println(cache.Get(4)) // 返回  4
	// assert.Equal(t, -1,cache.Get(2))
	// cache.Put(2,6)
	// assert.Equal(t, -1, cache.Get(1))
	// cache.Put(1,5)
	// cache.Put(1,2)
	// assert.Equal(t, 2,cache.Get(1))
	// assert.Equal(t, 6,cache.Get(2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	res := make([][]int, 0, 5000)
	path := make([]int, 10000)
	pathSumDfs(root, target, 0, 0, path, &res)
	m := make([][]int, 0, len(res))
	for i, v := range res {
		if i%2 == 0 {
			continue
		}
		m = append(m, v)
	}
	return m
}

func pathSumDfs(root *TreeNode, target, layer int, currentSum int, path []int, res *[][]int) {
	if root == nil {
		fmt.Println(path[0:layer], layer)
		if currentSum != target {
			return
		}
		*res = append(*res, append([]int{}, path[0:layer]...))
		return
	}

	currentSum = currentSum + root.Val

	path[layer] = root.Val
	pathSumDfs(root.Left, target, layer+1, currentSum, path, res)
	pathSumDfs(root.Right, target, layer+1, currentSum, path, res)
}

func TestPathSum(t *testing.T) {
	// assert.Equal(t, [][]int{
	// 	[]int{5,4,11,2},
	// 	[]int{5,8,4,5},
	// }, pathSum(&TreeNode{
	// 	Val:   5,
	// 	Left:  &TreeNode{
	// 		Val:   4,
	// 		Left:  &TreeNode{
	// 			Val:   11,
	// 			Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
	// 			Right: &TreeNode{Val: 7,Left: nil, Right: nil},
	// 		},
	// 		Right: nil,
	// 	},
	// 	Right: &TreeNode{
	// 		Val:   8,
	// 		Left:  &TreeNode{
	// 			Val:   13,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: &TreeNode{
	// 			Val:   4,
	// 			Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
	// 			Right: &TreeNode{Val: 1,Left: nil,Right: nil},
	// 		},
	// 	},
	// },22))

	assert.Equal(t, [][]int{
		[]int{-2, -3},
	}, pathSum(&TreeNode{
		Val:  -2,
		Left: nil,
		Right: &TreeNode{
			Val:   -3,
			Left:  nil,
			Right: nil,
		},
	}, -5))
}

func checkZeroOnes(s string) bool {
	var seq1Count, seq0Count int
	if len(s) == 0 {
		return false
	}

	var curSeq0, curSeq1 int
	cur := s[0]
	if cur == '0' {
		curSeq0++
	} else {
		curSeq1++
	}
	for i := 1; i < len(s); i++ {
		if s[i] != cur {
			cur = s[i]
			if curSeq1 > seq1Count {
				seq1Count = curSeq1
			}
			if curSeq0 > seq0Count {
				seq0Count = curSeq0
			}
			curSeq0, curSeq1 = 0, 0
		}
		if cur == '0' {
			curSeq0++
		} else {
			curSeq1++
		}
	}
	if curSeq1 > seq1Count {
		seq1Count = curSeq1
	}
	if curSeq0 > seq0Count {
		seq0Count = curSeq0
	}

	return seq1Count > seq0Count
}

func totalHammingDistance(nums []int) int {
	k := make([]int, 64)
	for _, v := range nums {
		v := v
		k[0] = k[0] + v&1
		for i := 1; i < 64; i++ {
			v = v >> 1
			k[i] = k[i] + v&1
		}
	}

	var ans int
	for _, v := range k {
		ans = ans + v*(len(nums)-v)
	}
	return ans
}

var (
	inStr = ""
)

func TestTotalHammingDistance(t *testing.T) {
	// fmt.Println(totalHammingDistance(ins))
}

// func canReach(s string, minJump int, maxJump int) bool {
// 	flag:=make([]bool,len(s))
// 	for i:=range s {
// 		if s[i]=='0' {
// 			flag[i] = true
// 		}
// 	}
// 	cur,max := minJump, maxJump
// 	for cur < len(s) && cur <= max {
// 		if !flag[cur] {
//
// 		} else {
// 			curMax:=cur+maxJump
// 			if curMax > max {
// 				max = curMax
// 			}
// 			if max >= len(s)-1 {
// 				return true
// 			}
// 		}
// 		cur++
// 	}
// 	return false
// }

func TestCanReach(t *testing.T) {
	fmt.Println(canReach(inStr, 1, 49999))
	fmt.Println(canReach("00111010", 3, 5))
	fmt.Println(canReach("011010", 2, 3))
	fmt.Println(canReach("0000000000", 8, 8))
}

func canReach(s string, minJump int, maxJump int) bool {
	flag := make([]bool, len(s))
	for i := range s {
		if s[i] == '0' {
			flag[i] = true
		}
	}
	if !flag[len(s)-1] {
		return false
	}

	visited := make([]bool, len(s))
	visited[0] = true
	rear := maxJump
	for i := 0; i < len(s); i++ {
		if !flag[i] || !visited[i] {
			continue
		}
		prev := i + minJump
		if prev >= len(s) {
			continue
		}

		var start int
		if visited[prev] {
			start = rear
		} else {
			start = prev
		}
		for j := start; j < len(s) && j <= i+maxJump; j++ {
			if j == len(s)-1 {
				return true
			}
			visited[j] = true
		}
		rear = i + maxJump
	}
	return false
}

func maxTurbulenceSize(arr []int) int {
	max := maxT(arr, func(res, index, prev, now int) int {
		if index%2 == 1 {
			if prev < now {
				return res + 1
			}
			return 1
		} else {
			if prev > now {
				return res + 1
			}
			return 1
		}
	})
	max1 := maxT(arr, func(res, index, prev, now int) int {
		if index%2 == 1 {
			if prev > now {
				return res + 1
			}
			return 1
		} else {
			if prev < now {
				return res + 1
			}
			return 1
		}
	})
	if max < max1 {
		return max1
	}
	return max
}

func maxT(arr []int, cmp func(res, index, prev, now int) int) int {
	res := make([]int, len(arr))
	for i := range res {
		res[i] = 1
	}

	for i := range arr {
		if i == 0 {
			continue
		}
		res[i] = cmp(res[i-1], i, arr[i-1], arr[i])
	}

	var max int
	for _, v := range res {
		if v > max {
			max = v
		}
	}
	return max
}

func TestMax(t *testing.T) {
	maxTurbulenceSize([]int{4, 8, 12, 16})
}

func TestMapStruct(t *testing.T) {
	type Student struct {
		id   int
		name string
	}
	students := []Student{
		{id: 1, name: "a"},
		{id: 2, name: "b"},
		{id: 3, name: "c"},
	}
	m := make(map[string]*Student)
	for _, student := range students {
		m[student.name] = &student
	}
	fmt.Printf("%+v", m)
}

// func equationsPossible(equations []string) bool {
// 	disRootMap:=make(map[rune]map[rune]bool)
// 	rootMap:=make(map[rune]rune)
// 	for _,eq:=range equations {
// 		left,right:=eq[0],eq[3]
// 		rootLeft,rootRight := findRoot(left, right)
// 		if rootRight == rootLeft {
// 			continue
// 		}
//
// 	}
// }
