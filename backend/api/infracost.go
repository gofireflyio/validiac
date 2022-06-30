package api

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/thoas/go-funk"
)

var InfraCostExec = getEnv("INFRACOST_EXEC", fmt.Sprintf("%s/infracost", BIN_PATH))

func InfraCost(in []byte) ([]byte, error) {
	var infraCostApiKey = funk.GetOrElse(os.Getenv("INFRACOST_API_KEY"), "infracost-api-key")
	path, err := asTempDir(".tf", in)
	if err != nil {
		return nil, err
	}

	defer os.Remove(path) // nolint: errcheck

	cmd := exec.Command(
		InfraCostExec,
		"breakdown",
		"--path",
		path,
		"--no-color",
		"--log-level=error",
	)
	cmd.Env = append(
		cmd.Env,
		fmt.Sprintf("INFRACOST_API_KEY=%s", infraCostApiKey),
		fmt.Sprintf("PATH=%s", os.Getenv("PATH")),
	)
	return cmd.CombinedOutput()
}
