package _223

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//感觉是一道体力题 主要是对坐标系要熟悉
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	s1 := (C - A) * (D - B)
	s2 := (G - E) * (H - F)

	if E >= C || G <= A || F >= D || H <= B {
		return s1 + s2
	}

	var cover int
	if E <= A && G <= C {
		//左边穿过
		cover = (G - A) * (min(D, H) - max(B, F))
	} else if E >= A && G >= C {
		//右边穿过
		cover = (C - E) * (min(D, H) - max(B, F))
	} else if E <= A && G >= C {
		//左右都穿过
		cover = (C - A) * (min(D, H) - max(B, F))
	} else {
		//左右都没穿过
		cover = (G - E) * (min(D, H) - max(B, F))
	}
	//fmt.Println(cover)
	return s1 + s2 - cover
}
