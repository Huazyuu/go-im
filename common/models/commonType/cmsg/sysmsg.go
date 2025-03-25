package cmsg

import (
	"database/sql/driver"
	"encoding/json"
)

type SysMsg struct {
	Type uint8 `json:"type"` // 违规 1:黄 2:恐 3:政 4:不正当言论
}

// Scan 取出来的时候的数据
func (c *SysMsg) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}

// Value 入库的数据
func (c *SysMsg) Value() (driver.Value, error) {
	b, err := json.Marshal(*c)
	return string(b), err
}
