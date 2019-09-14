package api

import (
	"bufio"
	"fmt"
	"github.com/xoreo/flash-encrypt/crypto"
	"github.com/xoreo/flash-encrypt/fs"
	"os"
	"strconv"
	"strings"
)

// Encrypt is an abstracted method to encrypt a directory.
func Encrypt(targetDriveID string) error {
	reader := bufio.NewReader(os.Stdin)

	// Get the connected drives
	drives, err := fs.GetDrivesDarwin()
	if err != nil {
		return err
	}

	// Find the drive name
	for i, drive := range drives {
		if strconv.Itoa(i) == targetDriveID {
			// Confirm
			fmt.Print("Are you sure you want to encrypt " + drive + " (yes/no)? ")
			confirmation, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			confirmation = strings.TrimSuffix(confirmation, "\n")

			if confirmation == "yes" {
				// Ask for passphrase
				fmt.Print("Passphrase: ")
				passphrase, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				passphrase = strings.TrimSuffix(passphrase, "\n")

				// Encrypt the entire flash drive
				err = crypto.EncryptDir(fs.GetDrivePath(drive), passphrase)
				if err != nil {
					return err
				}

				fmt.Printf("Encrypted %s\n", drive)

			} else {
				return nil
			}

		}
	}
	return nil
}

// Decrypt is an abstracted method to decrypt a directory.
func Decrypt(targetDriveID string) error {
	reader := bufio.NewReader(os.Stdin)

	// Get the connected drives
	drives, err := fs.GetDrivesDarwin()
	if err != nil {
		return err
	}

	// Get the drive name
	for i, drive := range drives {
		if strconv.Itoa(i) == targetDriveID {
			// Ask for confirmation
			fmt.Print("Are you sure you want to decrypt " + drive + " (yes/no)? ")
			confirmation, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			confirmation = strings.TrimSuffix(confirmation, "\n")

			if confirmation == "yes" {
				// Ask for passphrase
				fmt.Print("Passphrase: ")
				passphrase, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				passphrase = strings.TrimSuffix(passphrase, "\n")

				// Decrypt the entire directory
				err = crypto.DecryptDir(fs.GetDrivePath(drive), passphrase)
				if err != nil {
					return err
				}

				fmt.Printf("Decrypted %s\n", drive)

			} else {
				return nil
			}
		}
	}
	return nil
}

func ListDrives() error {
	// Get connected drives
	drives, err := fs.GetDrivesDarwin()
	if err != nil {
		panic(err)
	}

	// Print these drives
	fmt.Println("Connected drives:")
	for i, drive := range drives {
		fmt.Printf("[%d] %s\n", i, drive)
	}

	return nil
}

func Status() error {
	fmt.Println("status")
	return nil
}