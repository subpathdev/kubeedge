# Setting up cloud usage

## Prerequisites

+ [Install docker](https://docs.docker.com/install/)

+ Installed Packages / Programmes:
	- openssl

### Clone KubeEdge

```shell
git clone https://github.com/kubeedge/kubeedge.git $GOPATH/src/github.com/kubeedge/kubeedge
cd $GOPATH/src/github.com/kubeedge/kubeedge
```
### Configuring MQTT mode

The Edge part of KubeEdge uses MQTT for communication between deviceTwin and devices. KubeEdge supports 3 MQTT modes:
1) internalMqttMode: internal mqtt broker is enabled.
2) bothMqttMode: internal as well as external broker are enabled.
3) externalMqttMode: only external broker is enabled.

Use mode field in [edge.yaml](https://github.com/kubeedge/kubeedge/blob/master/edge/conf/edge.yaml#L4) to select the desired mode.

To use KubeEdge in double mqtt or external mode, you need to make sure that [mosquitto](https://mosquitto.org/) or [emqx edge](https://www.emqx.io/downloads/emq/edge?osType=Linux#download) is installed on the edge node as an MQTT Broker.

### Certificates
[Certificate creation](./setup-certificates.md)

## Run KubeEdge Edge
### Run as a binary
+ Build Edge

    ```shell
    cd $GOPATH/src/github.com/kubeedge/kubeedge
    make all WHAT=edge
    ```

    KubeEdge can also be cross compiled to run on ARM based processors.
    Please follow the instructions given below or click [Cross Compilation](../setup/cross-compilation.md) for detailed instructions.

    ```shell
    cd $GOPATH/src/github.com/kubeedge/kubeedge/edge
    make edge_cross_build
    ```

    KubeEdge can also be compiled with a small binary size. Please follow the below steps to build a binary of lesser size:

    ```shell
    apt-get install upx-ucl
    cd $GOPATH/src/github.com/kubeedge/kubeedge/edge
    make edge_small_build
    ```

    **Note:** If you are using the smaller version of the binary, it is compressed using upx, therefore the possible side effects of using upx compressed binaries like more RAM usage, 
    lower performance, whole code of program being loaded instead of it being on-demand, not allowing sharing of memory which may cause the code to be loaded to memory 
    more than once etc. are applicable here as well.

+ Modify the `$GOPATH/src/github.com/kubeedge/kubeedge/edge/conf/edge.yaml` configuration file
    + Replace `edgehub.websocket.certfile` and `edgehub.websocket.keyfile` with your own certificate path
    + Update the IP address of the master in the `websocket.url` field. 
    + replace `fb4ebb70-2783-42b8-b3ef-63e2fd6d242e`q with edge node name in edge.yaml for the below fields :
        + `websocket:URL`
        + `controller:node-id`
        + `edged:hostname-override`

+ Run edge

    ```shell
    # run mosquitto
    mosquitto -d -p 1883
    # or run emqx edge
    # emqx start
    
    # run edge_core
    # `conf/` should be in the same directory as the cloned KubeEdge repository
    # verify the configurations before running edge(edge_core)
    ./edge_core
    # or
    nohup ./edge_core > edge_core.log 2>&1 &
    ```

    **Note:** Please run edge using the users who have root permission.
### [Run as container](../../build/edge/README.md)
### [Run as Kubernetes deployment](../../build/edge/kubernetes/README.md)
