package main

import (
	"context"
	"fmt"
	"io"

	v1 "k8s.io/api/core/v1"
	k8s "k8s.io/client-go/kubernetes"
)

func GetPodLogs(cliset *k8s.Clientset, namespace string, podName string, containerName string, follow bool) error {
	count := int64(100)
	podLogOptions := v1.PodLogOptions{
		Container: containerName,
		Follow:    follow,
		TailLines: &count,
	}

	podLogRequest := cliset.CoreV1().
		Pods(namespace).
		GetLogs(podName, &podLogOptions)
	stream, err := podLogRequest.Stream(context.TODO())
	if err != nil {
		return err
	}
	defer stream.Close()

	for {
		buf := make([]byte, 2000)
		numBytes, err := stream.Read(buf)
		if numBytes == 0 {
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		message := string(buf[:numBytes])
		fmt.Print(message)
	}
	return nil
}
