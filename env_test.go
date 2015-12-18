package configenv

import (
	"testing"
)

var configGetListProdTests = []struct {
	Spec string
	Want []string
}{
	{"memc_cache.servers", []string{"110.14.216.18:11211", "120.43.342.226:11211"}},
}
var configGetListTests = []struct {
	Spec string
	Want []string
}{
	{"memc_cache.servers", []string{"1.14.216.18:11211", "2.43.342.226:11211"}},
}

func TestGetList(t *testing.T) {
	configProd := NewEnv("config.yml", "production")

	for _, test := range configGetListProdTests {
		got := configProd.GetList(test.Spec, nil)
		if want := test.Want; !testEqString(got, want) {
			t.Errorf("Get(%q) = %q, want %q", test.Spec, got, want)
		}
	}

	config := NewEnv("config.yml", "")

	for _, test := range configGetListTests {
		got := config.GetList(test.Spec, nil)
		if want := test.Want; !testEqString(got, want) {
			t.Errorf("Get(%q) = %q, want %q", test.Spec, got, want)
		}
	}

}

func testEqString(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
