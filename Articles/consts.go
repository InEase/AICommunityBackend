package Articles

var CATEGORY = map[int]string{
	0: "行业资讯",
	1: "AI论坛",
	2: "订单滚动",
	3: "合作方",
	4: "项目招揽",
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
