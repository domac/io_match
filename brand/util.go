package brand

//自定义哈希函数
func hashBytes(data []byte) uint64 {
	var h uint64 = 14695981039346656037
	for _, c := range data {
		h = (h ^ uint64(c)) * 10245
	}
	return h
}

func lasIndexN(s []byte, n int, c byte) int {
	for i := len(s) - n; i >= 1; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func lasIndexIdx(s []byte, idx int, c byte) int {
	for i := idx - 1; i >= 1; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

//把字符数组转化为无符号整型
func parsebyteToUint64(b []byte) (n uint64) {
	for i := 0; i < len(b); i++ {
		var v byte
		d := b[i]
		v = d - 48
		n *= uint64(10)
		n1 := n + uint64(v)
		n = n1
	}
	return n
}

func parsebyteToInt(b []byte) (n int) {
	lenb := len(b)
	for i := 0; i < lenb; i++ {
		d := b[i]
		v := d - 48
		n *= 10
		n1 := n + int(v)
		n = n1
	}
	return n
}

func combinehashBytes(data []byte, xh int) uint64 {
	var h uint64 = 14695981039346656037
	for _, c := range data {
		h = (h ^ uint64(c)) * 1099511628211
	}
	return h + uint64(xh)
}

func combinehashBytes32(data []byte, xh uint32) uint64 {
	var h uint64 = 14695981039346656037
	for _, c := range data {
		h = (h ^ uint64(c)) * 1099511628211
	}
	return h + uint64(xh)
}
