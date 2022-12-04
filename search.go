package exhentai_go_api

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
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
	categories []CategoryType     `json:"categories"` // work type
	artists    PropertyManagement `json:"artists"`    // artist name
	characters PropertyManagement `json:"characters"` // character name
	cosplayers PropertyManagement `json:"cosplayers"` // cosplayer name
	females    PropertyManagement `json:"females"`    // female
	groups     PropertyManagement `json:"groups"`     // group
	males      PropertyManagement `json:"males"`      // male
	parodies   PropertyManagement `json:"parodies"`   // parody
	uploaders  PropertyManagement `json:"uploaders"`  // uploader
	languages  PropertyManagement `json:"languages"`  // language
	keywords   PropertyManagement `json:"keywords"`   // keywords
	search     url.Values         // querystring to e-hentai
}

type Artwork struct {
	Category string
	Title    string
}
type Response struct {
	Before   string    `json:"before"`
	After    string    `json:"after"`
	Artworks []Artwork `json:"artworks"`
}

func New() *Request {
	var req = Request{}
	req.artists = NewStringKeyword("a")
	req.characters = NewStringKeyword("c")
	req.cosplayers = NewStringKeyword("cos")
	req.females = NewStringKeyword("f")
	req.males = NewStringKeyword("m")
	req.parodies = NewStringKeyword("p")
	req.uploaders = NewStringKeyword("u")
	req.groups = NewStringKeyword("g")
	req.languages = NewStringKeyword("l")
	req.keywords = NewStringKeyword("")

	req.search = url.Values{}
	req.search.Set("inline_set", "dm_l")
	return &req
}

func (r *Request) Artist() PropertyManagement {
	return r.artists
}

func (r *Request) Character() PropertyManagement {
	return r.characters
}

func (r *Request) Cosplayer() PropertyManagement {
	return r.cosplayers
}

func (r *Request) Female() PropertyManagement {
	return r.females
}

func (r *Request) Male() PropertyManagement {
	return r.males
}

func (r *Request) Parody() PropertyManagement {
	return r.parodies
}

func (r *Request) Uploader() PropertyManagement {
	return r.uploaders
}

func (r *Request) Group() PropertyManagement {
	return r.groups
}

func (r *Request) Language() PropertyManagement {
	return r.languages
}

func (r *Request) Keyword() PropertyManagement {
	return r.keywords
}

func (r *Request) Search() (*Response, error) {
	r.getAndParse(SearchEndpoint)

	//h, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(h))
	return &Response{}, nil
}

func (r *Request) SearchFavorite() (*Response, error) {
	return &Response{}, nil
}

func (r *Request) getAndParse(url string) (*Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	q, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	var output = make([]Artwork, 0)
	q.Find(".itg.gltc").Children().Find("tr").Each(func(idx int, s *goquery.Selection) {
		fmt.Println(s.Html())
		if idx > 0 {
			category, _ := s.Find("td.gl1c.glcat").Find("div").Html()
			title, _ := s.Find("td:eq(2)").Html()
			fmt.Println(title)
			output = append(output, Artwork{
				Category: category,
				Title:    title,
			})
		}
	})
	fmt.Println(output)
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
