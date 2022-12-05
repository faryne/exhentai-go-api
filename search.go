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

type Tag struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}
type Artwork struct {
	Category    string `json:"category"`
	Title       string `json:"title"`
	Thumb       string `json:"thumb"`
	PublishTime string `json:"publish_time"`
	Tags        []Tag  `json:"tags"`
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
	return r.getAndParse(SearchEndpoint)

	//h, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(h))
	//return &Response{}, nil
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
	var output []Artwork
	q.Find(".itg.gltc").Children().Find("tr").Each(func(idx int, s *goquery.Selection) {
		if idx == 2 {
			//fmt.Println(s.Html())
		}
		if idx > 0 {
			category, err1 := s.Find("td.gl1c.glcat").Find("div").Html()
			if err1 != nil {
				fmt.Printf("err1: %s \n", err1.Error())
			}
			obj := s.Find("td.gl2c")
			title := obj.Find("div.glthumb > div > img").AttrOr("alt", "aaaa")
			thumb := obj.Find("div.glthumb > div > img").AttrOr("data-src", "bbbb")
			if thumb == "bbbb" {
				thumb = obj.Find("div.glthumb > div > img").AttrOr("src", "bbbb")
			}
			publishTime, err1 := obj.Find("div.glthumb > div > div > div ").Eq(1).Html()
			if err1 != nil {
				fmt.Println(err1)
			}
			tags := make([]Tag, 0)
			obj.Find("td").Eq(3).Find("div").Eq(2).Find("div.gt").Each(func(_ int, s *goquery.Selection) {
				shortTag, _ := s.Html()
				tags = append(tags, Tag{
					Long:  s.AttrOr("title", ""),
					Short: shortTag,
				})
			})
			output = append(output, Artwork{
				Category:    category,
				Title:       title,
				Thumb:       thumb,
				PublishTime: publishTime,
				Tags:        tags,
			})
		}
	})
	//fmt.Println(output)
	return &Response{
		Before:   "",
		After:    "",
		Artworks: output,
	}, nil
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
