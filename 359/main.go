package _359

type Logger struct {
    printed map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
    return Logger{map[string]int{}}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    lastPrinted, ok := this.printed[message]
    result := false
    if !ok {
        result = true
        this.printed[message] = timestamp
    } else {
        if timestamp-lastPrinted >= 10 {
            result = true
            this.printed[message] = timestamp
        }
    }
    return result
}
