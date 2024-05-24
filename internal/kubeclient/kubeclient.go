package kubeclient

import (
	"node-describe/constants"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Kubeclient struct {
}

func NewClient() *Kubeclient {
	return &Kubeclient{}
}

func getKubeConfig() string {
	term := constants.GetCfgFile()
	// fmt.Println("the term is ", term)
	return term
}

func (k *Kubeclient) CreateClientObject() *kubernetes.Clientset {

	kubeconfig := getKubeConfig()
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset

}
