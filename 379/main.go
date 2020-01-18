package _379

type PhoneDirectory struct {
    data       map[int]bool
    maxNumbers int
}

/** Initialize your data structure here
  @param maxNumbers - The maximum numbers that can be stored in the phone directory. */
func Constructor(maxNumbers int) PhoneDirectory {
    data := map[int]bool{}
    for i := 0; i < maxNumbers; i++ {
        data[i] = true
    }
    return PhoneDirectory{data: data, maxNumbers: maxNumbers}
}

/** Provide a number which is not assigned to anyone.
  @return - Return an available number. Return -1 if none is available. */
func (this *PhoneDirectory) Get() int {
    if len(this.data) == 0 {
        return -1
    }
    number := -1
    for k := range this.data {
        number = k
        break
    }
    delete(this.data, number)
    return number
}

/** Check if a number is available or not. */
func (this *PhoneDirectory) Check(number int) bool {
    _, ok := this.data[number]
    return ok
}

/** Recycle or release a number. */
func (this *PhoneDirectory) Release(number int) {
    this.data[number] = true
}
