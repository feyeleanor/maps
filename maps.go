//	Copyright 2022 Game Design Analytics
//	All Rights Reserved
//
//	Authors:
//		Eleanor McHugh (eleanor@games-with-brains.com)
//
package maps

func Len[K comparable, V any](m map[K]V) int {
	return len(m)
}

func Get[K comparable, V any](m map[K]V, k K) (V, bool) {
	v, ok := m[k]
	return v, ok
}

func Set[K comparable, V any](m map[K]V, k K, v V) {
	m[k] = v
}

func Keys[K comparable, V any](m map[K]V) (r []K) {
	for k, _ := range m {
		r = append(r, k)
	}
	return
}

func KeysMatch[K comparable, V any](m map[K]V, o map[K]V) (r bool) {
	if r = len(m) == len(o); r {
		for k, _ := range m {
			if _, r = o[k]; !r {
				break
			}
		}
	}
	return
}

func Copy[K comparable, V any](m map[K]V) (r map[K]V) {
	r = make(map[K]V)
	for k, v := range m {
		r[k] = v
	}
	return
}

func Merge[K comparable, V any](m map[K]V, o ...map[K]V) (r map[K]V) {
	r = Copy(m)
	for _, on := range o {
		for k, v := range on {
			r[k] = v
		}
	}
	return
}

func Select[K comparable, V any](m map[K]V, k ...K) (r map[K]V) {
	r = make(map[K]V)
	for _, k := range k {
		if v, ok := m[k]; ok {
			r[k] = v
		}
	}
	return
}
