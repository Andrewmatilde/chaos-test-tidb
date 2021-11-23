package sysbench

import (
	"context"
	"go.uber.org/zap"
	"log"
	"os"
	"os/exec"
	"strings"
)

func SysbenchPrepare() *exec.Cmd {
	return exec.Command("sysbench", strings.Split("oltp_insert prepare --mysql-port=31010 --mysql-host=127.0.0.1 --mysql-user=root --tables=32 --table-size=100000000 --num-threads=64", " ")...)
}

func SysbenchCleanup() *exec.Cmd {
	return exec.Command("sysbench", strings.Split("oltp_insert cleanup --mysql-port=31010 --mysql-host=127.0.0.1 --mysql-user=root --tables=32 --table-size=100000000 --num-threads=64", " ")...)
}

func RunSysbench(ctx context.Context) error {
	cmdPrepare := SysbenchPrepare()
	cmdCleanup := SysbenchCleanup()
	log.Println("sysbench creating")
	_ = execWithTimeout(ctx, cmdPrepare)
	log.Println("sysbench cleaning")
	output, err := cmdCleanup.CombinedOutput()
	if err != nil {
		return err
	}
	log.Println(string(output))
	return nil
}

func execWithTimeout(ctx context.Context, cmd *exec.Cmd) error {
	done := make(chan error, 1)
	go func() {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		done <- err
	}()

	select {
	case <-ctx.Done():
		if err := cmd.Process.Kill(); err != nil {
			log.Println("failed to kill process: ", zap.Error(err))
			return err
		}
	case err := <-done:
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
