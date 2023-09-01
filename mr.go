package mr

func Map[T any, R any](ts []T, f func(T) R) []R {
	nt := make([]R, 0, len(ts))
	for _, t := range ts {
		nt = append(nt, f(t))
	}
	return nt
}

func Reduce[T any](ts []T, f func(T, T) T, init T) T {
	t := init
	for _, tt := range ts {
		t = f(t, tt)
	}
	return t
}

func Filter[T any](ts []T, f func(T) bool) []T {
	nt := make([]T, 0, len(ts))
	for _, t := range ts {
		if f(t) {
			nt = append(nt, t)
		}
	}
	return nt
}

func ToMap[T any, KEY comparable](ts []T, f func(T) KEY) map[KEY]T {
	m := make(map[KEY]T, len(ts))
	for _, v := range ts {
		m[f(v)] = v
	}
	return m
}

func Unique[T any, KEY comparable](ts []T, f func(T) KEY) []T {
	m := make(map[KEY]struct{}, len(ts))
	nt := make([]T, 0, len(ts))
	for _, t := range ts {
		k := f(t)
		if _, ok := m[k]; ok {
			continue
		}
		nt = append(nt, t)
		m[k] = struct{}{}
	}
	return nt
}

func Diff[T any, KEY comparable](ts1, ts2 []T, f func(T) KEY) []T {
	nt := make([]T, 0, len(ts1))
	ts2Map := ToMap(ts2, f)
	for _, t := range ts1 {
		if _, ok := ts2Map[f(t)]; !ok {
			nt = append(nt, t)
		}
	}
	return nt
}

func Intersect[T any, KEY comparable](ts1, ts2 []T, f func(T) KEY) []T {
	nt := make([]T, 0, len(ts1))
	ts2Map := ToMap(ts2, f)
	for _, t := range ts1 {
		if _, ok := ts2Map[f(t)]; ok {
			nt = append(nt, t)
		}
	}
	return nt
}

func Merge[T any, KEY comparable](ts1, ts2 []T, f func(T) KEY) []T {
	nt := make([]T, 0, len(ts1)+len(ts2))
	ts2Map := ToMap(ts2, f)
	replaced := make(map[KEY]struct{}, len(ts2))
	for _, t := range ts1 {
		if v, ok := ts2Map[f(t)]; ok {
			nt = append(nt, v)
			replaced[f(t)] = struct{}{}
		} else {
			nt = append(nt, t)
		}
	}

	for _, t := range ts2 {
		if _, ok := replaced[f(t)]; !ok {
			nt = append(nt, t)
		}
	}
	return nt
}

func Contains[T any, KEY comparable](ts []T, t T, f func(T) KEY) bool {
	tsMap := ToMap(ts, f)
	_, ok := tsMap[f(t)]
	return ok
}
