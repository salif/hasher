package hasher

import (
	"testing"
)

const PASSWORD string = "]qadа{w!| Ωf!;f௸ [qyt=}dgu \\jua^_hrц)e  ;'eп$`h:Ɋчn&hआ -д/zuxd'm- пعгr%vm ()жуrr,\"ckд bt<#чшwuxષn +ס*х>tвrfdʩфr?b@~кc,wddh.lxnࡨa"
const HASH_OUTPUT_LENGTH int = 64
const SALT_OUTPUT_LENGTH int = 32

func Test(t *testing.T) {
	var hash, salt, version = Hash(PASSWORD)
	if len(hash) != HASH_OUTPUT_LENGTH {
		t.Errorf("got: %v, want: %v", len(hash), HASH_OUTPUT_LENGTH)
	}
	if len(salt) != SALT_OUTPUT_LENGTH {
		t.Errorf("got: %v, want: %v", len(salt), SALT_OUTPUT_LENGTH)
	}
	var match = Verify(PASSWORD, hash, salt, version)
	if !match {
		t.Errorf("got: %v, want: %v", match, true)
	}
	var match2 = Verify(PASSWORD, hash, "", version)
	if match2 {
		t.Errorf("got: %v, want: %v", match2, false)
	}
}
