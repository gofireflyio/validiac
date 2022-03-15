package api

import (
    "fmt"
    "os"
    "os/exec"
)

var TFLintExec = ""
var TfLintConfig = "./bin/.tflint.hcl"

func TFLint(in []byte) ([]byte, error) {
	path, err := asTempFile("", ".tf", in)
	if err != nil {
		return nil, err
	}

	defer os.Remove(path) // nolint: errcheck
	return exec.Command(TFLintExec, fmt.Sprintf("--config=%s", TfLintConfig), "--enable-plugin=aws", "--enable-plugin=azurerm", "--enable-plugin=google", path, "--no-color").CombinedOutput()
}
