package schedule

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"time"
)

const (
	// DefaultExitCode default err code if not get run task2 code
	DefaultExitCode int = -1
)

// Please Implment io.ReadCloser
// reader last 3 byte must be exit code
type Runner interface {
	Run(ctx context.Context) (out io.ReadCloser)
	Type() string
}

var _ Runner = DataCode{}

type Lang string

// DataCode run code
type DataCode struct {
	Lang       Lang   `json:"kind"`
	LangDesc   string `json:"langDesc" comment:"Lang"`
	ScriptCode string `json:"script_code" comment:"ScriptCode"`
}

func getCmd(ctx context.Context, Lang string, code string) (*exec.Cmd, string, error) {
	switch Lang {
	case "shell":
		return runShell(ctx, code)
	case "python":
		return runPython(ctx, code)
	case "python3":
		return runPython3(ctx, code)
	case "golang":
		return runGolang(ctx, code)
	case "nodejs":
		return runNodejsV1(ctx, code)
	case "windowsbat":
		return runWindowsBat(ctx, code)
	default:
		return nil, "", fmt.Errorf("can not support lang: %s ", Lang)
	}
}

// Shell
// run shell code
func runShell(ctx context.Context, scriptCode string) (*exec.Cmd, string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}
	tmpFile, err := ioutil.TempFile("", "*.sh")
	if err != nil {
		return nil, "", err
	}
	shellCodePath := tmpFile.Name()
	_, err = tmpFile.WriteString(scriptCode)
	if err != nil {
		return nil, "", err
	}

	_ = tmpFile.Sync()
	tmpFile.Close()
	cmd := exec.CommandContext(ctx, shell, shellCodePath)
	return cmd, shellCodePath, nil
}

// Python
// run python code
func runPython(ctx context.Context, scriptCode string) (*exec.Cmd, string, error) {
	tmpFile, err := ioutil.TempFile("", "*.py")
	if err != nil {
		return nil, "", err
	}
	pythonCodePath := tmpFile.Name()
	_, err = tmpFile.WriteString(scriptCode)
	if err != nil {
		return nil, "", err
	}
	_ = tmpFile.Sync()
	tmpFile.Close()
	cmd := exec.CommandContext(ctx, "python", pythonCodePath)
	return cmd, pythonCodePath, nil
}

// Python3
// run python code
func runPython3(ctx context.Context, scriptCode string) (*exec.Cmd, string, error) {
	tmpFile, err := ioutil.TempFile("", "*.py")
	if err != nil {
		return nil, "", err
	}
	python3CodePath := tmpFile.Name()
	_, err = tmpFile.WriteString(scriptCode)
	if err != nil {
		return nil, "", err
	}
	tmpFile.Sync()
	tmpFile.Close()
	cmd := exec.CommandContext(ctx, "python3", python3CodePath)
	return cmd, python3CodePath, nil
}

// Golang
const (
	modContent = `module jobor

go `
	modName   = "go.mod"
	goNamePre = "crocodile_"
)

// run golang code
func runGolang(ctx context.Context, scriptCode string) (*exec.Cmd, string, error) {
	// golang version must rather equal 1.11
	// GO111MODULE ust be on
	cmd := exec.CommandContext(context.Background(), "go", "version")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, "", err
	}

	pattern := `[0-1]\.[0-9]{1,2}`
	re := regexp.MustCompile(pattern)
	goVersion := re.FindString(string(out))
	if goVersion < "1.11" {
		err := fmt.Errorf("go version must rather equal go1.11 and enable go module")
		return nil, "", err
	}
	if os.Getenv("GO111MODULE") != "on" {
		_ = os.Setenv("GO111MODULE", "on")
	}
	modContent := modContent + goVersion + "\n"

	tmpdir, err := ioutil.TempDir("", "crocodile_")
	if err != nil {
		return nil, "", err
	}

	err = ioutil.WriteFile(path.Join(tmpdir, modName), []byte(modContent), os.ModePerm)
	if err != nil {
		return nil, "", err
	}
	goNameFile := goNamePre + strconv.FormatInt(time.Now().Unix(), 10) + ".go"
	err = ioutil.WriteFile(path.Join(tmpdir, goNameFile), []byte(scriptCode), os.ModePerm)
	if err != nil {
		return nil, "", err
	}

	_ = os.Chdir(tmpdir)

	goCmd := exec.CommandContext(ctx, "go", "run", goNameFile)

	return goCmd, tmpdir, nil
}

// Windows bat
// run bat code
func runWindowsBat(ctx context.Context, scriptCode string) (*exec.Cmd, string, error) {
	tmpFile, err := ioutil.TempFile("", "*.bat")
	if err != nil {
		return nil, "", err
	}
	batCodePath := tmpFile.Name()
	_, err = tmpFile.WriteString(scriptCode)
	if err != nil {
		return nil, "", err
	}

	tmpFile.Sync()
	tmpFile.Close()
	cmd := exec.CommandContext(ctx, "cmd", "/C", batCodePath)
	return cmd, batCodePath, nil
}

func (d DataCode) Run(ctx context.Context) (out io.ReadCloser) {
	pr, pw := io.Pipe()
	go func() {
		var (
			exitCode = DefaultExitCode
			err      error
			codePath string
			cmd      *exec.Cmd
		)
		defer pw.Close()
		defer func() {
			now := time.Now().Local().Format("2006-01-02 15:04:05: ")
			pw.Write([]byte(fmt.Sprintf("%sTask Run Finished, Return ScriptCode:%5d", now, exitCode))) // write exitCode,total 5 byte
			if codePath != "" {
				_ = os.Remove(codePath)
			}
		}()
		cmd, codePath, err = getCmd(ctx, string(d.Lang), d.ScriptCode)
		if err != nil {
			pw.Write([]byte(err.Error()))
			return
		}
		cmd.Stdout = pw
		cmd.Stderr = pw
		err = cmd.Start()
		if err != nil {
			pw.Write([]byte(err.Error()))
			return
		}

		err = cmd.Wait()
		if err != nil {
			// deal err
			// if context err,will change err to custom msg
			switch ctx.Err() {
			case context.DeadlineExceeded:
				//pw.Write([]byte(resp.GetMsg(resp.ErrCtxDeadlineExceeded)))
			case context.Canceled:
				//pw.Write([]byte(resp.GetMsg(resp.ErrCtxCanceled)))
			default:
				pw.Write([]byte(err.Error()))
			}

			// try to get the exit code
			if exitError, ok := err.(*exec.ExitError); ok {
				exitCode = exitError.ExitCode()
			}
		} else {
			exitCode = 0
		}

	}()
	return pr
}

func (d DataCode) Type() string {
	panic("implement me")
}

// String return Lanf str
func (l Lang) String() string {
	switch string(l) {
	case "shell":
		return "shell"
	case "python":
		return "python"
	case "python3":
		return "python3"
	case "golang":
		return "golang"
	case "nodejs":
		return "nodejs"
	case "windowsbat":
		return "windowsbat"
	default:
		return "unknow lang"
	}
}

