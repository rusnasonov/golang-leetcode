package main

func hash(str string) [26]uint8 {
	h := [26]uint8{}
	for i := 0; i < len(str); i++ {
		h[str[i]-97]++
	}
	return h
}

func groupAnagrams(strs []string) [][]string {
	var out [][]string
	var res map[[26]uint8][]string = make(map[[26]uint8][]string)
	for _, v := range strs {
		h := hash(v)
		values := res[h]
		res[h] = append(values, v)
	}
	for _, v := range res {
		out = append(out, v)
	}
	return out
}

