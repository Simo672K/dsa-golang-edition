package leet

type HashSet struct {
	set    map[int]struct{}
	length int
}

func NewHashSet() *HashSet {
	return &HashSet{
		set:    make(map[int]struct{}),
		length: 0,
	}
}

func (hs *HashSet) Add(n int) {
	if _, exists := hs.set[n]; !exists {
		hs.set[n] = struct{}{}
		hs.length++
	}
}

func ContainsDuplicate(nums []int) bool {
	hashSet := NewHashSet()
	numsLen := len(nums)

	for _, num := range nums {
		hashSet.Add(num)
	}

	return !(numsLen == hashSet.length)
}
