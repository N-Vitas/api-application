package helpers

import "log"

// Log wrapper
func Info(template string, values ...interface{}) {
	log.Printf("[api] "+template+"\n", values...)
}
