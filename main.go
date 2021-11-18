package main

import (
	"chaos-client/chaosmesh"
	"chaos-client/tidb"
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

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

	iochaosList, err := client.IoChaos("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	for _, iochaos := range (*iochaosList).Items {
		fmt.Println(iochaos)
	}
	dbClient := tidb.NewClient()
	err = tidb.InsertCase(dbClient)
	if err != nil {
		panic(err)
	}
}
