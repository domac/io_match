package brand

import (
	"bufio"
	"github.com/domac/io_match/log"
	"os"
	"strings"
)

//数据文件读入处理
func ReadAndHandle2(sign, dataDisk string) *IoMatchResult {
	log.GetLogger().Infoln("2-read and handle start")
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
			onlineDate := b[index1+1:]
			price := b[index2+1 : index1]
			combineHashHey := combinehashBytes(onlineDate, xh)
			BRANDDB[xh] += parsebyteToInt(price)
			if _, ok := ONLINESMAP[combineHashHey]; !ok {
				ONLINESMAP[combineHashHey] = '1'
				dv := dataList[xh] + 1
				dataList[xh] = dv
				if dv == 1 {
					namedList[xh] = string(name)
				}
			}

			name = name[:0]
			price = price[:0]
			onlineDate = onlineDate[:0]
			name = nil
			price = nil
			onlineDate = nil
		}
	}
	taskResult := ListResult2()
	res := NewIoMatchResult(sign, taskResult)
	log.GetLogger().Infoln("2-read and handle end")
	return res
}

func ListResult2() string {
	values := make([]BrandItem, len(BRANDKEYS), len(BRANDKEYS))
	cid := 0
	for _, idx := range BRANDKEYS {
		d := dataList[idx]
		tv := BRANDDB[idx]
		name := namedList[idx]
		if d > 100 {
			values[cid] = BrandItem{Name: name, TotalValue: tv, DateCount: d, xh: idx}
			cid++
		}
	}
	values = values[:cid]

	quickSort2(values, 0, len(values)-1)

	newCount := TOPNUM + 5
	values2 := []BrandItem{}
	for i := 0; i < newCount; i++ {
		vi := values[i]
		values2 = append(values2, vi)
	}
	values2 = compareSortValue2(values2)
	values2 = compareSortXh2(values2)

	res := []string{}
	for i := 0; i < TOPNUM; i++ {
		currentItem := values2[i]
		res = append(res, currentItem.Name)
	}
	return strings.Join(res, ",")
}

func compareSortValue2(arr []BrandItem) []BrandItem {
	lenS := len(arr)
	currentDateCount := arr[0].DateCount
	currentStartIndex := 0
	for i := 1; i < lenS; i++ {
		targetDateCount := arr[i].DateCount
		if currentDateCount > targetDateCount || i == lenS-1 {
			currentDateCount = targetDateCount
			quickSubValueArray2(arr, currentStartIndex, i-1)
			currentStartIndex = i
		}
	}
	return arr
}

func compareSortXh2(arr []BrandItem) []BrandItem {
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
			quickSubXhArray2(arr, currentStartIndex, i-1)
			currentStartIndex = i
		}
	}
	return arr
}

func quickSubValueArray2(arr []BrandItem, start, end int) {
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
			quickSubValueArray2(arr, start, j)
		}
		if end > i {
			quickSubValueArray2(arr, i, end)
		}
	}
}

func quickSubXhArray2(arr []BrandItem, start, end int) {
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
			quickSubXhArray2(arr, start, j)
		}
		if end > i {
			quickSubXhArray2(arr, i, end)
		}
	}
}

func quickSort2(arr []BrandItem, start, end int) {
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
			quickSort2(arr, start, j)
		}
		if end > i {
			quickSort2(arr, i, end)
		}
	}

}
