package helper

import (
	"testing"
)

func TestHelper(t *testing.T) {
	a := []string{"aaa", "bbb", "ccc"}
	if !Contains(a, "aaa") {
		t.Error("should be contain aaa")
	}
	if !Contains(a, "bbb") {
		t.Error("should be contain bbb")
	}
	if !Contains(a, "ccc") {
		t.Error("should be contain ccc")
	}
	if Contains(a, "ddd") {
		t.Error("should be contain ddd")
	}
	if Contains(a, "") {
		t.Error("should be contain ddd")
	}
}
