package _57

import "testing"

func TestInsert(T *testing.T) {
    intervalGrid := [][]Interval{
        {
            {1, 3}, {6, 9},
        },
        {
            {1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
        },
        {},
        {
            {1, 5},
        },
        {
            {1, 5},
        },
        {
            {1, 5},
        },
    }
    newIntervalGrid := []Interval{
        {2, 5},
        {4, 8},
        {5, 7},
        {2, 3},
        {6, 8},
        {0, 3},
    }
    expects := [][]Interval{
        {
            {1, 5}, {6, 9},
        },
        {
            {1, 2}, {3, 10}, {12, 16},
        },
        {
            {5, 7},
        },
        {
            {1, 5},
        },
        {
            {1, 5}, {6, 8},
        },
        {
            {0, 5},
        },
    }
    for i, interval := range intervalGrid {
        newInterval := newIntervalGrid[i]
        actual := insert(interval, newInterval)
        expect := expects[i]
        if len(actual) != len(expect) {
            T.Errorf("expected: %v, actual: %v", expect, actual)
            continue
        }
        for i := range expect {
            if actual[i] != expect[i] {
                T.Errorf("expected: %v, actual: %v", expect, actual)
                break
            }
        }
    }
}
