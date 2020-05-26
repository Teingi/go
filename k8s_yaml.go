package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//这段代码等同于kubectl apply -f k8s-test.yaml
func main() {
	var (
		err    error
		data   []byte
		conf   *rest.Config
		client *kubernetes.Clientset
	)
	if data, err = ioutil.ReadFile("/Users/5bug/codes/projects/k8s-demo/k8s-test.yaml"); err != nil {
		fmt.Print(err)
	}
	if data, err = yaml2.ToJSON(data); err != nil {
		return
	}
	deployment := &v1.Deployment{}
	if err = json.Unmarshal(data, deployment); err != nil {
		return
	}
	cluster := deployment.ObjectMeta.ClusterName
	namespace := deployment.ObjectMeta.Namespace
	deploymentName := deployment.ObjectMeta.Name
	if conf, err = clientcmd.BuildConfigFromFlags("", "/Users/5bug/.kube/config"); err != nil {
		log.Println("BuildConfigFromFlags", err)
		return
	}
	if client, err = kubernetes.NewForConfig(conf); err != nil {
		log.Println("NewForConfig", err)
		return
	}
	if deployment, err = client.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{}); err != nil {
		return
	}
	fmt.Println(cluster, namespace, deploymentName)
	fmt.Println(deployment)
	return
}
