package main

import (
	"os"

	"github.com/argoproj/argo-events/common"
	"github.com/argoproj/argo-events/gateways/server"
	"github.com/argoproj/argo-events/gateways/server/hdfs"
	"k8s.io/client-go/kubernetes"
)

func main() {
	kubeConfig, _ := os.LookupEnv(common.EnvVarKubeConfig)
	restConfig, err := common.GetClientConfig(kubeConfig)
	if err != nil {
		panic(err)
	}
	namespace, ok := os.LookupEnv(common.EnvVarNamespace)
	if !ok {
		panic("namespace is not provided")
	}
	clientset := kubernetes.NewForConfigOrDie(restConfig)
	server.StartGateway(&hdfs.EventListener{
		Logger:    common.NewArgoEventsLogger(),
		K8sClient: clientset,
		Namespace: namespace,
	})
}
