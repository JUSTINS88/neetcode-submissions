func groupAnagrams(strs []string) [][]string {
    group := make(map[string][]string)
    out := make([][]string, 0, len(strs))
    for _, str := range strs {
        byted := []byte(str)
        sort.Slice(byted, func (i, j int) bool {return byted[i] < byted[j]})
        group[string(byted)] = append(group[string(byted)], str)
    }

    for val := range group {
        out = append(out, group[val])
    }

    return out
}
