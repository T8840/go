package os

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func GetProjectPath() string {
	dir, _ := os.Getwd()
	return dir
}

func GetLocalDirPath() string {
	dir, _ := os.Getwd()
	comma := strings.Index(dir, "godemo")
	local_dir_path := dir[:comma]
	return local_dir_path
}

func GetExecCmd(cmd *exec.Cmd) string{
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout  // 标准输出
	cmd.Stderr = &stderr  // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return outStr
}

func CmdAndChangeDir(dir string, commandName string, params []string) (string, error) {
	cmd := exec.Command(commandName, params...)
	fmt.Println("CmdAndChangeDir", dir, cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}

func WatchExecCmdDir(dir string, commandName string, params []string) (string) {
	cmd := exec.Command(commandName, params...)
	fmt.Println("CmdAndChangeDir", dir, cmd.Args)
	cmd.Dir = dir
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start failed with %s\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = CopyAndCapture(os.Stdout, stdoutIn)
		wg.Done()
	}()

	stderr, errStderr = CopyAndCapture(os.Stderr, stderrIn)
	wg.Wait()

	if err = cmd.Wait(); err != nil {
		log.Fatalf("cmd.Wait failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatalf("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	log.Fatalf("out: %s\n err: %s\n", outStr, errStr)
	return outStr
}

func WatchExecCmd(cmd *exec.Cmd) string{
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start failed with %s\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = CopyAndCapture(os.Stdout, stdoutIn)
		wg.Done()
	}()

	stderr, errStderr = CopyAndCapture(os.Stderr, stderrIn)
	wg.Wait()

	if err = cmd.Wait(); err != nil {
		log.Fatalf("cmd.Wait failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatalf("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	log.Fatalf("out: %s\n err: %s\n", outStr, errStr)
	return outStr
}

func CopyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func RunGitCmd(git_path string,name string,  arg ...string)(string,error) {
	cmd :=exec.Command(name,arg...)
	cmd.Dir = git_path
	msg,err := cmd.CombinedOutput() //混合输出stdout + stderr
	cmd.Run()
	// 报错时exit status 1
	return string(msg),err
}


func GetLocalIp()string {
	return ""
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}