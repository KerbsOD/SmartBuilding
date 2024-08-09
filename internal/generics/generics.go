package generics

func RepeatedElements[T comparable](slice []T) []T {
	countMap := make(map[T]int)
	for _, element := range slice {
		countMap[element]++
	}

	result := []T{}
	for element, count := range countMap {
		if count > 1 {
			result = append(result, element)
		}
	}

	return result
}

func MaxMapped[T any](arr []T, mapFunction func(T) int) int {
	mappedElements := []int{}
	for _, element := range arr {
		mappedElements = append(mappedElements, mapFunction(element))
	}

	currentMax := mappedElements[0]
	for _, element := range mappedElements {
		if element > currentMax {
			currentMax = element
		}
	}
	return currentMax
}

func MinimizeElementByComparer[T any](arr []T, compare func(a, b T) bool) T {
	currentMin := arr[0]
	for _, elem := range arr {
		if compare(elem, currentMin) {
			currentMin = elem
		}
	}
	return currentMin
}
