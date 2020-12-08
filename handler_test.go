package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(8, 9)
    if total != 17 {
       t.Errorf("OOps incorrect! got: %d, want: %d.", total, 17)
    }
}