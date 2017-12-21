package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySecurityIpsRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	SecurityIps          string `position:"Query" name:"SecurityIps"`
}

func (req *ModifySecurityIpsRequest) Invoke(client *sdk.Client) (resp *ModifySecurityIpsResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "ModifySecurityIps", "", "")
	resp = &ModifySecurityIpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySecurityIpsResponse struct {
	responses.BaseResponse
	RequestId string
}
