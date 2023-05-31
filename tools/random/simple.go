package random

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

// SimpleRandomConfig ..
type SimpleRandomConfig struct {
	MapShifting       int64 // 字典偏移，任意大于零的随机数
	ConfusionShifting int64 // 混淆偏移，任意大于零的随机数
	ValueShifting     int64 // 数值偏移，任意大于零的随机数
}

// simpleRandom ..
type simpleRandom struct {
	crossoverMap   map[string]string // 交叉字典
	confusionSlice []string          // 混淆字典
	valueShifting  int64             // 数值偏移，任意大于零的随机数
}

// Simple ..
func Simple(config SimpleRandomConfig) simpleRandom {
	var (
		keys   = strings.Split("0123456789abcdefghijklmnopqrstuvwxyz", "")
		values = strings.Split("bd7hij8n3op9ak5l1mef0sct4uvw6gxyz2qr", "")

		shifting        int64 = config.MapShifting % 36        // 字典偏移
		confusionLenght int64 = 8 + config.ConfusionShifting%7 // 混淆长度

		crossoverSlice = append(values[shifting:], values[:shifting]...)[confusionLenght:]
		confusionSlice = append(values[shifting:], values[:shifting]...)[:confusionLenght]
		crossoverMap   = map[string]string{}
	)

	for k, v := range keys[:36-confusionLenght] {
		crossoverMap[v] = crossoverSlice[k]
	}

	return simpleRandom{
		crossoverMap:   crossoverMap,
		confusionSlice: confusionSlice,
		valueShifting:  config.ValueShifting % 51,
	}
}

// Encode .. 数值可逆编码，混淆输出 大写英文 + 数字，解码时不区分大小写
func (s simpleRandom) Encode(number int64, minLength int) (randomCode string) {
	// 进制转换
	encodeStr := strconv.FormatInt(number+s.valueShifting, len(s.crossoverMap))

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
	return decodeNumber - s.valueShifting
}
