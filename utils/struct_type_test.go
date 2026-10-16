package utils_test

import (
	"testing"

	"github.com/nghialthanh/morn-go/utils"
)

func TestIsStructTypeRequiresTheTemplateType(t *testing.T) {
	type User struct{ Name string }
	type Order struct{ Name string }

	if !utils.IsStructType(User{}, &User{Name: "a"}) {
		t.Fatal("pointer to the template type is the same document")
	}
	if !utils.IsStructType(User{}, User{Name: "a"}) {
		t.Fatal("template value matches the document value")
	}
	if utils.IsStructType(User{}, &Order{Name: "a"}) {
		t.Fatal("a different struct is not the template")
	}
	if utils.IsStructType(User{}, (*User)(nil)) {
		t.Fatal("nil document is not the template")
	}
}
