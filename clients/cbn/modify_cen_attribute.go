package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCenAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyCenAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyCenAttributeResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "ModifyCenAttribute", "cbn", "")
	resp = &ModifyCenAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCenAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
