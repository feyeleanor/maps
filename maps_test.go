//	Copyright 2022 Game Design Analytics
//	All Rights Reserved
//
//	Authors:
//		Eleanor McHugh (eleanor@games-with-brains.com)
//
package maps

import (
	"sort"
	"testing"
)

func TestMaps_Len(t *testing.T) {
	ConfirmLen := func(m map[int]bool, l int) {
		if r := Len(m); r != l {
			t.Errorf("Len(%v) should be %v not %v", m, l, r)
		}
	}
	ConfirmLen(nil, 0)
	ConfirmLen(map[int]bool{}, 0)
	ConfirmLen(map[int]bool{0: true}, 1)
	ConfirmLen(map[int]bool{0: true, 1: false}, 2)
	ConfirmLen(map[int]bool{0: true, 1: false, 2: true}, 3)
}

func TestMaps_Get(t *testing.T) {
	ConfirmGet := func(m map[int]bool, k int, v bool) {
		switch r, ok := Get(m, k); {
		case !ok:
			t.Errorf("Get(%v, %v) did not find a value", m, k)
		case r != v:
			t.Errorf("Get(%v, %v) should be %v not %v", m, k, v, r)
		}
	}

	ConfirmGet(map[int]bool{0: true, 1: false, 2: true}, 0, true)
	ConfirmGet(map[int]bool{0: true, 1: false, 2: true}, 1, false)
	ConfirmGet(map[int]bool{0: true, 1: false, 2: true}, 2, true)
}

func TestMaps_Set(t *testing.T) {
	ConfirmSet := func(m map[int]bool, k int, v bool) {
		Set(m, k, v)
		switch r, ok := Get(m, k); {
		case !ok:
			t.Errorf("Set(%v, %v) did not set a value", m, k)
		case r != v:
			t.Errorf("Set(%v, %v) should store %v not %v", m, k, v, r)
		}
	}

	ConfirmSet(map[int]bool{}, 0, true)
	ConfirmSet(map[int]bool{}, 0, false)
	ConfirmSet(map[int]bool{}, 1, true)
	ConfirmSet(map[int]bool{}, 1, false)
	ConfirmSet(map[int]bool{0: true}, 0, false)
	ConfirmSet(map[int]bool{0: false}, 0, true)
	ConfirmSet(map[int]bool{1: true}, 1, false)
	ConfirmSet(map[int]bool{1: false}, 1, true)
}

func TestMaps_Keys(t *testing.T) {
	ConfirmKeys := func(m map[int]bool, k ...int) {
		switch r := Keys(m); {
		case len(m) != len(r):
			t.Errorf("Keys(%v) only returned %v keys of %v", m, len(r), len(m))
		case len(r) != len(k):
			t.Errorf("Keys(%v) wrong number of keys: expected %v rather than %v", m, len(k), len(m))
		default:
			sort.Ints(k)
			sort.Ints(r)
			for i, v := range k {
				if r[i] != v {
					t.Errorf("%v.Keys()[%v] should be %v not %v", m, i, v, r[i])
				}
			}
		}
	}

	ConfirmKeys(nil)
	ConfirmKeys(make(map[int]bool))
	ConfirmKeys(map[int]bool{})
	ConfirmKeys(map[int]bool{0: true}, 0)
	ConfirmKeys(map[int]bool{0: true, 1: true}, 0, 1)
	ConfirmKeys(map[int]bool{0: true, 1: true, 2: true}, 0, 1, 2)
	ConfirmKeys(map[int]bool{0: false, 1: true, 2: true}, 0, 1, 2)
	ConfirmKeys(map[int]bool{0: true, 1: false, 2: true}, 0, 1, 2)
	ConfirmKeys(map[int]bool{0: true, 1: true, 2: false}, 0, 1, 2)
}

func TestMaps_KeysMatch(t *testing.T) {
	ConfirmKeysMatch := func(m1, m2 map[int]bool) {
		switch {
		case !KeysMatch(m1, m2):
			t.Errorf("KeysMatch(%v, %v) should be true", m1, m2)
		case !KeysMatch(m2, m1):
			t.Errorf("KeysMatch(%v, %v) should be true", m2, m1)
		}
	}

	ConfirmKeysMatch(nil, nil)
	ConfirmKeysMatch(map[int]bool{0: true}, map[int]bool{0: true})
	ConfirmKeysMatch(map[int]bool{0: true}, map[int]bool{0: false})
	ConfirmKeysMatch(map[int]bool{0: true, 1: true}, map[int]bool{0: true, 1: false})
	ConfirmKeysMatch(map[int]bool{0: true, 1: true, 2: false}, map[int]bool{0: true, 1: false, 2: false})

	RefuteKeysMatch := func(m1, m2 map[int]bool) {
		switch {
		case KeysMatch(m1, m2):
			t.Errorf("KeysMatch(%v, %v) should be false", m1, m2)
		case KeysMatch(m2, m1):
			t.Errorf("KeysMatch(%v, %v) should be false", m2, m1)
		}
	}

	RefuteKeysMatch(nil, map[int]bool{0: true})
	RefuteKeysMatch(map[int]bool{0: true}, map[int]bool{1: true})
	RefuteKeysMatch(map[int]bool{0: true, 1: true}, map[int]bool{1: true, 2: true})
}

func TestMaps_Copy(t *testing.T) {
	ConfirmCopy := func(m map[int]bool) {
		switch r := Copy(m); {
		case len(m) != len(r):
			t.Errorf("Copy(%v) length should be %v not %v", m, len(m), len(r))
		case !KeysMatch(r, m):
			t.Errorf("Copy(%v) has incorrect keys %v", m, Keys(r))
		default:
			for k, v := range r {
				switch vo, ok := m[k]; {
				case !ok:
					t.Errorf("Copy(%v)[%v] should exist in copy", m, k)
				case vo != v:
					t.Errorf("Copy(%v)[%v] should be %v not %v", m, k, v, vo)
				}
			}
		}
	}

	ConfirmCopy(nil)
	ConfirmCopy(map[int]bool{})
	ConfirmCopy(map[int]bool{0: true})
	ConfirmCopy(map[int]bool{0: true, 1: false})
	ConfirmCopy(map[int]bool{0: true, 1: false, 2: true})
	ConfirmCopy(map[int]bool{0: true, 1: false, 2: true, 3: true})
}

func TestMaps_Merge(t *testing.T) {
	ConfirmMerge := func(r, m map[int]bool, o ...map[int]bool) {
		switch x := Merge(m, o...); {
		case len(r) != len(x):
			t.Errorf("Merge(%v, %v) is the wrong size", m, o)
		case !KeysMatch(x, r):
			t.Errorf("Merge(%v, %v) has incorrect keys", m, o)
		default:
			for k, v := range r {
				switch vo, ok := x[k]; {
				case !ok:
					t.Errorf("Merge(%v, %v)[%v] should exist in copy", m, o, k)
				case vo != v:
					t.Errorf("Merge(%v, %v)[%v] should be %v not %v", m, o, k, v, vo)
				}
			}
		}
	}

	ConfirmMerge(nil, nil, nil)
	ConfirmMerge(map[int]bool{}, nil, map[int]bool{})
	ConfirmMerge(map[int]bool{}, map[int]bool{}, nil)
	ConfirmMerge(map[int]bool{0: true}, nil, map[int]bool{0: true})
	ConfirmMerge(map[int]bool{0: true}, map[int]bool{0: true}, nil)
	ConfirmMerge(map[int]bool{0: true, 1: false}, nil, map[int]bool{0: true, 1: false})
	ConfirmMerge(map[int]bool{0: true, 1: false}, map[int]bool{0: true, 1: false}, nil)

	ConfirmMerge(map[int]bool{0: true, 1: false}, nil, map[int]bool{0: true}, map[int]bool{1: false})
	ConfirmMerge(map[int]bool{0: true, 1: false}, map[int]bool{0: true}, nil, map[int]bool{1: false})
	ConfirmMerge(map[int]bool{0: true, 1: false}, map[int]bool{0: true}, map[int]bool{1: false}, nil)

	ConfirmMerge(map[int]bool{0: true, 1: false, 2: true}, nil, map[int]bool{}, map[int]bool{0: true, 1: false, 2: true})
	ConfirmMerge(map[int]bool{0: true, 1: false, 2: true}, nil, map[int]bool{0: true}, map[int]bool{1: false, 2: true})
	ConfirmMerge(map[int]bool{0: true, 1: false, 2: true}, map[int]bool{0: true}, map[int]bool{1: false}, map[int]bool{2: true})
	ConfirmMerge(map[int]bool{0: true, 1: false, 2: true}, map[int]bool{0: true, 1: false}, map[int]bool{2: true}, nil)
	ConfirmMerge(map[int]bool{0: true, 1: false, 2: true}, map[int]bool{0: true, 1: false, 2: true}, map[int]bool{}, nil)

	ConfirmMerge(map[int]bool{0: true}, map[int]bool{0: true}, map[int]bool{0: true}, map[int]bool{0: true})
	ConfirmMerge(map[int]bool{0: true}, map[int]bool{0: false}, map[int]bool{0: true}, map[int]bool{0: true})
	ConfirmMerge(map[int]bool{0: true}, map[int]bool{0: true}, map[int]bool{0: false}, map[int]bool{0: true})
	ConfirmMerge(map[int]bool{0: false}, map[int]bool{0: true}, map[int]bool{0: true}, map[int]bool{0: false})
}

func TestMaps_Select(t *testing.T) {
	ConfirmSelect := func(r, m map[int]bool, k ...int) {
		x := Select(m, k...)
		if len(r) != len(x) {
			t.Errorf("Select(%v, %v) should contain %v elements not %v", m, k, len(r), len(x))
		}
		for _, v := range k {
			switch vo, ok := x[v]; {
			case !ok:
				t.Errorf("Select(%v, %v) should contain key %v", m, k, v)
			case vo != r[v]:
				t.Errorf("Select(%v, %v)[%v] should be %v not %v", m, k, v, r[v], vo)
			}
		}
	}

	ConfirmSelect(map[int]bool{}, map[int]bool{0: true, 1: false, 2: true})
	ConfirmSelect(map[int]bool{0: true}, map[int]bool{0: true, 1: false, 2: true}, 0)
	ConfirmSelect(map[int]bool{1: false}, map[int]bool{0: true, 1: false, 2: true}, 1)
	ConfirmSelect(map[int]bool{2: true}, map[int]bool{0: true, 1: false, 2: true}, 2)

	ConfirmSelect(map[int]bool{0: true}, map[int]bool{0: true, 1: false, 2: true}, 0, 0)
	ConfirmSelect(map[int]bool{1: false}, map[int]bool{0: true, 1: false, 2: true}, 1, 1)
	ConfirmSelect(map[int]bool{2: true}, map[int]bool{0: true, 1: false, 2: true}, 2, 2)

	ConfirmSelect(map[int]bool{0: true, 1: false}, map[int]bool{0: true, 1: false, 2: true}, 0, 1)
	ConfirmSelect(map[int]bool{1: false, 2: true}, map[int]bool{0: true, 1: false, 2: true}, 1, 2)
	ConfirmSelect(map[int]bool{0: true, 2: true}, map[int]bool{0: true, 1: false, 2: true}, 0, 2)

	ConfirmSelect(map[int]bool{1: false}, map[int]bool{0: true, 1: false, 2: true}, 1)

	ConfirmSelect(map[int]bool{2: true}, map[int]bool{0: true, 1: false, 2: true}, 2)

}
