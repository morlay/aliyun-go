package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VerifyUserRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *VerifyUserRequest) Invoke(client *sdk.Client) (resp *VerifyUserResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "VerifyUser", "ess", "")
	resp = &VerifyUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type VerifyUserResponse struct {
	responses.BaseResponse
}
