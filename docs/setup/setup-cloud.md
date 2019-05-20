# Setting up cloud part

## Prequisites

+ [Install docker](https://docs.docker.com/install)

+ [Install kubeadm/kubectl](https://kubernetes.io/docs/setup/independent/install-kubeadm/)

+ [Creating cluster with kubeadm](<https://kubernetes.io/docs/setup/independent/create-cluster-kubeadm/>)

+ Installed Packages / Programmes:
	- openssl

+ After initializing Kubernetes master, we need to expose insecure port 8080 for edgecontroller/kubectl to work with http connection to Kubernetes apiserver.
  Please follow below steps to enable http port in Kubernetes apiserver.

    ```shell
    vi /etc/kubernetes/manifests/kube-apiserver.yaml
    # Add the following flags in spec: containers: -command section
    - --insecure-port=8080
    - --insecure-bind-address=0.0.0.0
    ```

+ (**Optional**)KubeEdge also supports https connection to Kubernetes apiserver. Follow the steps in [Kubernetes Documentation](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/) to create the kubeconfig file.

  Enter the path to kubeconfig file in controller.yaml
  ```yaml
  controller:
    kube:
      ...
      kubeconfig: "path_to_kubeconfig_file" #Enter path to kubeconfig file to enable https connection to k8s apiserver
  ```
### Certificates
[Certificate creation](./setup-certificates.md)

## Clone KubeEdge

```shell
git clone https://github.com/kubeedge/kubeedge.git $GOPATH/src/github.com/kubeedge/kubeedge
cd $GOPATH/src/github.com/kubeedge/kubeedge
```

## Run Cloud

### Run as a binary

+ Build cloud and edge

	```shell
	cd $GOPATH/src/github.com/kubeedge/kubeedge
	make
	```

+ Build Cloud only
	```shell
	cd $GOPATH/src/github.com/kubeedge/kubeedge
	make all WHAT=cloud
	```

+ The path to the generated certificates should be updated in `$GOPATH/src/github.com/kubeedge/kubeedge/cloud/conf/controller.yaml`. Please update the correct paths for the following :
    + cloudhub.ca
    + cloudhub.cert
    + cloudhub.key

+ Create device model and device CRDs.
    ```shell
    cd $GOPATH/src/github.com/kubeedge/kubeedge/build/crds/devices
    kubectl create -f devices_v1alpha1_devicemodel.yaml
    kubectl create -f devices_v1alpha1_device.yaml
    ```

+ Run cloud
    ```shell
    cd $GOPATH/src/github.com/kubeedge/kubeedge/cloud
    # run edge controller
    # `conf/` should be in the same directory as the cloned KubeEdge repository
    # verify the configurations before running cloud(edgecontroller)
    ./edgecontroller
    ```

### [Run as Kubernetes deployment](../../build/cloud/README.md)
