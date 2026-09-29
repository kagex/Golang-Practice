package sublist

// Relation type is defined in relations.go file.

func Sublist(a, b []int) Relation {
	switch {
	case equal(a, b):
		return RelationEqual
	case isSublist(a, b):
		return RelationSublist
	case isSublist(b, a):
		return RelationSuperlist
	default:
		return RelationUnequal
	}
}

// equal проверяет, равны ли два списка
func equal(a, b []int) bool {
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

// isSublist проверяет, является ли small подсписком large
func isSublist(small, large []int) bool {
	// Пустой список — подсписок любого списка
	if len(small) == 0 {
		return true
	}
	
	// Если small длиннее large, он не может быть подсписком
	if len(small) > len(large) {
		return false
	}
	
	// Ищем small в large как непрерывную последовательность
	for i := 0; i <= len(large)-len(small); i++ {
		if equal(small, large[i:i+len(small)]) {
			return true
		}
	}
	
	return false
}