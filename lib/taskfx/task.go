package taskfx

type Task struct {
	Title string   `yaml:"title" validate:"required"`
	Cmds  []string `yaml:"cmds" validate:"required,min=1,dive,required"`
}
