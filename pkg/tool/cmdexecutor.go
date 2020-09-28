package tool

import (
	"strings"
	"text/template"

	"github.com/helm/chart-testing/v3/pkg/exec"
)

type CmdTemplateExecutor struct {
	exec exec.ProcessExecutor
}

func NewTester(exec exec.ProcessExecutor) CmdTemplateExecutor {
	return CmdTemplateExecutor{
		exec: exec,
	}
}

func (t CmdTemplateExecutor) RunCommand(cmdTemplate string, data interface{}) error {
	var template = template.Must(template.New("command").Parse(cmdTemplate))
	var b strings.Builder
	err := template.Execute(&b, data)
	renderedCommand := b.String()
	split := strings.Split(renderedCommand, " ")
	if err != nil {
		return err
	}
	return t.exec.RunProcess(split[0], split[1:])
}
