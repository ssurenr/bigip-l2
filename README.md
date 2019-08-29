# BIGIP L2 OPS
A Go package for handling BIG-IP L2 operations. Go implementation of F5 CCCL Net Schema.

## Usage

Download the respository and run the main.go file
```shell
go get github.com/ssurenr/bigip-l2
cd $GOPATH/src/github.com/ssurenr/bigip-l2/
go run main.go
```

The program lists the parsed data from the arp.json file under the sample directory
```
Name                 IP Address           Mac Address
k8s-10.244.1.234     10.244.1.234         56:69:31:1d:30:62
k8s-10.244.1.232     10.244.1.232         56:69:31:1d:30:63
```
