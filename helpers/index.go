package helpers

import (
	"log"
)

// Log wrapper
func Info(template string, values ...interface{}) {
	log.Printf("[api] "+template+"\n", values...)
}

func AnwerMessage(code int64, message string) map[string]interface{} {
	return map[string]interface{}{"Code": code, "Message":message}
}