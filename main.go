package main

import (
	"chaos-client/chaosmesh"
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
	"strconv"
	"time"
)

type t struct {
	start      time.Time
	chaosStart time.Time
	end        time.Time
}

func ApplyIochaosToKV(client *chaosmesh.Client, duration time.Duration, target int) error {
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
	chaos := client.HTTPChaos("default")
	list, _ := chaos.List(context.TODO(), metav1.ListOptions{})
	fmt.Println(list)

	httpchaos := chaosmesh.Hhhchaos("star", "default")
	result, err := chaos.Create(context.TODO(), &httpchaos, metav1.CreateOptions{})
	fmt.Println(result)
	fmt.Println(err)
}
