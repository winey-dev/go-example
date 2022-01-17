package main

import (
	"context"
	"log"
	"os"

	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatalf("Config build fail %v\n", err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	ing, err := clientset.NetworkingV1beta1().Ingresses("").List(context.Background(), v1meta.ListOptions{})
	if err != nil {
		log.Printf("NetworkingV1beta1() Faill err=%v\n", err)
	} else {
		log.Printf("v1beta: %v\n", ing)
	}

	ing2, err := clientset.NetworkingV1().Ingresses("").List(context.Background(), v1meta.ListOptions{})
	if err != nil {
		log.Printf("NetworkingV1beta1() Faill err=%v\n", err)
	} else {
		log.Printf("v1: %v\n", ing2)
	}
}
