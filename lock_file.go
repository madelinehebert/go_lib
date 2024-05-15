package go_lib

import (
	"os"
	"syscall"
	"time"

	"github.com/madelinehebert/go_lib"
)

func LockFile(file *os.File) error {
	//Set a lock
	for {
		//Attempt to set a lock on the file; if this fails, wait for a random amount of
		//time between 0 and one seconds before retrying
		if err := syscall.Flock((int(file.Fd())), syscall.LOCK_EX); err != nil {
			time.Sleep(time.Duration(go_lib.GenerateRandomTime(0, 1)))
		} else {
			break
		}
	}

	//If all is well, return nil, signalling the file was locked
	return nil
}
