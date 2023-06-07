package change

import "strconv"

// ToString ..
func ToString[T any](input T) string {
	var (
		t interface{} = input
	)
	switch v := t.(type) {
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return v
	default:
		return ""
	}
}
