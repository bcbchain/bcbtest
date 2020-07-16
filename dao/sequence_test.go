package dao

import (
	"fmt"
	"testing"
)

func TestGetKindIDSeq(t *testing.T) {
	seq := GetKindIDSeq()
	fmt.Println(seq)
	seq1 := GetKindIDSeq()
	fmt.Println(seq1)
}
