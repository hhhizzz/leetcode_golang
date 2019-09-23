package _887

import "testing"

func TestSuperEggDrop(t *testing.T) {
    //if superEggDrop(1, 4) != 4 {
    //    t.Errorf("expected: 4")
    //}
    if superEggDrop(2, 100) != 14 {
        t.Errorf("expected: 14")
    }
    if superEggDrop(10, 10000) != 14 {
        t.Errorf("expected: 14")
    }
}
