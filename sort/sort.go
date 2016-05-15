package sort

// please refer stdlib's `sort` package, this is for my own education purposes

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
