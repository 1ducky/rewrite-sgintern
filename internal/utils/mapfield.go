package utils

func MapField[T, R any](items []T, fn func(T) (R, bool)) []R {
	out := make([]R, 0, len(items))
	for _, item := range items {
		if v, ok := fn(item); ok {
			out = append(out, v)
		}
	}
	return out
}
