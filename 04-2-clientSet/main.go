/**
 * @Author: jijunhua@bytedance.com
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2022/4/15 17:19
 */

package main

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	// clientSet中会指定config，所以不用手动设置

	// clientSet, 多个client的集合
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// clientSet找到coreV1
	coreV1 := clientSet.CoreV1()
	pod, err := coreV1.Pods("default").Get(context.TODO(), "test", v1.GetOptions{})
	if err != nil {
		println("\nerr:" + err.Error())
	} else {
		println("\npodName: " + pod.Name)
	}
}
