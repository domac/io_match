package brand

import (
	"bufio"
	"github.com/domac/io_match/log"
	"os"
)

//全局常量
const (
	TOPNUM       = 40
	BrandCount   = 30000000
	OnlinesCount = 70000000
)

//全局变量
var (
	BRANDKEYS = make(map[uint64]int, BrandCount)
	BRANDDB   = []int{}

	toplist [TOPNUM]BrandItem
	topMap  = make(map[uint64]int)

	ONLINESMAP = make(map[uint64]uint32, OnlinesCount)
	dataList   = []uint32{}
	namedList  = []string{}
)

//初始化
func InitKeys(dataFile string) error {
	log.GetLogger().Infof("brand file (%s) init start", dataFile)
	f, err := os.Open(dataFile)
	if err != nil {
		return err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	idx := 0
	for s.Scan() {
		b := s.Bytes()
		//逆向切割
		hashKey := hashBytes(b)
		BRANDKEYS[hashKey] = idx
		idx++
	}
	keysLen := idx
	BRANDDB = make([]int, keysLen, keysLen)
	dataList = make([]uint32, keysLen, keysLen)
	namedList = make([]string, keysLen, keysLen)
	for i := 0; i < keysLen; i++ {
		BRANDDB[i] = 0
		dataList[i] = 0
		namedList[i] = ""
	}

	for i := 0; i < TOPNUM; i++ {
		toplist[i] = BrandItem{
			xh: -1,
		}
	}
	log.GetLogger().Infoln("brand file init finish")
	return nil
}

type BrandItem struct {
	Name       string
	HashKey    uint64
	xh         int
	TotalValue int

	DateCount uint32
}

type IoMatchResult struct {
	Sign       string `json:"sign"`
	TaskResult string `json:"taskResult"`
}

func NewIoMatchResult(sign, taskResult string) *IoMatchResult {
	return &IoMatchResult{
		Sign:       sign,
		TaskResult: taskResult,
	}
}

func HandleDiskData(sign, dataDisk string, dataCheckequence int) *IoMatchResult {
	res := NewIoMatchResult(sign, "Unknow CheckQuence")
	if dataCheckequence == 1 {
		res = ReadAndHandle1(sign, dataDisk)
	} else if dataCheckequence == 2 {
		res = ReadAndHandle2(sign, dataDisk)
	} else if dataCheckequence == 3 {
		res = ReadAndHandle3(sign, dataDisk)
	}

	if res == nil {
		res = NewIoMatchResult(sign, "empty data")
	}
	return res
}
