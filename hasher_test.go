package hasher

import (
	"testing"
)

const PASSWORD string = "]qadа{w!| Ωf!;f௸ [qyt=}dgu \\jua^_hrц)e ;'eп$`h:Ɋчn&hआ -д/zuxd'm- пعгr%vm ()жуrr,\"ckд bt<#чшwuxષn +ס*х>tвrfdʩфr?b@~кc,wddhр.lxnࡨa "
const HASH_OUTPUT_LENGTH int = 43
const SALT_OUTPUT_LENGTH int = 43

func Test(t *testing.T) {
	hash, salt, err := Hash(PASSWORD)
	if err != nil {
		t.Error(err.Error())
	}
	if len(hash) != HASH_OUTPUT_LENGTH {
		t.Errorf("got: %v, want: %v", len(hash), HASH_OUTPUT_LENGTH)
	}
	if len(salt) != SALT_OUTPUT_LENGTH {
		t.Errorf("got: %v, want: %v", len(salt), SALT_OUTPUT_LENGTH)
	}
	match, err := Verify(PASSWORD, hash, salt)
	if err != nil {
		t.Error(err.Error())
	}
	if !match {
		t.Errorf("got: %v, want: %v", match, true)
	}
	match, err = Verify(PASSWORD, hash, "")
	if err != nil {
		t.Error(err.Error())
	}
	if match {
		t.Errorf("got: %v, want: %v", match, false)
	}
}
