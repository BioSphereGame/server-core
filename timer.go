package timer

import (
	"fmt"
	"os"
	"time"
)

type Timer struct {
	startTime time.Time
	endTime   time.Time
}

func NewTimer() *Timer {
	return &Timer{}
}

func (t *Timer) SetStart() {
	t.startTime = time.Now()
}

func (t *Timer) SetEnd(operationName string) {
	t.endTime = time.Now()
	duration := t.endTime.Sub(t.startTime)
	microseconds := duration.Microseconds()
	globalTime := time.Now().Format(time.RFC3339)

    data := fmt.Sprintf("%s - %s: %dus\n",
		globalTime, operationName, microseconds)

	file, err := os.Create("timer.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
