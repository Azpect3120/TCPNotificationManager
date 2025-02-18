package notify

import (
	"os/exec"
	"runtime"

	"github.com/Azpect3120/TCPNotificationManager/internal/client"
	"github.com/Azpect3120/TCPNotificationManager/internal/logger"
)

// Notify sends a notification to the desktop, this should
// only be used by the client, as the server has no GUI or
// reason to have a UI/UX.
//
// The functions called will differ by system, but for now
// only Linux is supported.
//
// TODO: Implement Windows support.
func Notify(client *client.TcpClient, title, message string) {
	switch runtime.GOOS {
	case "linux":
		_notifyLinux(client, title, message)
	default:
		client.Logger.Log("Unsupported OS for notifications", logger.ERROR)
	}
}

// Notify a Linux system using the notify-send command. The
// system should have notify-send installed.
func _notifyLinux(client *client.TcpClient, title, message string) {
	if err := exec.Command("notify-send", title, message).Run(); err != nil {
		client.Logger.Log("Error sending notification", logger.ERROR)
	}
}
