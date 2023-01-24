package k8s

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

type K8sConfig struct {
	Name  string
	Host  string
	Token string
	CA    string
}

var (
	k3sconfig K8sConfig = K8sConfig{
		Name: "k3s",
		Host: "https://jiuyebeso.top:6443",
		Token: "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBBNmlxNWEyNVBaZktsWEh0WjdaRlpQM2ZrOWN4blJvVi1mcEZYTHlCencifQ.eyJpc3MiOiJr" +
			"dWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbS" +
			"IsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkZWZhdWx0LXRva2VuLWp6NDVyIiwia3ViZXJuZXRlcy5p" +
			"by9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImRlZmF1bHQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3N" +
			"lcnZpY2UtYWNjb3VudC51aWQiOiJkMzk2NmY2MC0wNWE5LTQ3MWItODRlNS0yMGExOTVhODczNzYiLCJzdWIiOiJzeXN0ZW06c2VydmljZW" +
			"FjY291bnQ6a3ViZS1zeXN0ZW06ZGVmYXVsdCJ9.Tga0gvXYVbCxcQ4HLoPOcxNq41AZxGoAgQ3gF3m_YFa0k8dGNCzGHREVBSJxUM2a8sr5B" +
			"yigMihsgzv_hWdxwTh9adtfs3MaZvPiA9WTl1pK68bQrbrstQTv74PnqEQcm5_whQqJA98yRE1zNwah4-YqycVD8FyyDJVKnafvGL36Y_qc8i" +
			"T0wL3_XSO44IzXDIxU6lD199qfDOhDU0qlOBO2bZcrrayzppgDTJ8AeHIsAK1VXXTczqdBgQDdch2gDWWl-JtnSShE3R4gZgH29F3pzTE1x4RG" +
			"L-AFg6DlMdwkb39hNijxEMhfZOdA_9HsCmBpqUA4eAAT0HIka-QQtg",
		CA: k3sCa,
	}
	miniconfig K8sConfig = K8sConfig{
		Name:  "minikube",
		Host:  "",
		Token: "",
		CA:    miniCa,
	}
)

const k3sCa = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJkekNDQVIyZ0F3SUJBZ0lCQURBS0JnZ3Foa2pPUFFRREFqQWpNU0V3SHdZRFZRUUREQmhyTTNNdGMyVnkKZG1WeUxXTmhRREUyTlRRM05UYzBNakV3SGhjTk1qSXdOakE1TURZMU1ESXhXaGNOTXpJd05qQTJNRFkxTURJeApXakFqTVNFd0h3WURWUVFEREJock0zTXRjMlZ5ZG1WeUxXTmhRREUyTlRRM05UYzBNakV3V1RBVEJnY3Foa2pPClBRSUJCZ2dxaGtqT1BRTUJCd05DQUFSdWJrYUlzZmo4ekFoNHlEaytiVGR3NVNLeVd2Qk1Ma25YdW9QVzBkTDYKWi9WZkYzZXArL0YyWENhQzdqZ0tDek1aWnVpLzJqREd5TnVpRTRjajNTWFhvMEl3UURBT0JnTlZIUThCQWY4RQpCQU1DQXFRd0R3WURWUjBUQVFIL0JBVXdBd0VCL3pBZEJnTlZIUTRFRmdRVUo3RmRORUw0aVZ2dWQ1YkE3N0R5CklIRFhKYVl3Q2dZSUtvWkl6ajBFQXdJRFNBQXdSUUlnQ0NETTBuU1F6WlNRSXJPMmNBTzhQSEdHT1BlY0wwUDkKN0ZMWHRvWmRhS2NDSVFDY3dNWkVxMHdGUWJVNkFER2ZQaHNzWmszM0ZvOWFjdmM3NmxqenZYT1lrUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
const miniCa = ""

func SelectClusterConfig(env string) (*rest.Config, error) {
	var c K8sConfig
	switch env {
	case "k3s":
		c = k3sconfig
	case "minikube":
		c = miniconfig
	default:
		log.Printf("环境: %s 不支持", env)
		return nil, fmt.Errorf("环境: %s 不支持", env)
	}

	return &rest.Config{
		Host:            c.Host,
		BearerToken:     c.Token,
		BearerTokenFile: "",
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
			//CAData: []byte(c.CA),
		},
	}, nil
}

func GetPodByName(env, namespace, podname string) {
	config, err := SelectClusterConfig(env)

	if err != nil {
		log.Println("xxxxx", err)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("xxxxx", err)
		return
	}

	pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podname, metav1.GetOptions{})
	if err != nil {
		log.Println("xxxxx", err)
		return
	}
	log.Println("pod := ", pod.Status.PodIP)

	node, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Println("xxxxx", err)
		return
	}
	log.Println("node := ", node.Items[0].Status.Addresses)
}

func GetPods(namespace string) {
	// 在 kubeconfig 中使用当前上下文
	// path-to-kubeconfig -- 例如 /root/.kube/config
	//config, _ := clientcmd.BuildConfigFromFlags("", "/Users/yizhiya/.kube/config-k3s")
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/yizhiya/.kube/config-k3s")
	// 创建 clientset
	clientset, _ := kubernetes.NewForConfig(config)
	// 访问 API 以列出 Pod
	pods, _ := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	for _, pod := range pods.Items {
		fmt.Println("pod name: ", pod.Name)
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}

func getPods() {

	clientcmd.BuildConfigFromKubeconfigGetter("", nil)
}
