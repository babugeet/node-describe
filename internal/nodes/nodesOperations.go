package nodes

import (
	"context"
	"fmt"
	"node-describe/internal/kubeclient"
	"os"
	"text/tabwriter"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateKubeclient() *kubernetes.Clientset {
	client := kubeclient.NewClient()
	clientset := client.CreateClientObject()
	return clientset
}

func GetNodes(clientset *kubernetes.Clientset) []v1.Node {
	// client := kubeclient.NewClient()
	// clientset := client.CreateClientObject()

	// List pods
	nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	// fmt.Println(nodeList)

	// var target map[string]any

	// _ = json.Unmarshal([]byte(nodeList.Items), &target)
	// for i, k := range nodeList.Items {
	// 	fmt.Println(i)
	// 	fmt.Println(k.Name)
	// }
	return nodeList.Items
}

func GetPods(nodeList []v1.Node, client *kubernetes.Clientset) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	// tabwriter.Writerzzz
	// Write table header
	fmt.Fprintln(writer, "Node\tCPU Request\tCPU Limit\tMemory Request\tMemory Limit\t")
	for _, node := range nodeList {
		// fmt.Println(node.Name)

		var nodeCPUreq, nodeCPUlimit, nodeMemReq, nodeMemLimit int64
		// fmt.Println("               ")
		podList, _ := GetPodsByNode(node, client)
		for _, pod := range podList.Items {
			totalCPULimit, totalCPURequest, totalMemLimit, totalMemRequest := GetPodResource(pod)
			nodeCPUreq = nodeCPUreq + totalCPURequest
			nodeCPUlimit = nodeCPUlimit + totalCPULimit
			nodeMemReq = nodeMemReq + totalMemLimit
			nodeMemLimit = nodeMemLimit + totalMemRequest
		}
		// fmt.Println(nodeCPUreq, node.Status.Allocatable.Cpu().MilliValue())
		// fmt.Println(node.Status.Allocatable.Memory().MilliValue())
		pernodeCPUreq := CalPercentageUsage(nodeCPUreq, node.Status.Allocatable.Cpu().MilliValue())
		pernodeCPUlimit := CalPercentageUsage(nodeCPUlimit, node.Status.Allocatable.Cpu().MilliValue())
		pernodeMemReq := CalPercentageUsage(nodeMemReq, node.Status.Allocatable.Memory().MilliValue())
		pernodeMemLimit := CalPercentageUsage(nodeMemLimit, node.Status.Allocatable.Memory().MilliValue())
		// fmt.Print(no÷deCPUreq, nodeCPUlimit, nodeMemReq, nodeMemLimit)
		fmt.Fprintf(writer, "%s\t%f\t%f\t%f\t%f\t\n", node.Name, pernodeCPUreq, pernodeCPUlimit, pernodeMemReq, pernodeMemLimit)
	}
	writer.Flush()
}

func CalPercentageUsage(nodeuse int64, nodemax int64) float64 {

	return (float64(nodeuse) / float64(nodemax)) * 100
}

func GetPodsByNode(node v1.Node, client *kubernetes.Clientset) (*v1.PodList, error) {
	return client.CoreV1().Pods(v1.NamespaceAll).List(context.TODO(), metav1.ListOptions{
		FieldSelector: "spec.nodeName=" + node.Name + ",status.phase!=Failed,status.phase!=Succeeded",
	})

}

func GetPodResource(podName v1.Pod) (int64, int64, int64, int64) {
	containerList := podName.Spec.Containers
	totalCPULimit, totalCPURequest, totalMemLimit, totalMemRequest := GetContainerResource(containerList)
	return totalCPULimit, totalCPURequest, totalMemLimit, totalMemRequest
}

func GetContainerResource(containerList []v1.Container) (int64, int64, int64, int64) {
	var totalCPULimit, totalCPURequest, totalMemLimit, totalMemRequest int64

	for _, j := range containerList {
		cpuLimit, cpuRequest := cpuLimitRequests(j)
		memLimit, memRequest := memoryLimitRequests(j)
		totalCPURequest = cpuRequest + totalCPURequest
		totalCPULimit = cpuLimit + totalCPULimit
		totalMemLimit = memLimit + totalMemLimit
		totalMemRequest = memRequest + totalMemRequest
	}
	return totalCPULimit, totalCPURequest, totalMemLimit, totalMemRequest
}

func cpuLimitRequests(containerName v1.Container) (int64, int64) {
	// containerName.Resources.Requests.Cpu().MilliValue()
	return containerName.Resources.Limits.Cpu().MilliValue(), containerName.Resources.Requests.Cpu().MilliValue()
}

func memoryLimitRequests(containerName v1.Container) (int64, int64) {
	return containerName.Resources.Limits.Memory().MilliValue(), containerName.Resources.Requests.Memory().MilliValue()
}

// fieldSelector=spec.nodeName%3Dk8clusters33.fyre.ibm.com%2Cstatus.phase%21%3DFailed%2Cstatus.phase%21%3DSucceeded
func DescribeNode() {
	client := CreateKubeclient()
	nodeList := GetNodes(client)
	GetPods(nodeList, client)

}
