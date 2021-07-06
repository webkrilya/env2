package envreader

import (
	"fmt"
	"testing"
)

func TestEnv(t *testing.T) {
	fmt.Println(ReadEnv())
}
