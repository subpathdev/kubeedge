# Certificate creation

## Edge / Cloud on the same system
You can use the same certificates in both programm path. We have written a script to greate these certificates for you.
In the script are set default values for the following varibles:
| variable | default value |
| -------- | ------------- |
| CA\_PATH | /etc/kubeedge/ca |
| CA\_SUBJECT | /C=CN/ST=Zhejiang/L=Hangzhou/O=KubeEdge/CN=kubeedge.io |
| CERT\_PATH | /etc/kubeedge/certs |
| SUBJECT | /C=CN/ST=Zhejiang/L=Hangzhou/O=KubeEdge/CN=kubeedge.io |

If you want to change them you can execute the following command:
```shell
export <vaiable name>=<value>
```

To generate the certifcates you execute the following command:
```shell
$GOPATH/src/github.com/kubeedge/kubeedge/build/tools/certgen.sh genCertAndKey
```

## Using different systems

This part describe to generate the certificates if you use a different system of the cloud and edge.
In this part you have a dedicated system for cloud and edge. (no part is running in a cloud)

In this scenario you have to create a root certificate and key on one system. A certificate request 
and a key on the other system. 
