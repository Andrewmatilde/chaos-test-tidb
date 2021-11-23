package main

import (
	"chaos-client/chaosmesh"
	"chaos-client/sysbench"
	"chaos-client/tidb"
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type t struct {
	start      time.Time
	chaosStart time.Time
	end        time.Time
}

func ApplyIochaosToKV(client *chaosmesh.Client,duration time.Duration, target int) error {
	chaos := chaosmesh.IoChaosForTikv("iochaos"+strconv.Itoa(target), "tidb-c0", "advanced-tidb-tikv-"+strconv.Itoa(target))

	log.Println("Creating Iochaos : ", chaos.Name)
	chaosr, err := client.IoChaos("tidb-c0").Create(context.TODO(), &chaos, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	time.Sleep(duration)
	log.Println("Deleting Iochaos : ", chaos.Name)
	err = client.IoChaos("tidb-c0").Delete(context.TODO(), chaosr.Name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	client, err := chaosmesh.NewClientFor(config)
	if err != nil {
		panic(err)
	}
	tl := make([]t, 3)
	for i := 0; i < 3; i++ {
		_ = tidb.NewClient()
		tl[i].start = time.Now()
		ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(time.Minute * 9))

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			err = sysbench.RunSysbench(ctx)
			wg.Done()
		}()
		time.Sleep(time.Minute * 3)
		tl[i].chaosStart = time.Now()

		err := ApplyIochaosToKV(client, time.Minute * 3, i)
		if err != nil {
			panic(err)
		}
		wg.Wait()
		cancel()
		tl[i].end = time.Now()
		if err != nil {
			panic(err)
		}
	}

	for _, tt := range tl {
		fmt.Println(tt.start.Format("2006-01-02 03:04:05 PM"),
			tt.chaosStart.Format("2006-01-02 03:04:05 PM"),
			tt.end.Format("2006-01-02 03:04:05 PM"))
	}

}
