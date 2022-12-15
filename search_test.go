package exhentai_go_api

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_A(t *testing.T) {
	s := New()
	resp, err := s.Search()
	if err != nil {
		t.Fatal(err.Error())
	}
	//fmt.Printf("%#v \n", resp.ArtworksV2)
	c, _ := json.Marshal(resp.ArtworksV2)
	fmt.Println(string(c))
	//fmt.Printf("%v\n", resp)
	//jsonContent, _ := json.MarshalIndent(resp, "", "\t")
	//fmt.Println(string(jsonContent))
	//fmt.Println(err)
	//fmt.Println(resp)

}
