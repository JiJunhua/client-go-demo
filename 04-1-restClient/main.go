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

	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	// client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// get data of pod
	pod := v1.Pod{}
	// Do之前的都是在拼接url,
	// do发送执行命令， 实际执行发送的url，返回都在这里完成
	// 获取的结果放到pods中（手动指定的方式）

	err = restClient.Get().Namespace("default").Resource("pods").Name("test").Do(context.TODO()).Into(&pod)
	if err != nil {
		println("\nerr:" + err.Error())
	} else {
		println("\npodName: " + pod.Name)
	}
}
