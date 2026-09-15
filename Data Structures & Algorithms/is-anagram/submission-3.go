func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    mapS := make(map[int]int)
    mapT := make(map[int]int)

    for i := 0; i < len(s); i++ {
        mapS[int(s[i])]++
        mapT[int(t[i])]++
    }

    for i := int('a'); i <= int('z'); i++ {
        if mapS[i] != mapT[i] {
            return false
        }
    }
    return true
}
