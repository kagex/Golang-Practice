package listops

type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	result := initial
	for _, el := range s {
		result = fn(result, el)
	}
	return result
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	result := initial
	for i := len(s) - 1; i >= 0; i-- {
		result = fn(s[i], result)
	}
	return result
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := make(IntList, 0, len(s))
	for _, el := range s {
		if fn(el) {
			result = append(result, el)
		}
	}
	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, len(s))
	for i, el := range s {
		result[i] = fn(el)
	}
	return result
}

func (s IntList) Reverse() IntList {
	result := make(IntList, len(s))
	for i, el := range s {
		result[len(s)-1-i] = el
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	result := make(IntList, len(s)+len(lst))
	copy(result, s)
	copy(result[len(s):], lst)
	return result
}

func (s IntList) Concat(lists []IntList) IntList {
	totalLen := len(s)
	for _, lst := range lists {
		totalLen += len(lst)
	}
	
	result := make(IntList, 0, totalLen)
	result = append(result, s...)
	for _, lst := range lists {
		result = append(result, lst...)
	}
	return result
}