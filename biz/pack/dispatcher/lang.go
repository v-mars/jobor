package dispatcher

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"io/ioutil"
	"jobor/biz/response"
	task2 "jobor/kitex_gen/task"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	// DefaultExitCode default err code if not get run task2 code
	DefaultExitCode int = -1
	LangShell           = "shell"
	LangPython          = "python"
	LangPython3         = "python3"
	LangGolang          = "golang"
	LangWindowsBat      = "windowsbat"
	LangApi             = "api"
)

// Runner Please Implement io.ReadCloser
// reader last 3 byte must be exit code
type Runner interface {
	Run(ctx context.Context) (out io.ReadCloser)
	Type() string
}

// GetDataRun get Task Runner
// get api or other
func GetDataRun(t *task2.TaskRequest) (Runner, error) {
	var data task2.TaskData
	err := json.Unmarshal(t.TaskData, &data)
	if err != nil {
		return nil, err
	}
	switch t.TaskLang {
	case LangShell, LangPython, LangPython3, LangGolang, LangWindowsBat:
		var code = DataCode{Lang: Lang(t.TaskLang), ScriptCode: *data.Content}
		code.LangDesc = code.Lang.String()
		return code, err
	case LangApi:
		var api = DataAPI{URL: data.Api.Url, Method: data.Api.Method, PayLoad: data.Api.Body}
		if api.Header == nil {
			api.Header = make(map[string]string)
		}
		//if len(data.Api.ContentType) > 0 {
		//	api.Header["Content-Type"] = data.Api.ContentType
		//}
		for _, mapVal := range data.Api.HeaderList {
			api.Header[mapVal.Key] = mapVal.Value
		}
		if data.Api.ContentType == "application/x-www-form-urlencoded" {
			api.PayLoad = ""
			for _, mapForm := range data.Api.WwwFormList {
				if len(api.PayLoad) > 0 {
					api.PayLoad = fmt.Sprintf("%s&%s=%s", api.PayLoad, mapForm.Key, mapForm.Value)
				} else {
					api.PayLoad = fmt.Sprintf("%s=%s", mapForm.Key, mapForm.Value)
				}
			}
		} else if data.Api.ContentType == "application/form-data" {
			api.PayLoad = ""
			for _, mapForm := range data.Api.FormDataList {
				if len(api.PayLoad) > 0 {
					api.PayLoad = fmt.Sprintf("%s&%s=%s", api.PayLoad, mapForm.Key, mapForm.Value)
				} else {
					api.PayLoad = fmt.Sprintf("%s=%s", mapForm.Key, mapForm.Value)
				}
			}
		}
		return api, nil
	default:
		err := fmt.Errorf("unsupport Task lang %s", t.TaskLang)
		return nil, err
	}
}

var _ Runner = DataCode{}

type DataAPI struct {
	Method  string
	URL     string
	PayLoad string
	Header  map[string]string
}

func (d DataAPI) Run(ctx context.Context) (out io.ReadCloser) {
	pr, pw := io.Pipe()
	go func() {
		var exitCode = DefaultExitCode
		defer pw.Close()
		defer func() {
			now := time.Now().Local().Format("2006-01-02 15:04:05: ")
			_, _ = pw.Write([]byte(fmt.Sprintf("\n%sRun Finished,Return Code:%5d", now, exitCode))) // write exitCode,total 5 byte
			// _,_=pw.Write([]byte(fmt.Sprintf("%3d", exitCode))) // write exitCode,total 3 byte
		}()
		// go1.13 use NewRequestWithContext
		hlog.CtxDebugf(ctx, "method: %s, url: %s, payLoad:", d.Method, d.URL, d.PayLoad)
		req, err := http.NewRequestWithContext(ctx, strings.ToUpper(d.Method), d.URL, bytes.NewReader([]byte(d.PayLoad)))
		if err != nil {
			_, _ = pw.Write([]byte(err.Error()))
			hlog.Errorf("http NewRequest failed: %s", err)
			return
		}

		for k, v := range d.Header {
			req.Header.Add(k, v)
		}

		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			hlog.Errorf("client Do failed: %s", err)
			var customErr bytes.Buffer
			switch ctx.Err() {
			case context.DeadlineExceeded:
				customErr.WriteString(response.GetMsg(response.ErrCtxDeadlineExceeded))
			case context.Canceled:
				customErr.WriteString(response.GetMsg(response.ErrCtxCanceled))
			default:
				customErr.WriteString(err.Error())
			}
			_, _ = pw.Write(customErr.Bytes())
			return
		}
		defer resp.Body.Close()

		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			hlog.Errorf("read failed: %s", err)
			return
		}
		_, _ = pw.Write(bs)

		if resp.StatusCode > 0 {
			exitCode = resp.StatusCode
		}
	}()
	return pr
}

func (d DataAPI) Type() string {
	return "api"
}

type Lang string

// DataCode run code
type DataCode struct {
	Lang       Lang   `json:"kind"`
	LangDesc   string `json:"langDesc" comment:"Lang"`
	ScriptCode string `json:"script_code" comment:"ScriptCode"`
}

// Run run shell
func (d DataCode) Run(ctx context.Context) (out io.ReadCloser) {
	pr, pw := io.Pipe()
	go func() {
		var (
			exitCode = DefaultExitCode
			err      error
			codePath string
			cmd      *exec.Cmd
		)
		defer func(pw *io.PipeWriter) { _ = pw.Close() }(pw)
		defer func() {
			now := time.Now().Local().Format("2006-01-02 15:04:05")
			finishTxt := fmt.Sprintf("#################### Bash Shell Finish ##################")
			_, _ = pw.Write([]byte(fmt.Sprintf("\n%s\n%s Task Run Finished, Return exitCode:%5d", finishTxt, now, exitCode))) // write exitCode,total 5 byte
			if codePath != "" {
				_ = os.Remove(codePath)
			}
		}()
		cmd, codePath, err = getCmd(ctx, string(d.Lang), d.ScriptCode)
		if err != nil {
			_, _ = pw.Write([]byte(err.Error()))
			return
		}
		cmd.Stdout = pw
		cmd.Stderr = pw
		startTxt := fmt.Sprintf("#################### Bash Shell Start ##################\n\n")
		_, _ = pw.Write([]byte(startTxt))
		err = cmd.Start()
		if err != nil {
			_, _ = pw.Write([]byte(err.Error()))
			return
		}

		err = cmd.Wait()
		if err != nil {
			// deal err
			// if context err,will change err to custom msg
			switch ctx.Err() {
			case context.DeadlineExceeded:
				_, err = pw.Write([]byte("ErrCtxDeadlineExceeded"))
				if err != nil {
					return
				}
			case context.Canceled:
				_, _ = pw.Write([]byte("ErrCtxCanceled"))
			default:
				_, err = pw.Write([]byte(err.Error()))
				if err != nil {
					return
				}
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
	return string(d.Lang)
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
	//case "nodejs":
	//	return runNodejs(ctx, code)
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
	err = tmpFile.Close()
	if err != nil {
		return nil, "", err
	}
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
	_ = tmpFile.Close()
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
	_ = tmpFile.Sync()
	_ = tmpFile.Close()
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

	tmpdir, err := ioutil.TempDir("", "jobor_")
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

	_ = tmpFile.Sync()
	_ = tmpFile.Close()
	cmd := exec.CommandContext(ctx, "cmd", "/C", batCodePath)
	return cmd, batCodePath, nil
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

func apiTest() {
	url := "www.baidu.com"
	method := "POST"

	payload := strings.NewReader(`{
    "key1": "v1"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("h1", "value1")
	req.Header.Add("h2", "value2")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "BAIDUID=865E9EA2BD4B781410C139B8CA1E3198:FG=1")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
