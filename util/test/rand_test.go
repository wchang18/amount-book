package test

import "testing"
import "github.com/wchang18/amount-book/util/helper"

func TestRandInt(t *testing.T) {
	t.Log(helper.GetRandInt(6))
}
