/**
 * @Author: jijunhua@bytedance.com
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2022/5/2 17:38
 */

package main

import (
	"fmt"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//create config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	//create client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	//create informer in namespace "default"
	//factory := informers.NewSharedInformerFactory(clientset, 0)	 // in all namespace
	factory := informers.NewSharedInformerFactoryWithOptions(clientset, 0, informers.WithNamespace("default"))

	// 拿到Pod对应的informer，注册EventHandler
	// 当Pod发生变化时，会触发EventHandler
	// 未注册的handler，不会触发，可以在 `ResourceEventHandler`中看到
	informer := factory.Core().V1().Pods().Informer()

	// add workqueue
	//rateLimitingQueue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "my_test_queue")

	//register event handler
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("add Event")
			//key, err := cache.MetaNamespaceIndexFunc(obj)
			//if err != nil {
			//	fmt.Println("err")
			//}
			//rateLimitingQueue.Add(key)
		},
		UpdateFunc: func(old, new interface{}) {
			fmt.Println("update Event")
			//key, err := cache.MetaNamespaceIndexFunc(new)
			//if err != nil {
			//	fmt.Println("err")
			//}
			//rateLimitingQueue.Add(key)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("delete Event")
			//key, err := cache.MetaNamespaceIndexFunc(obj)
			//if err != nil {
			//	fmt.Println("err")
			//}
			//rateLimitingQueue.Add(key)
		},
	})

	//start factory
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	// 等待同步完成
	factory.WaitForCacheSync(stopCh)
	<-stopCh
}
