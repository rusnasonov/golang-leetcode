package main

import (
	"github.com/stretchr/testify/assert"
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T){
	tests := []struct{
	input []string
	expected [][]string
  }{
	   {[]string{"eat","tea","tan","ate","nat","bat"}, [][]string{[]string{"eat", "tea", "ate"}, []string{"tan", "nat"}, []string{"bat"}}},
	   {[]string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, [][]string{[]string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}}},
  }
  for _, tt := range tests {
	  t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
		  out := groupAnagrams(tt.input)
		  assert.Equal(t, tt.expected, out) 
	  })
  }
}