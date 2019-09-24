package collections

func Unique(collection []string) []string {
	seen := make(map[string]struct{}, len(collection))
	offset := 0
	for _, elem := range collection {
		if _, ok := seen[elem]; ok {
			continue
		}
		seen[elem] = struct{}{}
		collection[offset] = elem
		offset++
	}

	return collection[:offset]
}

// difference returns the elements in 'a' that aren't in 'b'
func Difference(collectionA, collectionB []string) []string {
	mb := make(map[string]struct{}, len(collectionB))
	for _, x := range collectionB {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range collectionA {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
