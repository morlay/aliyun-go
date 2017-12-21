package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateLoginProfileRequest struct {
	Password              string `position:"Query" name:"Password"`
	PasswordResetRequired string `position:"Query" name:"PasswordResetRequired"`
	MFABindRequired       string `position:"Query" name:"MFABindRequired"`
	UserName              string `position:"Query" name:"UserName"`
}

func (r UpdateLoginProfileRequest) Invoke(client *sdk.Client) (response *UpdateLoginProfileResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateLoginProfileRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdateLoginProfile", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateLoginProfileResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateLoginProfileResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateLoginProfileResponse struct {
	RequestId string
}
