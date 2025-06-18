package types

type Command string

const (
	CmdTask Command = "task"
	CmdSession Command = "session"
)

type TaskSubCommand string

const (
	TaskAdd TaskSubCommand = "add"
	TaskDelete TaskSubCommand = "delete"
	TaskEdit TaskSubCommand = "edit"
	TaskList TaskSubCommand = "list"
)

type SessionSubCommand string

const (
	SessionStart SessionSubCommand = "start"
	SessionPause SessionSubCommand = "pause"
	SessionStop SessionSubCommand = "stop"
)
