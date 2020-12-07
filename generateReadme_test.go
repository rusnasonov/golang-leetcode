package main

import (
	"testing"
)

func TestBody(t *testing.T) {
	leetcode := NewLeetCode()
	tasks, err := leetcode.Tasks()
	if err != nil {
		t.Error(err)
	}
	if len(tasks) == 0 {
		t.Error("task list is empty")
	}
}
