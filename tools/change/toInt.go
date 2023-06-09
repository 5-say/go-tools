package change

import (
	"fmt"
	"strconv"
)

// ToInt ..
func ToInt[T any](input T) int {
	var (
		t interface{} = input
	)
	switch v := t.(type) {
	case string:
		result, _ := strconv.Atoi(v)
		return result
	default:
		result, _ := strconv.Atoi(fmt.Sprint(v))
		return result
	}
}
