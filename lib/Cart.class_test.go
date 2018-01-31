package lib

import (
	"fmt"
	"testing"
)

func TestCart(t *testing.T) {
	var cart = NewCart()
	cart.Add("11", "a", "ia", 10, 1)
	cart.Add("21", "b", "ib", 11, 2)

	var compare = func(a []*goodInfo, expect []*goodInfo) {
		if len(a) != len(expect) {
			t.Fatal("len err")
		}
		for i, g := range a {
			tmp := expect[i]
			s1 := fmt.Sprintf("%#v", g)
			s2 := fmt.Sprintf("%#v", tmp)
			if s1 != s2 {
				t.Errorf("%d, expect %#v but got %#v", i, tmp, g)
			}
		}
	}

	compare(cart.Items(), []*goodInfo{
		&goodInfo{"21", "b", "ib", 11, 2},
		&goodInfo{"11", "a", "ia", 10, 1},
	})

	cart.Add("11", "a", "ia", 10, 2)
	compare(cart.Items(), []*goodInfo{
		&goodInfo{"11", "a", "ia", 10, 3},
		&goodInfo{"21", "b", "ib", 11, 2},
	})

	var n = cart.GetGoodsNum()
	if n != 5 {
		t.Errorf("GetGoodsNum expect %d but got %d", 5, n)
	}

	cart.Del("11")
	compare(cart.Items(), []*goodInfo{
		&goodInfo{"21", "b", "ib", 11, 2},
	})

	var totalMoney = cart.GetTotalMoney()
	if totalMoney != 22 {
		t.Errorf("GetTotalMoney expect %d but got %d", 22, totalMoney)
	}

	cart.Incr("21")
	compare(cart.Items(), []*goodInfo{
		&goodInfo{"21", "b", "ib", 11, 3},
	})

	cart.Decr("21")
	compare(cart.Items(), []*goodInfo{
		&goodInfo{"21", "b", "ib", 11, 2},
	})

	cart.Clear()
	compare(cart.Items(), []*goodInfo{})
}
