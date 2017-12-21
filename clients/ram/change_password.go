package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ChangePasswordRequest struct {
	OldPassword string `position:"Query" name:"OldPassword"`
	NewPassword string `position:"Query" name:"NewPassword"`
}

func (r ChangePasswordRequest) Invoke(client *sdk.Client) (response *ChangePasswordResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ChangePasswordRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ChangePassword", "", "")

	resp := struct {
		*responses.BaseResponse
		ChangePasswordResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ChangePasswordResponse

	err = client.DoAction(&req, &resp)
	return
}

type ChangePasswordResponse struct {
	RequestId string
}
