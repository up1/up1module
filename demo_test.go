package up1module_test

import (
	"testing"

	"github.com/up1/up1module"
)

func TestSayHiWithSomkiat(t *testing.T) {
	r := up1module.SayHi("somkiat")
	if r != "Say hi somkiat" {
		t.Fatalf("Error	%v", r)
	}
}

func TestSayHiWithEmptyString(t *testing.T) {
	r := up1module.SayHi("")
	if r != "Error" {
		t.Fatalf("Error	%v", r)
	}
}
