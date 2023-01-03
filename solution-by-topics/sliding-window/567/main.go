package main

func checkInclusion(s1 string, s2 string) bool {
	ls1, ls2 := len(s1), len(s2)
	if (ls1 > ls2) {
		return false;
	}
	s1map := make([]int, 26)
	s2map := make([]int, 26);
	for i := 0; i < ls1; i++ {
		s1map[s1[i] - 'a']++;
		s2map[s2[i] - 'a']++;
	}

	count := 0;
	for i := 0; i < 26; i++ {
		if (s1map[i] == s2map[i]){ count++ }
	}

	for i := 0; i < ls2 - ls1; i++ {
		r,l := s2[(i + ls1)] - 'a', s2[i] - 'a';
		if (count == 26){ return true }
		s2map[r]++;
		if (s2map[r] == s1map[r]) {
			count++;
		} else if (s2map[r] == s1map[r] + 1) {
			count--;
		}
		s2map[l]--;
		if (s2map[l] == s1map[l]) {
			count++;
		} else if (s2map[l] == s1map[l] - 1) {
			count--;
		}
	}
	return count == 26;
}

func main(){
	println(checkInclusion("ab", "eidbaooo"))
	println(checkInclusion("ab", "eidboaoo"))
}
