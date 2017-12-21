package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetLoginProfileRequest struct {
	UserName string `position:"Query" name:"UserName"`
}

func (r GetLoginProfileRequest) Invoke(client *sdk.Client) (response *GetLoginProfileResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetLoginProfileRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "GetLoginProfile", "", "")

	resp := struct {
		*responses.BaseResponse
		GetLoginProfileResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetLoginProfileResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetLoginProfileResponse struct {
	RequestId    string
	LoginProfile GetLoginProfileLoginProfile
}

type GetLoginProfileLoginProfile struct {
	UserName              string
	PasswordResetRequired bool
	MFABindRequired       bool
	CreateDate            string
}
