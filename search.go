package exhentai_go_api

import (
	"net/url"
	"reflect"
	"strings"
)

type PropertyManagement interface {
	Add(input []interface{})
	Remove(input []interface{})
	String(builder *strings.Builder)
}

type Request struct {
	Categories []CategoryType     `json:"categories"` // work type
	Artists    PropertyManagement `json:"artists"`    // artist name
	Characters PropertyManagement `json:"characters"` // character name
	Cosplayers PropertyManagement `json:"cosplayers"` // cosplayer name
	Females    PropertyManagement `json:"females"`    // female
	Groups     PropertyManagement `json:"groups"`     // group
	Males      PropertyManagement `json:"males"`      // male
	Parodies   PropertyManagement `json:"parodies"`   // parody
	Uploaders  PropertyManagement `json:"uploaders"`  // uploader
	Languages  PropertyManagement `json:"languages"`  // language
	search     *url.Values        // querystring to e-hentai
}

type Response struct {
}

func New() *Request {
	var req = Request{}
	req.Artists = NewStringKeyword("a")
	req.Characters = NewStringKeyword("c")
	req.Cosplayers = NewStringKeyword("cos")
	req.Females = NewStringKeyword("f")
	req.Males = NewStringKeyword("m")
	req.Parodies = NewStringKeyword("p")
	req.Uploaders = NewStringKeyword("u")
	req.Groups = NewStringKeyword("g")
	req.Languages = NewStringKeyword("l")
	return &req
}

func (r *Request) Search() (*Response, error) {
	return &Response{}, nil
}

func inArray(input interface{}, collections interface{}) bool {
	var kind = reflect.TypeOf(collections).Kind()

	switch kind {
	case reflect.Slice:
		s := reflect.ValueOf(collections)

		for i := 0; i < s.Len(); i++ {
			if reflect.ValueOf(input) == reflect.ValueOf(s.Field(i)) {
				return true
			}
		}
	}
	return false
}
