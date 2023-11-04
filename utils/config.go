package handlers

import (
	"fmt"
	"os"
)

func setTTLConfig(ttl int) error {
	// Check if the ~/.ioson directory exists
	iosonDir := os.Getenv("HOME") + "/.ioson"
	if _, err := os.Stat(iosonDir); os.IsNotExist(err) {
		// If the directory does not exist, create it
		err = os.Mkdir(iosonDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	// Open the config file for writing
	file, err := os.OpenFile(iosonDir+"/.ioson.ini", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	// Write the TTL value to the config file
	_, err = fmt.Fprintf(file, "ttl=%d", ttl)
	if err != nil {
		return fmt.Errorf("failed to write to config file: %v", err)
	}

	return nil
}
