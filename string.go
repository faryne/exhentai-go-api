package exhentai_go_api

import (
	"fmt"
	"strings"
)

type stringKeyword struct {
	Prefix   string
	Keywords []string
}

func NewStringKeyword(key ...string) PropertyManagement {
	var prefix = ""
	if len(key) > 0 {
		prefix = key[0]
	}

	var i = stringKeyword{
		Prefix:   prefix,
		Keywords: make([]string, 0),
	}
	return &i
}

func (l *stringKeyword) Add(input []interface{}) {
	for _, v := range input {
		if inArray(v, l.Keywords) == false {
			l.Keywords = append(l.Keywords, v.(string))
			continue
		}
	}
}

func (l *stringKeyword) Remove(input []interface{}) {
	for k, v := range input {
		if inArray(v, l.Keywords) == true {
			l.Keywords = append(l.Keywords[:k], l.Keywords[k+1:]...)
			break
		}
	}
}

func (l *stringKeyword) String(builder *strings.Builder) {
	if len(l.Keywords) == 0 {
		builder.WriteString(fmt.Sprintf("%s:%s ", l.Prefix, l.Keywords[0]))
	} else {
		for _, v := range l.Keywords {
			builder.WriteString(fmt.Sprintf("~%s:%s ", l.Prefix, v))
		}
	}
}
