package script_executor

import (
	"encoding/json"
	"fmt"
)

type ScriptDefinition struct {
	TaskID int      `json:"id"`
	Script string   `json:"script"`
	Args   []string `json:"args"`
}

type ScriptResult struct {
	TaskID   int    `json:"id"`
	Output   string `json:"output"`
	Error    string `json:"error"`
	Success  bool   `json:"success"`
	ExitCode int    `json:"exit_code"`
}

func (s ScriptResult) String() string {
	return fmt.Sprintf("TaskID: %s\nOutput: %s\nError: %s\nSuccess: %t\nExitCode: %d\n", s.TaskID, s.Output, s.Error, s.Success, s.ExitCode)
}

func (s ScriptResult) JSON() ([]byte, error) {
	return json.Marshal(s)
}

type ResultCallbackType func(*ScriptResult, error)
