package main

import (
	"flag"
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	exampleType := flag.String("type", "terminal", "--type=terminal or --type=log")
	ns := flag.String("ns", "nginx-ingress", "--ns=nginx-ingress (NameSpace)")
	pn := flag.String("pn", "ingress-controller-57b6745d-j89fm", "--pn=ingress-controller-57b6745d-j89fm (PodName)")
	cn := flag.String("cn", "", "--cn=ingress-controller (ContainerName)")
	flag.Parse()

	log.Printf("Flag Pasrse ExampleType=[%s] Namespace=[%s] PodName=[%s] ContainerName=[%s]\n",
		*exampleType, *ns, *pn, *cn)

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		log.Fatalf("Config build fail %v\n", err)
	}

	cliset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("NewForConfig Fail %v\n", err)
	}

	if *exampleType == "terminal" {
		err = GetTerminal(cliset, config, *ns, *pn, *cn, "bash")
		if err != nil {
			log.Fatalf("Get Terminal Fail %v\n", err)
		}
	} else if *exampleType == "log" {
		err = GetPodLogs(cliset, *ns, *pn, *cn, true)
		if err != nil {
			log.Fatalf("GetPodLods Fail %v\n", err)
		}
	} else {
		log.Fatalf("Not Support Type %s\n", *exampleType)
	}

	return
}
