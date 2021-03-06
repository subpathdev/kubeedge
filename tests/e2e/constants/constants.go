package constants

const (
	RunController = "cd ${GOPATH}/src/github.com/kubeedge/kubeedge/cloud/; sudo nohup ./edgecontroller > edgecontroller.log 2>&1 &"
	RunEdgecore   = "cd ${GOPATH}/src/github.com/kubeedge/kubeedge/edge/; sudo nohup ./edgecore > edgecore.log 2>&1 &"
	RunEdgeSite   = "cd ${GOPATH}/src/github.com/kubeedge/kubeedge/edgesite/; sudo nohup ./edgesite > edgesite.log 2>&1 &"

	AppHandler        = "/api/v1/namespaces/default/pods"
	NodeHandler       = "/api/v1/nodes"
	DeploymentHandler = "/apis/apps/v1/namespaces/default/deployments"
	ConfigmapHandler  = "/api/v1/namespaces/default/configmaps"
	ServiceHandler    = "/api/v1/namespaces/default/services"
	CrdHandler        = "/apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions"
)
