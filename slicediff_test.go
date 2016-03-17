package slicediff

import (
	"testing"
	"fmt"
)

type Elem struct {
	KeyElem      string
	IgnoredField string
	NoticedField string
}

func (e Elem) Key() string {
	return e.KeyElem
}

func (e Elem) String() string {
	return fmt.Sprintf("KEY:%s IGNORE:%s NOTICE:%s", e.Key, e.IgnoredField, e.NoticedField)
}

func TestNewDiff(t *testing.T) {
	src := []Comparerable{
		Elem{
			KeyElem: "DEL",
			IgnoredField: "hogehogehoge",
			NoticedField: "piyopiyopiyo",
		},
		Elem{
			KeyElem: "MOD",
			IgnoredField: "hogehogehoge",
			NoticedField: "piyopiyopiyo",
		},
		Elem{
			KeyElem: "NOCHEANGE",
			IgnoredField: "hogehogehoge",
			NoticedField: "piyopiyopiyo",
		},
		// Key:ADD is nil
	}
	dst := []Comparerable{
		// Key:DEL is Deleted
		Elem{
			KeyElem: "MOD",
			IgnoredField: "hogehogehoge",
			NoticedField: "piyopiyo", // NoticedField is changed
		},
		Elem{
			KeyElem: "NOCHEANGE",
			IgnoredField: "hogehoge", // IgnoredField is changed
			NoticedField: "piyopiyopiyo",
		},
		Elem{
			KeyElem: "ADD",
			IgnoredField: "hogehogehoge",
			NoticedField: "piyopiyopiyo",
		},
	}
	diff := NewDiff(src, dst, func(s, d Comparerable) bool {
		if s.Key() != d.Key() {
			t.Logf("SrcKey: %s", s.Key())
			t.Logf("DstKey: %s", d.Key())
			t.FailNow()
			return false
		} else if sElem, ok := s.(Elem); !ok {
			t.FailNow()
			return false
		} else if dElem, ok := d.(Elem); !ok {
			t.FailNow()
			return false
		} else {
			return sElem.NoticedField == dElem.NoticedField
		}
	})
	fmt.Printf("%#v\n", diff[0])
	fmt.Printf("%#v\n", diff[1])
	fmt.Printf("%#v\n", diff[2])
}
