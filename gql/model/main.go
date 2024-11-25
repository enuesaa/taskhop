// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Health struct {
	Time string `json:"time"`
	Ok   bool   `json:"ok"`
	Code string `json:"code"`
}

type LogInput struct {
	Output string `json:"output"`
}

type Mutation struct {
}

type Query struct {
}

type RunCmdInput struct {
	Memo string `json:"memo"`
	Cmd  string `json:"cmd"`
}

type RunCmdOutput struct {
	Output string `json:"output"`
}

type Subscription struct {
}

type Task struct {
	Cmds []string `json:"cmds"`
}
