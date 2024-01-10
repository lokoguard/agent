package script_executor

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func (s ScriptDefinition) Run() (*ScriptResult, error) {
	// Generate a random hash for the file name
	hash, err := generateRandomHash(16)
	if err != nil {
		return nil, fmt.Errorf("error generating random hash: %v", err)
	}
	filePath := fmt.Sprintf("/tmp/lokoguard_%s.sh", hash)

	// create the script file
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close() // Close the file when done

	// Write the script to the file
	_, err = io.WriteString(file, s.Script)
	if err != nil {
		return nil, fmt.Errorf("error writing script to file: %v", err)
	}

	// Make the file executable
	err = file.Chmod(0755)
	if err != nil {
		return nil, fmt.Errorf("error changing file permissions: %v", err)
	}

	// Prepare command execution
	cmd := exec.Command("bash", append([]string{filePath}, s.Args...)...)

	// Create separate stdout and stderr buffers
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the Bash script using bash command
	err = cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// The command had a non-zero exit code
			return nil, fmt.Errorf("script failed with exit code %d\nStdout:\n%s\nStderr:\n%s", exitErr.ExitCode(), stdout.String(), stderr.String())
		}
		return nil, fmt.Errorf("error running script: %v", err)
	}

	// Return the result
	return &ScriptResult{
		TaskID:   s.TaskID,
		Output:   stdout.String(),
		Error:    stderr.String(),
		Success:  cmd.ProcessState.Success(),
		ExitCode: cmd.ProcessState.ExitCode(),
	}, nil
}


func (s ScriptDefinition) RunWithCallback(callback ResultCallbackType){
	go func (s1 ScriptDefinition) {
		res, err := s1.Run()
		callback(res, err)
	}(s)
	
}

// For generate name for script file
func generateRandomHash(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}