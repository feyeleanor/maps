//	Copyright 2022 Game Design Analytics
//	All Rights Reserved
//
//	Authors:
//		Eleanor McHugh (eleanor@games-with-brains.com)
//
package maps

func Keys[K comparable, V any](m map[K]V) (r []K) {
	for k, _ := range m {
		r = append(r, k)
	}
	return
}

func KeysMatch[K comparable, V any](m, o map[K]V) (r bool) {
	if r = len(m) == len(o); r {
		for k, _ := range m {
			if _, r = o[k]; !r {
				break
			}
		}
	}
	return
}

func Equal[K, V comparable](m, o map[K]V) (r bool) {
	if r = len(m) == len(o); r {
		var vo V
		for k, v := range m {
			if vo, r = o[k]; r {
				if r = vo == v; !r {
					break
				}
			} else {
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
