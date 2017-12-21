package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateLoginProfileRequest struct {
	Password              string `position:"Query" name:"Password"`
	PasswordResetRequired string `position:"Query" name:"PasswordResetRequired"`
	MFABindRequired       string `position:"Query" name:"MFABindRequired"`
	UserName              string `position:"Query" name:"UserName"`
}

func (r CreateLoginProfileRequest) Invoke(client *sdk.Client) (response *CreateLoginProfileResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateLoginProfileRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateLoginProfile", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateLoginProfileResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateLoginProfileResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateLoginProfileResponse struct {
	RequestId    string
	LoginProfile CreateLoginProfileLoginProfile
}

type CreateLoginProfileLoginProfile struct {
	UserName              string
	PasswordResetRequired bool
	MFABindRequired       bool
	CreateDate            string
}
