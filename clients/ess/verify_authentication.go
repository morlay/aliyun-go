package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type VerifyAuthenticationRequest struct {
	requests.RpcRequest
	Uid                  int64  `position:"Query" name:"Uid"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *VerifyAuthenticationRequest) Invoke(client *sdk.Client) (resp *VerifyAuthenticationResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "VerifyAuthentication", "ess", "")
	resp = &VerifyAuthenticationResponse{}
	err = client.DoAction(req, resp)
	return
}

type VerifyAuthenticationResponse struct {
	responses.BaseResponse
	RequestId common.String
}
