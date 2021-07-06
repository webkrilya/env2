package envreader

import (
	"fmt"
	"testing"
)

func TestEnv(t *testing.T) {
	x:=ReadEnv()
	for k,v:=range x{
		fmt.Println(k)
		fmt.Println(v)
	}
}
