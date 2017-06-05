package main

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func vegetaserversOnKube() []string {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{
		LabelSelector: "app=vegetaserver",
	})

	if err != nil {
		panic(err.Error())
	}

	var servers = make([]string, 0)

	for _, pod := range pods.Items {
		if pod.Status.PodIP == "" {
			continue
		}
		servers = append(servers, fmt.Sprintf("http://%v:8081/", pod.Status.PodIP))
	}

	return servers
}
