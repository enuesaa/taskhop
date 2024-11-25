package cmdsfile

type CmdsFile struct {
	Title string `yaml:"title"`
	Workdir string `yaml:"workdir"`
	Cmds []string `yaml:"cmds"`
}
