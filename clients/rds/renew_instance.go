package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RenewInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               string `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
}

func (req *RenewInstanceRequest) Invoke(client *sdk.Client) (resp *RenewInstanceResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "RenewInstance", "rds", "")
	resp = &RenewInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
	OrderId   common.Long
}
