package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartDBInstanceDiagnoseRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ProxyId              string `position:"Query" name:"ProxyId"`
}

func (req *StartDBInstanceDiagnoseRequest) Invoke(client *sdk.Client) (resp *StartDBInstanceDiagnoseResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "StartDBInstanceDiagnose", "rds", "")
	resp = &StartDBInstanceDiagnoseResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartDBInstanceDiagnoseResponse struct {
	responses.BaseResponse
	RequestId      string
	DBInstanceName string
	DBInstanceId   string
}
