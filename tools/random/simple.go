package random

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

// simpleRandom ..
type simpleRandom struct {
	crossoverMap   map[string]string // 交叉字典
	confusionSlice []string          // 混淆字典
}

// Simple ..
func Simple() simpleRandom {
	return simpleRandom{
		crossoverMap: map[string]string{
			"0": "h", "1": "8", "2": "d", "3": "6", "4": "g", "5": "b", "6": "0", "7": "5", "8": "a",
			"9": "7", "a": "1", "b": "e", "c": "4", "d": "c", "e": "2", "f": "9", "g": "f", "h": "3",
		},
		confusionSlice: []string{
			"j", "k", "l", "m", "n", "q", "r", "t", "v", "w", "x", "y", "z",
		},
	}
}

// Encode .. 数值可逆编码，混淆输出 大写英文 + 数字，解码时不区分大小写
func (s simpleRandom) Encode(number int64, minLength int) (randomCode string) {
	// 进制转换
	encodeStr := strconv.FormatInt(number, len(s.crossoverMap))

	// 执行交叉
	encodeSlice := strings.Split(encodeStr, "")
	for k, v := range encodeSlice {
		encodeSlice[k] = s.crossoverMap[v]
	}

	// 随机构造混淆切片
	length := minLength
	if len(encodeSlice) > minLength {
		length = len(encodeSlice)
	}
	randomCodeSlice := make([]string, 0, length-len(encodeSlice))
	for len(randomCodeSlice) < length-len(encodeSlice) {
		randInt, _ := rand.Int(rand.Reader, big.NewInt(int64(len(s.confusionSlice))))
		randomCodeSlice = append(randomCodeSlice, s.confusionSlice[randInt.Int64()])
	}

	// 合并切片
	randomCodeSlice = append(randomCodeSlice, encodeSlice...)

	// 连接并转大写
	randomCode = strings.ToUpper(strings.Join(randomCodeSlice, ""))

	// 返回最终的 string
	return randomCode
}

// Decode .. 数值可逆解码，不区分大小写
func (s simpleRandom) Decode(randomCode string) (decodeNumber int64) {
	// 转小写
	randomCode = strings.ToLower(randomCode)

	// 排除混淆
	for _, v := range s.confusionSlice {
		randomCode = strings.ReplaceAll(randomCode, v, "")
	}

	// 构造解码用的交叉字典
	decodeMap := map[string]string{}
	for k, v := range s.crossoverMap {
		decodeMap[v] = k
	}

	// 交叉恢复
	decodeSlice := strings.Split(randomCode, "")
	for k, v := range decodeSlice {
		decodeSlice[k] = decodeMap[v]
	}

	// 连接
	decodeStr := strings.Join(decodeSlice, "")

	// 进制转换
	decodeNumber, _ = strconv.ParseInt(decodeStr, len(s.crossoverMap), 64)

	// 返回解码后的 int64
	return decodeNumber
}
