package ulist

import (
	"github.com/zeromicro/go-zero/core/logx"
	"regexp"
)

func InListByRegex(list []string, key string) (ok bool) {
	for _, s := range list {
		regex, err := regexp.Compile(s)
		if err != nil {
			logx.Error(err)
			return
		}
		if regex.MatchString(key) {
			return true
		}
	}
	return false
}

func InList(list []string, key string) (ok bool) {
	for _, s := range list {
		if s == key {
			return true
		}
	}
	return false
}
