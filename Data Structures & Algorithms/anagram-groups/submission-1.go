func groupAnagrams(strs []string) [][]string {
    group := make(map[[26]int][]string)
    out := make([][]string, 0, len(strs))
    for _, str := range strs {
        count := [26]int{}
        for _, letter := range str {
            count[letter - 'a']++;
        }
        group[count] = append(group[count], str)
    }

    for val := range group {
        out = append(out, group[val])
    }

    return out
}
