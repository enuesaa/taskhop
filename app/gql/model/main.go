// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Health struct {
	Time string `json:"time"`
	Ok   bool   `json:"ok"`
	Code string `json:"code"`
}

type LogInput struct {
	Type   LogType `json:"type"`
	Output string  `json:"output"`
}

type Mutation struct {
}

type Query struct {
}

type Task struct {
	Status TaskStatus `json:"status"`
	Cmds   []string   `json:"cmds"`
}

type LogType string

const (
	LogTypeCommandOutput LogType = "COMMAND_OUTPUT"
	LogTypeCommand       LogType = "COMMAND"
)

var AllLogType = []LogType{
	LogTypeCommandOutput,
	LogTypeCommand,
}

func (e LogType) IsValid() bool {
	switch e {
	case LogTypeCommandOutput, LogTypeCommand:
		return true
	}
	return false
}

func (e LogType) String() string {
	return string(e)
}

func (e *LogType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LogType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LogType", str)
	}
	return nil
}

func (e LogType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskStatus string

const (
	TaskStatusRegistration   TaskStatus = "REGISTRATION"
	TaskStatusDownloadAssets TaskStatus = "DOWNLOAD_ASSETS"
	TaskStatusPrompt         TaskStatus = "PROMPT"
	TaskStatusProceeding     TaskStatus = "PROCEEDING"
	TaskStatusCompleted      TaskStatus = "COMPLETED"
)

var AllTaskStatus = []TaskStatus{
	TaskStatusRegistration,
	TaskStatusDownloadAssets,
	TaskStatusPrompt,
	TaskStatusProceeding,
	TaskStatusCompleted,
}

func (e TaskStatus) IsValid() bool {
	switch e {
	case TaskStatusRegistration, TaskStatusDownloadAssets, TaskStatusPrompt, TaskStatusProceeding, TaskStatusCompleted:
		return true
	}
	return false
}

func (e TaskStatus) String() string {
	return string(e)
}

func (e *TaskStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskStatus", str)
	}
	return nil
}

func (e TaskStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
