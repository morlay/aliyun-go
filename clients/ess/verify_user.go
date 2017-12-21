package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VerifyUserRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r VerifyUserRequest) Invoke(client *sdk.Client) (response *VerifyUserResponse, err error) {
	req := struct {
		*requests.RpcRequest
		VerifyUserRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "VerifyUser", "ess", "")

	resp := struct {
		*responses.BaseResponse
		VerifyUserResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.VerifyUserResponse

	err = client.DoAction(&req, &resp)
	return
}

type VerifyUserResponse struct {
}
