package cverify

import (
	"database/sql/driver"
	"encoding/json"
)

type VerificationQuestion struct {
	Question1 *string `json:"question1"`
	Question2 *string `json:"question2"`
	Question3 *string `json:"question3"`
	Answer1   *string `json:"answer1"`
	Answer2   *string `json:"answer2"`
	Answer3   *string `json:"answer3"`
}

// Scan 取出来的时候的数据
func (c *VerificationQuestion) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}

// Value 入库的数据
func (c VerificationQuestion) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
