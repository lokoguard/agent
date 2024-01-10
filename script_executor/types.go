package script_executor

import "fmt"

type ScriptDefinition struct {
	TaskID string   `json:"task_id"`
	Script string   `json:"script"`
	Args   []string `json:"args"`
}

type ScriptResult struct {
	TaskID   string `json:"task_id"`
	Output   string `json:"output"`
	Error    string `json:"error"`
	Success  bool   `json:"success"`
	ExitCode int    `json:"exit_code"`
}

func (s ScriptResult) String() string {
	return fmt.Sprintf("TaskID: %s\nOutput: %s\nError: %s\nSuccess: %t\nExitCode: %d\n", s.TaskID, s.Output, s.Error, s.Success, s.ExitCode)
}

type ResultCallbackType func (*ScriptResult, error)