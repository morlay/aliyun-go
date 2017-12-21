package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDBInstanceReplicaRequest struct {
	ConnectionMode        string `position:"Query" name:"ConnectionMode"`
	ReplicaDescription    string `position:"Query" name:"ReplicaDescription"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     int    `position:"Query" name:"DBInstanceStorage"`
	SystemDBCharset       string `position:"Query" name:"SystemDBCharset"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	EngineVersion         string `position:"Query" name:"EngineVersion"`
	Engine                string `position:"Query" name:"Engine"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	DBInstanceNetType     string `position:"Query" name:"DBInstanceNetType"`
	Period                string `position:"Query" name:"Period"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	UsedTime              string `position:"Query" name:"UsedTime"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	SecurityIPList        string `position:"Query" name:"SecurityIPList"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string `position:"Query" name:"PrivateIpAddress"`
	SourceDBInstanceId    string `position:"Query" name:"SourceDBInstanceId"`
	VPCId                 string `position:"Query" name:"VPCId"`
	ZoneId                string `position:"Query" name:"ZoneId"`
	PayType               string `position:"Query" name:"PayType"`
	InstanceNetworkType   string `position:"Query" name:"InstanceNetworkType"`
}

func (r CreateDBInstanceReplicaRequest) Invoke(client *sdk.Client) (response *CreateDBInstanceReplicaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateDBInstanceReplicaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateDBInstanceReplica", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CreateDBInstanceReplicaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateDBInstanceReplicaResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateDBInstanceReplicaResponse struct {
	RequestId    string
	DBInstanceId string
	OrderId      int64
	ReplicaId    string
	WorkflowId   string
}
