package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstancePayTypeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               string `position:"Query" name:"Period"`
	AgentId              string `position:"Query" name:"AgentId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Resource             string `position:"Query" name:"Resource"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	UsedTime             int    `position:"Query" name:"UsedTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	PayType              string `position:"Query" name:"PayType"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
}

func (req *ModifyDBInstancePayTypeRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstancePayTypeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstancePayType", "rds", "")
	resp = &ModifyDBInstancePayTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstancePayTypeResponse struct {
	responses.BaseResponse
	RequestId    string
	DBInstanceId string
	OrderId      int64
}
