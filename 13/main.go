package _13

func romanToInt(s string) int {
    m := map[string]int{"IV": 3, "IX": 8, "XL": 30, "XC": 80, "CD": 300, "CM": 800}
    mw := map[uint8]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
    result := 0
    result += mw[s[0]]
    for i := 1; i < len(s); i++ {
        sub := s[i-1 : i+1]
        if sum, ok := m[sub]; ok {
            result += sum
        } else {
            result += mw[s[i]]
        }
    }
    return result
}
