package utility

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	// Menentukan perintah untuk membersihkan layar sesuai dengan sistem operasi yang digunakan
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
