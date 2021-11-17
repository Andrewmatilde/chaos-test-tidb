package main

import (
	"chaos-client/chaosmesh"
	"context"
	"flag"
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
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

	exampleRestClient, err := chaosmesh.NewClientFor(config)

	result := v1alpha1.IOChaosList{}
	err = exampleRestClient.
		Get().
		Namespace("").
		Resource("iochaos").
		Do(context.TODO()).
		Into(&result)

	if err != nil {
		panic(err.Error())
	}
	for _, item := range result.Items {
		fmt.Println(item)
	}
}
