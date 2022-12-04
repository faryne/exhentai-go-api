package exhentai_go_api

import (
	"fmt"
	"testing"
)

func Test_A(t *testing.T) {
	s := New()
	resp, err := s.Search()
	fmt.Println(err)
	fmt.Println(resp)
}
