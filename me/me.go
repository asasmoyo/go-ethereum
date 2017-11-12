package me

func End() {
	Log("Stopping ME")
	err := blockFile.Close()
	if err != nil {
		Log("Error closing file for block write", "err", err)
	}
	Log("ME stopped")
}
