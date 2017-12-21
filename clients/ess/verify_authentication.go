package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VerifyAuthenticationRequest struct {
	Uid                  int64  `position:"Query" name:"Uid"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r VerifyAuthenticationRequest) Invoke(client *sdk.Client) (response *VerifyAuthenticationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		VerifyAuthenticationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "VerifyAuthentication", "ess", "")

	resp := struct {
		*responses.BaseResponse
		VerifyAuthenticationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.VerifyAuthenticationResponse

	err = client.DoAction(&req, &resp)
	return
}

type VerifyAuthenticationResponse struct {
	RequestId string
}
