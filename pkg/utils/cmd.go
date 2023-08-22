package utils

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

type CmdResp struct {
	ExitCode int    `json:"exitCode"`
	Out      string `json:"out"`
	ErrOut   string `json:"errOut"`
	Stdout   string `json:"stdOut"`
}

// CheckLsExists 预先检查命令是否存在
func CheckLsExists(cmd string) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Printf("didn't find 'ls' executable\n")
	} else {
		fmt.Printf("'ls' executable is in '%s'\n", path)
	}
}

func read(ctx context.Context, wg *sync.WaitGroup, std io.ReadCloser, stdAndErr, out *string) {
	reader := bufio.NewReader(std)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Print(readString)
			*out = *out + readString
		}
	}
}

// Command 执行命令实时打印
func Command(cmd string, timeout int64, f *os.File) (CmdResp, error) {
	//c := exec.CommandContext(ctx, "cmd", "/C", cmd) // windows
	//fmt.Println("split:", strings.Split(cmd, "\\n"))
	if timeout <= 0 {
		timeout = 10
	}
	ctx, cancel := context.WithCancel(context.Background())
	go func(cancelFunc context.CancelFunc) {
		time.Sleep(time.Duration(timeout) * time.Second)
		cancelFunc()
	}(cancel)
	var result = CmdResp{}
	c := exec.CommandContext(ctx, "bash", "-cxe", cmd) // mac linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		result.ErrOut = err.Error()
		result.ExitCode = c.ProcessState.ExitCode()
		return result, err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		result.ErrOut = err.Error()
		result.ExitCode = c.ProcessState.ExitCode()
		return result, err
	}
	var wg sync.WaitGroup
	// 因为有2个任务, 一个需要读取stderr 另一个需要读取stdout
	wg.Add(2)
	go func() {
		reader := bufio.NewReader(stderr)
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				readString, err := reader.ReadString('\n')
				if err != nil || err == io.EOF {
					return
				}
				//fmt.Print(readString)
				result.ErrOut = result.ErrOut + readString
				result.Out = result.Out + readString
				if f != nil {
					_, _ = f.Write([]byte(readString))
				}
			}
		}
	}()

	go func() {
		reader := bufio.NewReader(stdout)
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				readString, err := reader.ReadString('\n')
				if err != nil || err == io.EOF {
					return
				}
				//fmt.Print(readString)
				result.Stdout = result.Stdout + readString
				result.Out = result.Out + readString
				if f != nil {
					_, _ = f.Write([]byte(readString))
				}
			}
		}
	}()
	//go read(ctx, &wg, stderr, &result.ErrOut)
	//go read(ctx, &wg, stdout, &result.Stdout)
	// 这里一定要用start,而不是run 详情请看下面的图
	err = c.Start()
	if err != nil {
		result.ErrOut = err.Error()
		result.ExitCode = c.ProcessState.ExitCode()
		return result, err
	}

	// 等待任务结束
	wg.Wait()
	//fmt.Println("c.ProcessState:", c.ProcessState.String())
	result.ExitCode = c.ProcessState.ExitCode()
	return result, err
}

// CommandOut 执行命令一次性返回结果
func CommandOut(cmd string, timeout int64) (CmdResp, error) {
	//c := exec.CommandContext(ctx, "cmd", "/C", cmd) // windows
	//fmt.Println("split:", strings.Split(cmd, "\\n"))
	if timeout <= 0 {
		timeout = 10
	}
	ctx, cancel := context.WithCancel(context.Background())
	go func(cancelFunc context.CancelFunc) {
		time.Sleep(time.Duration(timeout) * time.Second)
		cancelFunc()
	}(cancel)
	var result = CmdResp{}
	c := exec.CommandContext(ctx, "bash", "-cxe", cmd) // mac linux

	//output, err := c.CombinedOutput()
	//if err != nil {
	//	result.ErrOut = err.Error()
	//	result.ExitCode = c.ProcessState.ExitCode()
	//	return result, err
	//}
	//result.Stdout = string(output)

	c.Stdout = os.Stdout
	//c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		result.ErrOut = err.Error()
		result.ExitCode = c.ProcessState.ExitCode()
		return result, err
	}

	result.ExitCode = c.ProcessState.ExitCode()
	return result, err
}
