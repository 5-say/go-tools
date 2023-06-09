package t

import (
	"time"

	"github.com/5-say/go-tools/tools/change"
)

// Fix_MySQL_DATETIME ..
func Fix_MySQL_DATETIME(timeObj time.Time) (fixedTimestamp int64) {
	// 规避 MySQL 数据库 DATETIME 字段类型存储时，毫秒位四舍五入的问题
	originalTimestamp := timeObj.Unix()
	if change.ToInt(timeObj.Format(".000")[1:]) >= 500 {
		return originalTimestamp + 1
	}
	return originalTimestamp
}
