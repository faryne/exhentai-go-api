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
	jsonContent, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(jsonContent))
	//fmt.Println(err)
	//fmt.Println(resp)

}
