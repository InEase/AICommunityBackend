package Articles

var CATEGORY = map[int]string{
	0: "分类1",
	1: "分类2",
	2: "分类3",
	3: "分类4",
	4: "分类5",
	5: "分类6",
}

func Keys(m map[int]string) []int {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

var CategoryKeys = Keys(CATEGORY)

func IfInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
