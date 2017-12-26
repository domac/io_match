package brand

import (
	"bufio"
	"github.com/domac/io_match/log"
	"os"
	"strings"
)

//数据文件读入处理
func ReadAndHandle3(sign, dataDisk string) *IoMatchResult {
	log.GetLogger().Infoln("3-read and handle start")
	f, err := os.Open(dataDisk)
	if err != nil {
		return nil
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	for s.Scan() {
		b := s.Bytes()
		index1 := lasIndexN(b, 9, 32)
		index2 := lasIndexIdx(b, index1, 32)
		index3 := lasIndexIdx(b, index2, 32)
		index4 := lasIndexIdx(b, index3, 32)

		//基础数据
		name := b[:index4]
		hashKey := hashBytes(name)
		if xh, ok := BRANDKEYS[hashKey]; ok {
			price := b[index2+1 : index1]
			BRANDDB[xh] += parsebyteToInt(price)
			onlineDate := b[index1+1:]
			combineHashHey := combinehashBytes(onlineDate, xh)

			item := ONLINESMAP[combineHashHey] + 1
			ONLINESMAP[combineHashHey] = item

			dv := dataList[xh]

			if item > dv {
				dataList[xh] = item
			}

			if dv == 0 {
				namedList[xh] = string(name)
			}

			name = name[:0]
			price = price[:0]
			onlineDate = onlineDate[:0]
			name = nil
			price = nil
			onlineDate = nil
		}
	}

	taskResult := ListResult3()
	res := NewIoMatchResult(sign, taskResult)
	log.GetLogger().Infoln("3-read and handle end")
	return res
}

//输出结果
func ListResult3() string {
	values := make([]BrandItem, len(BRANDKEYS), len(BRANDKEYS))

	cid := 0
	for _, idx := range BRANDKEYS {
		d := dataList[idx]
		tv := BRANDDB[idx]
		name := namedList[idx]
		if d > 1 {
			values[cid] = BrandItem{Name: name, TotalValue: tv, DateCount: d, xh: idx}
			cid++
		}
	}
	values = values[:cid]

	quickSort3(values, 0, len(values)-1)

	newCount := cid
	values2 := []BrandItem{}
	for i := 0; i < newCount; i++ {
		vi := values[i]
		values2 = append(values2, vi)
	}
	values2 = compareSortValue3(values2)
	values2 = compareSortXh3(values2)

	res := []string{}
	for i := 0; i < TOPNUM; i++ {
		currentItem := values2[i]
		res = append(res, currentItem.Name)
	}
	return strings.Join(res, ",")
}

func compareSortValue3(arr []BrandItem) []BrandItem {
	lenS := len(arr)
	currentDateCount := arr[0].DateCount
	currentStartIndex := 0
	for i := 1; i < lenS; i++ {
		targetDateCount := arr[i].DateCount
		if currentDateCount > targetDateCount || i == lenS-1 {
			currentDateCount = targetDateCount
			quickSubValueArray3(arr, currentStartIndex, i-1)
			currentStartIndex = i
		}
	}
	return arr
}

func compareSortXh3(arr []BrandItem) []BrandItem {
	lenS := len(arr)
	currentDateCount := arr[0].DateCount
	currentTotalValue := arr[0].TotalValue
	currentStartIndex := 0

	for i := 1; i < lenS; i++ {
		targetDateCount := arr[i].DateCount
		targetTotalValue := arr[i].TotalValue

		if currentDateCount > targetDateCount || currentTotalValue > targetTotalValue || i == lenS-1 {
			currentDateCount = targetDateCount
			currentTotalValue = targetTotalValue
			quickSubXhArray3(arr, currentStartIndex, i-1)
			currentStartIndex = i
		}
	}
	return arr
}

func quickSubValueArray3(arr []BrandItem, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2].TotalValue
		for i <= j {
			for arr[i].TotalValue > key {
				i++
			}
			for arr[j].TotalValue < key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickSubValueArray3(arr, start, j)
		}
		if end > i {
			quickSubValueArray3(arr, i, end)
		}
	}
}

func quickSubXhArray3(arr []BrandItem, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2].xh
		for i <= j {
			for arr[i].xh < key {
				i++
			}
			for arr[j].xh > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickSubXhArray3(arr, start, j)
		}
		if end > i {
			quickSubXhArray3(arr, i, end)
		}
	}
}

func quickSort3(arr []BrandItem, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2].DateCount
		for i <= j {
			for arr[i].DateCount > key {
				i++
			}
			for arr[j].DateCount < key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickSort3(arr, start, j)
		}
		if end > i {
			quickSort3(arr, i, end)
		}
	}

}
