package _443

import "strconv"

func compress(chars []byte) int {
    current := chars[0]
    currentLength := 1
    nextPos := 0
    for i := 1; i < len(chars); i++ {
        if chars[i] == current {
            currentLength++
        } else {
            chars[nextPos] = current
            nextPos++
            if currentLength > 1 {
                lengthStr := strconv.Itoa(currentLength)
                for _, c := range lengthStr {
                    chars[nextPos] = byte(c)
                    nextPos++
                }
            }
            current = chars[i]
            currentLength = 1
        }
    }
    chars[nextPos] = current
    nextPos++
    if currentLength > 1 {
        lengthStr := strconv.Itoa(currentLength)
        for _, c := range lengthStr {
            chars[nextPos] = byte(c)
            nextPos++
        }
    }
    return nextPos
}
