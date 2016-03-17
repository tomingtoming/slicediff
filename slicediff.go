package slicediff

import (
	"fmt"
	"sort"
)

type Comparerable interface {
	Key() string
	String() string
}

type Diff interface {
	String() string
}

type Del struct {
	Before Comparerable
}

type Mod struct {
	Before Comparerable
	After  Comparerable
}

type Add struct {
	After Comparerable
}

func (d Del) String() string {
	return fmt.Sprintf("[DEL]%s", d.Before.String())
}

func (m Mod) String() string {
	return fmt.Sprintf("[Mod]%s\n--->>%s", m.Before.String(), m.After.String())
}

func (a Add) String() string {
	return fmt.Sprintf("[ADD]%s", a.After.String())
}

func NewDiff(s, d []Comparerable, isSame func(Comparerable, Comparerable) bool) []Diff {
	diffs := []Diff{}
	sMap := keyMap(s)
	dMap := keyMap(d)
	keySet := sortedUnionKeySet(sMap, dMap)
	for _, key := range keySet {
		sValue, sExists := sMap[key]
		dValue, dExists := dMap[key]
		if sExists && dExists {
			if isSame(sValue, dValue) {
				// Do nothing
			} else {
				diffs = append(diffs, Mod{ Before: sValue, After: dValue })
			}
		} else if sExists {
			diffs = append(diffs, Del{ Before: sValue })
		} else {
			diffs = append(diffs, Add{ After: dValue })
		}
	}
	return diffs
}

func keyMap(cs []Comparerable) map[string]Comparerable {
	km := map[string]Comparerable{}
	for _, c := range cs {
		km[c.Key()] = c
	}
	return km
}

func sortedUnionKeySet(ms ...map[string]Comparerable) []string {
	rm := map[string]interface{}{}
	for _, m := range ms {
		for k, _ := range m {
			rm[k] = nil
		}
	}
	ks := []string{}
	for k, _ := range rm {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}
