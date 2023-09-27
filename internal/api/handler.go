package api

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/pelletier/go-toml"
)

func ExecuteCommandHandler(w http.ResponseWriter, r *http.Request) {
	config, err := toml.LoadFile("./configs/config.toml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	command := config.Get("command.execute").(string)

	var out []byte
	switch runtime.GOOS {
	case "windows":
		// For Windows, you might need to adjust the command and use the appropriate shell
		cmd := exec.Command("cmd", "/C", command)
		out, err = cmd.Output()
	case "linux", "darwin":
		// For Linux and macOS, you might not need to adjust the command, using sh should be fine
		cmd := exec.Command("sh", "-c", command)
		out, err = cmd.Output()
	default:
		http.Error(w, "Unsupported Operating System", http.StatusInternalServerError)
		return
	}

	if err != nil {
		errorMessage := fmt.Sprintf("Error executing command: %s\n%s", command, err.Error())
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}

	// Replace newlines with <br> for better readability in HTML responses
	//output := strings.ReplaceAll(string(out), "\n", "<br>")
	//fmt.Fprint(w, output)

	fmt.Fprint(w, string(out))
}
