package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func example1() {
	// out of cluster

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error %s, building config from flags\n", err.Error())
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s, creating clientset\n", err.Error())
		panic(err.Error())
	}

	// for {

	// List the pods (Pods belongs to CoreV1())
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("error %s, while listing all the pods\n", err.Error())
		panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	fmt.Println("Pods in the default namespace:")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	ctx := context.Background()
	//List the deployments (Deployments belongs to AppsV1())
	fmt.Println("\nList of deployments:")
	deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s, while listing the deployments")
	}
	for _, d := range deployments.Items {
		fmt.Println(d.Name)
	}
}
