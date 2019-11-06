package _393

func validUtf8(data []int) bool {
    pos := 0
    goal := 1
    for i := 0; i < len(data); i++ {
        b := data[i]
        if pos == 0 {
            if b <= 0x7f {
                continue
            } else if b <= 0xdf {
                goal = 2
                pos++
                continue
            } else if b <= 0xef {
                goal = 3
                pos++
                continue
            } else if b <= 0xf7 {
                goal = 4
                pos++
                continue
            } else {
                return false
            }
        } else {
            if b >= 0x80 && b <= 0xbf {
                pos++
                if pos == goal {
                    pos = 0
                    goal = 1
                }
                continue
            } else {
                return false
            }
        }
    }
    if pos != 0 {
        return false
    }
    return true
}
