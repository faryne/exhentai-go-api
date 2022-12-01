package exhentai_go_api

import (
	"net/url"
)

type Request struct {
	Categories []CategoryType `json:"categories"` // 作品型態
	Artists    []string       `json:"artists"`    // 畫師
	Languages  []LanguageType `json:"languages"`  // 語言
	search     *url.Values    // 查詢字串
}

type Response struct {
}

func New() *Request {
	return &Request{}
}

func (r *Request) AddLanguages(lang ...LanguageType) {

}

func Search() (*Response, error) {
	return &Response{}, nil
}
