package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// Define some colors for print statements
var (
	Yellow = color.New(color.FgYellow).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
	Purple = color.New(color.FgMagenta).SprintFunc()
	Cyan   = color.New(color.FgCyan).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
	Blue   = color.New(color.FgBlue).SprintFunc()
)

func formatDuration(d time.Duration) string {
	days := d / (24 * time.Hour)
	hours := d % (24 * time.Hour) / time.Hour
	minutes := d % time.Hour / time.Minute
	seconds := d % time.Minute / time.Second

	return fmt.Sprintf("%d days %02d hours %02d minutes %02d seconds", days, hours, minutes, seconds)
}

// Create a wrapper function for fmt.Println
func Println(a ...interface{}) {
	fmt.Println(a...)
}

func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func Print(a ...interface{}) {
	fmt.Print(a...)
}
