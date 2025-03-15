
# jptime

`jptime` is a Go package that provides utilities for working with Japan Standard Time (JST) and other Japan-specific time-related operations.

## Features

- `NowJST()`: Returns the current time in Japan Standard Time (JST).
- `ToJST(t time.Time)`: Converts a given time to Japan Standard Time (JST).
- `TimeFormatJST(t time.Time)`: Formats a time in the Japanese date format (yyyy年MM月dd日).
- `WeekdaysJPN()`: Returns a slice of the weekdays in Japanese (Sunday to Saturday).
- `DayOfWeekJPN(t time.Time)`: Returns the Japanese weekday name for the given date.

## Installation

To install the `jptime` package, run:

```bash
go get github.com/Yamaguchi-Katsuya/jptime.git
```

## Example Usage

```go
package main

import (
	"fmt"
	"time"

    "github.com/Yamaguchi-Katsuya/jptime"
)

func main() {
	// Get current time in JST
	now, err := jptime.NowJST()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current time in JST:", now)

	// Convert time to JST
	utc := time.Now().UTC()
	jstTime, err := jptime.ToJST(utc)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Converted UTC time to JST:", jstTime)

	// Format time in Japanese date format
	formattedTime := jptime.TimeFormatJST(now)
	fmt.Println("Formatted time:", formattedTime)

	// Get Japanese weekday names
	weekdays := jptime.WeekdaysJPN()
	fmt.Println("Weekdays in Japanese:", weekdays)

	// Get the weekday for a specific date
	dayOfWeek := jptime.DayOfWeekJPN(now)
	fmt.Println("Day of the week in Japanese:", dayOfWeek)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
