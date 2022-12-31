package set

type Set map[rune]bool

func New() *Set {
	s := make(Set)
	return &s
}

func (s Set) Add(i rune) {
	s[i] = true
}

func (s Set) Remove(i rune) {
	delete(s, i)
}

func (s Set) ToSlice() []rune {
	result := []rune{}
	for k := range s {
		result = append(result, k)
	}
	return result
}

func Difference(a, b *Set) []rune {
	results := []rune{}
	for k := range *a {
		if _, found := (*b)[k]; !found {
			results = append(results, k)
		}
	}
	return results
}

func Union(a, b *Set) []rune {
	results := []rune{}
	for k := range *a {
		if _, found := (*b)[k]; found {
			results = append(results, k)
		}
	}
	return results
}

func FromSlice(s []rune) *Set {
	result := New()
	for _, v := range s {
		result.Add(v)
	}
	return result
}
