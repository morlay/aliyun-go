package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SwitchDBInstanceHARequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	EffectiveTime        string `position:"Query" name:"EffectiveTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeId               string `position:"Query" name:"NodeId"`
	Operation            string `position:"Query" name:"Operation"`
}

func (req *SwitchDBInstanceHARequest) Invoke(client *sdk.Client) (resp *SwitchDBInstanceHAResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "SwitchDBInstanceHA", "rds", "")
	resp = &SwitchDBInstanceHAResponse{}
	err = client.DoAction(req, resp)
	return
}

type SwitchDBInstanceHAResponse struct {
	responses.BaseResponse
	RequestId string
}
