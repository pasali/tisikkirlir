// Package  provides ...
package tisikkirlir

import (
	"testing"
)

func TestTsk(t *testing.T) {
	result := Tsk("mehmet bAşAl")
	if result != "mihmit bİşİl" {
		t.Error("Expected mihmit bişil got ", result)
	}
}
