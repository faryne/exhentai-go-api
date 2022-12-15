package exhentai_go_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
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

type Uploader struct {
	Name string `json:"name"`
}
type Artwork struct {
	Id          int64    `json:"id"`
	Category    string   `json:"category"`
	Title       string   `json:"title"`
	Thumb       string   `json:"thumb"`
	PublishTime string   `json:"publish_time"`
	Tags        []Tag    `json:"tags"`
	Uploader    Uploader `json:"uploader"`
	Pages       int64    `json:"pages"`
}
type Response struct {
	Before     string    `json:"before"`
	After      string    `json:"after"`
	Artworks   []Artwork `json:"artworks"`
	ArtworksV2 ArtworkV2 `json:"artworks_v2"`
}

type ArtworkV2 struct {
	GMetadata []struct {
		Gid          int64  `json:"gid"`
		Token        string `json:"token"`
		ArchiverKey  string `json:"archiver_key"`
		Title        string `json:"title"`
		TitleJpn     string `json:"title_jpn"`
		Category     string `json:"category"`
		Thumb        string `json:"thumb"`
		Uploader     string `json:"uploader"`
		Posted       string `json:"posted"`
		FileCount    string `json:"filecount"`
		FileSize     int64  `json:"filesize"`
		Expunged     bool   `json:"expunged"`
		Rating       string `json:"rating"`
		TorrentCount string `json:"torrentcount"`
		Torrents     []struct {
			Hash  string `json:"hash"`
			Added string `json:"added"`
			Name  string `json:"name"`
			TSize string `json:"tsize"`
			FSize string `json:"fsize"`
		} `json:"torrents"`
		Tags      []string `json:"tags"`
		ParentGid string   `json:"parent_gid"`
		ParentKey string   `json:"parent_key"`
		FirstGid  string   `json:"first_gid"`
		FirstKey  string   `json:"first_key"`
	} `json:"gmetadata"`
}

var gidList = make([][2]interface{}, 0)

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

func (r *Request) getAndParse(fullUrl string) (*Response, error) {
	resp, err := http.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	q, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	//var output []Artwork
	q.Find(".itg.gltc").Children().Find("tr").Each(func(idx int, s *goquery.Selection) {
		attr, _ := s.Find("td.gl3c.glname > a").Eq(0).Attr("href")
		if attr != "" {
			up, _ := url.Parse(attr)
			splitPath := strings.Split(up.Path, "/")
			artworkId, _ := strconv.Atoi(splitPath[2])
			artworkToken := splitPath[3]
			gidList = append(gidList, [2]interface{}{artworkId, artworkToken})
			//fmt.Printf("%+v, %+v\n", splitPath[2], splitPath[3])
			return
		}
	})
	//fmt.Printf("%#v \n", gidList)
	type GetDetailRequest struct {
		Method  string           `json:"method"`
		GidList [][2]interface{} `json:"gidlist"`
	}
	var rContent = GetDetailRequest{GidList: gidList, Method: "gdata"}
	reqBody, _ := json.Marshal(rContent)
	var buf = bytes.NewBuffer(reqBody)
	body, err := http.Post("https://api.e-hentai.org/api.php", "application/json", buf)
	if err != nil {
		fmt.Println("get detail error: ", err.Error())
		return nil, nil
	}
	defer body.Body.Close()
	content, _ := ioutil.ReadAll(body.Body)
	var respOutput ArtworkV2
	json.Unmarshal(content, &respOutput)
	//fmt.Println(string(content))
	return &Response{
		Before:     "",
		After:      "",
		ArtworksV2: respOutput,
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
