package _1103

func distributeCandies(candies int, num_people int) []int {
    result := make([]int, num_people)
    current := 0
    currentCandy := 1
    for candies != 0 {
        if candies > currentCandy {
            candies -= currentCandy
            result[current] += currentCandy
        } else {
            result[current] += candies
            candies = 0
        }
        current += 1
        current %= num_people
        currentCandy += 1
    }
    return result
}
