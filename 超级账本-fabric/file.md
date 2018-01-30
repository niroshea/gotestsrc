1、安装go语言环境，设置GOROOT，GOPATH

2、安装git环境，设置PATH环境变量

3、安装node.js，设置环境变量

4、安装docker 以及 docker-compose

5、在GOPATH创建文件夹
mkdir -p $GOPATH/src/src/github.com/hyperledger

6、切换到 $GOPATH/src/src/github.com/hyperledger
git clone https://github.com/hyperledger/fabric
git clone https://github.com/hyperledger/fabric-ca

7、执行make报错：
cp: cannot stat ‘build/docker/gotools/bin/protoc-gen-go’: No such file or directory
解决：
go get -u github.com/golang/protobuf/protoc-gen-go
执行make && make install
将$GOPATH/bin 下的 protoc-gen-go 拷贝到$GOPATH/src/github.com/hyperledger/fabric/build/docker/gotools/bin/ 下。

8、执行make报错：
ltdl.h: No such file or directory
解决：
如果在ubunt操作系统中，只需安装：
apt install libltdl3-dev
如果在centos操作系统中，只需安装：
yum install libtool-ltdl-devel

9、近期只能使用go1.8进行编译，如果是1.9会出现一个warning，出现意料意外的错误。
cd $GOPATH/src/github.com/hyperledger/fabric
ARCH=x86_64
BASEIMAGE_RELEASE=0.3.1
PROJECT_VERSION=1.0.0
LD_FLAGS="-X github.com/hyperledger/fabric/common/metadata.Version=${PROJECT_VERSION} \
-X github.com/hyperledger/fabric/common/metadata.BaseVersion=${BASEIMAGE_RELEASE} \
-X github.com/hyperledger/fabric/common/metadata.BaseDockerLabel=org.hyperledger.fabric \
-X github.com/hyperledger/fabric/common/metadata.DockerNamespace=hyperledger \
-X github.com/hyperledger/fabric/common/metadata.BaseDockerNamespace=hyperledger"
CGO_CFLAGS=" " go install -ldflags "$LD_FLAGS -linkmode external -extldflags '-static -lpthread' " \
github.com/hyperledger/fabric/peer

===========================================
go install -ldflags " -linkmode external -extldflags ' -static -lpthread '" github.com/hyperledger/fabric-ca/cmd/fabric-ca-client
go install -ldflags " -linkmode external -extldflags ' -static -lpthread '" github.com/hyperledger/fabric-ca/cmd/fabric-ca-server

#编译安装cryptogen
PROJECT_VERSION=1.0.0
CGO_CFLAGS=" " \
go install -tags "" \
-ldflags "-X github.com/hyperledger/fabric/common/tools/cryptogen/metadata.Version=${PROJECT_VERSION}" \
github.com/hyperledger/fabric/common/tools/cryptogen

#编译安装 configtxgen
CGO_CFLAGS=" " \
go install -tags "nopkcs11" \
-ldflags "-X github.com/hyperledger/fabric/common/configtx/tool/configtxgen/metadata.Version=${PROJECT_VERSION}" \
github.com/hyperledger/fabric/common/configtx/tool/configtxgen

#编译安装 configtxlator
PROJECT_VERSION=1.0.0
CGO_CFLAGS=" " \
go install -tags "" \
-ldflags "-X github.com/hyperledger/fabric/common/tools/configtxlator/metadata.Version=${PROJECT_VERSION}" \
github.com/hyperledger/fabric/common/tools/configtxlator

#下载chaintool脚本
curl -L https://github.com/hyperledger/fabric-chaintool/releases/download/v0.10.3/chaintool > /usr/local/bin/chaintool


#修改docker配置文件
curl -fsSL https://get.docker.com | sh
安装完成之后
修改docker配置文件，如果没有配置文件，修改
vim /lib/systemd/system/docker.service
==================
ExecStart=/usr/bin/dockerd -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock -H fd:// --api-cors-header="*" --default-ulimit=nofile=8192:16384 --default-ulimit=nproc=8192:16384

然后
sudo systemctl daemon-reload
sudo systemctl restart docker (sudo service docker restart)

#google被墙，如何解决golang tools 没法下载的问题
git clone https://github.com/golang/tools.git $GOPATH/src/golang.org/x/tools

go get github.com/golang/protobuf/protoc-gen-go \
&& go get github.com/kardianos/govendor \
&& go get github.com/golang/lint/golint \
&& go get golang.org/x/tools/cmd/goimports \
&& go get github.com/onsi/ginkgo/ginkgo \
&& go get github.com/axw/gocov \
&& go get github.com/client9/misspell/cmd/misspell \
&& go get github.com/AlekSi/gocov-xml

#从源码生成镜像
cd $GOPATH/src/github.com/hyperledger/fabric
make docker


#启动fabric网络
#准备相关网络配置文件

1、生成组织关系和身份证书
cryptogen generate --config=./crypto-config.yaml --output ./crypto-config

2、生成Ordering服务启动初始区块
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./orderer.genesis.block

3、生成新建应用通道的配置交易
CHANNEL_NAME=businesschannel
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx  ./businesschannel.tx -channelID $CHANNEL_NAME

4、生成锚节点配置更新文件
configtxgen \
-profile TwoOrgsChannel \
-outputAnchorPeersUpdate ./Org1MSPanchors.tx  \
-channelID $CHANNEL_NAME \
-asOrg Org1MSP

configtxgen \
-profile TwoOrgsChannel \
-outputAnchorPeersUpdate ./Org2MSPanchors.tx \
-channelID $CHANNEL_NAME \
-asOrg Org2MSP
======================================================================================
#启动ORDER节点
#需配置的环境环境变量
ORDERER_GENERAL_LOGLEVEL=INFO
ORDERER_GENERAL_LISTENADDERSS=192.168.56.101
ORDERER_GENERAL_LISTENPORT=7050
ORDERER_GENERAL_GENESISMETHOD=file
ORDERER_GENERAL_GENESISFILE=/etc/hyperledger/fabric/orderer.genesis.block
ORDERER_GENERAL_LOCALMSPID=OrdererMSPT1
ORDERER_GENERAL_LOCALMSPDIR=/etc/hyperledger/fabric/msp
ORDERER_GENERAL_LEDGERTYPE=file
ORDERER_GENERAL_BATCHTIMEOUT=10s
ORDERER_GENERAL_MAXMESSAGECOUNT=10
ORDERER_GENERAL_TLS_ENABLED=false true
ORDERER_GENERAL_TLS_PRIVATEKEY=/etc/hyperledger/fabric/tls/server.key
ORDERER_GENERAL_TLS_CERTIFICATE=/etc/hyperledger/fabric/tls/server.crt
ORDERER_GENERAL_TLS_ROOTCAS=[/etc/hyperledger/fabric/tls/ca.crt]

#指定fabric配置文件路劲非常重要
export FABRIC_CFG_PATH=/etc/hyperledger/fabric
#然后就可以开启排序节点了！
orderer start
======================================================================================
#启动peer节点
CORE_LOGGING_LEVEL=INFO
CORE_PEER_ID=p0.g1.c
CORE_PEER_ADDRESS=192.168.56.102:7051
CORE_PEER_GOSSIP_EXTERNALENDPOINT=192.168.56.102:7051
CORE_PEER_GOSSIP_USELEADERELECTION=true
CORE_PEER_GOSSIP_ORGLEADER=false
CORE_PEER_LOCALMSPID=Org1MSPT1X
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/fabric/msp
CORE_VM_ENDPOINT=unix:///var/run/docker.sock
CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=host
CORE_PEER_TLS_ENABLED=false true
CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/fabric/tls/server.crt
CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/fabric/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/fabric/tls/ca.crt

#启动节点
peer node start


#节点启动之后，在cli（客户端）中操作
#创建通道
CHANNEL_NAME=businesschannel
CORE_PEER_LOCALMSPID="Org1MSP" \
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/fabric/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp \
peer channel create \
-o orderer.example.com:7050 \
-c ${CHANNEL_NAME} \
-f ./businesschannel.tx \
--tls \
--cafile /etc/hyperledger/fabric/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com//msp/tlscacerts/tlsca.example.com-cert.pem

#fabric网络启动之后区块链账本信息存放在每个peer容器的
/var/hyperledger/production

=============================================================================================================
区块链应用开发

每个链码都需要实现 Chaincode 接口
type Chaincode interface{
    Init(stub ChaincodeStubInterface) pb.Response
    Invoke(stub ChaincodeStubInterface) pb.Response
}

Init:当链码收到实例化或者升级类型的交易时，Init方法会被调用。
Invoke：当链码收到调用invoke或者查询query类型的交易，Invoke方法会被调用。

链码结构：

package main
import (
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)
type SimpleChaincode struct{}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    //该方法中实现链码初始化或升级的处理逻辑
}

func (t *SimpleChaincode) Invoke (stub shim.ChaincodeStubInterface) pb.Response{
    //在该方法中实现链码运行照片那个被调用活查询时的处理逻辑
}

func main(){
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode :%s",err)
    }
}

其中shim包提供了链码与账本交互的中间层。链码通过shim.ChaincodeStub 提供的方法来读取和修改账本状态。
pb 提供了Init和Invoke方法需要返回的pb.Response类型。

节点和链码容器之间通过gRPC消息来交互。两者之间采用ChaincodeMessage消息，基本结构如下：

message ChaincodeMessage{
    enum Type{
        UNDEFINED = 0;
        REGISTER = 1;
        REGISTERED = 2;
        INIT = 3;
        READY = 4;
        TRANSACTION = 5;
        COMPLETED = 6 ;
        ERROR = 7;
        GET_STATE = 8;
        PUT_STATE = 9;
        DEL_STATE = 10;
        INVOKE_CHAINCODE = 11;
        RESPONSE = 13;
        GET_STATE_BY_RANGE = 14;
        GET_QUERY_RESULT = 15;
        QUERY_STATE_NEXT = 16;
        QUERY_STATE_CLOSE = 17;
        KEEPALIVE = 18;
        GET_HISTORY_FOR_KEY = 19;
    }
    Type type = 1;
    google.protobuf.Timestamp timestamp = 2
    bytes payload = 3;
    string txid = 4;
    SignedProposal  proposal = 5;
    ChaincodeEvent chaincode_event = 6;
}

链码容器的shim层是节点与链码交互的中间层，当链码的代码逻辑要读写账本的时候，链码会通过shim层发送响应操作类型的ChaincodeMessage给节点，节点本地操作账本后返回响应消息。客户端收到足够的响应消息，并且有足够的背书节点支持后，就会将这笔交易发送给排序节点进行排序，并最终写入区块链。