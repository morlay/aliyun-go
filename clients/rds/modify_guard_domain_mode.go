package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyGuardDomainModeRequest struct {
	requests.RpcRequest
	DomainMode           string `position:"Query" name:"DomainMode"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyGuardDomainModeRequest) Invoke(client *sdk.Client) (resp *ModifyGuardDomainModeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyGuardDomainMode", "rds", "")
	resp = &ModifyGuardDomainModeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyGuardDomainModeResponse struct {
	responses.BaseResponse
	RequestId string
}
