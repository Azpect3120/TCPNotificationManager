package notify

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Notify sends a notification to the desktop, this should
// only be used by the client, as the server has no GUI or
// reason to have a UI/UX.
//
// The functions called will differ by system, but for now
// only Linux is supported.
//
// TODO: Implement Windows support.
func Notify(title, message string) error {
	switch runtime.GOOS {
	case "linux":
		return _notifyLinux(title, message)
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
}

// Notify a Linux system using the notify-send command. The
// system should have notify-send installed.
func _notifyLinux(title, message string) error {
	if err := exec.Command("notify-send", title, message).Run(); err != nil {
		return fmt.Errorf("error sending notification: %v", err)
	}
	return nil
}
