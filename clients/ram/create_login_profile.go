package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateLoginProfileRequest struct {
	requests.RpcRequest
	Password              string `position:"Query" name:"Password"`
	PasswordResetRequired string `position:"Query" name:"PasswordResetRequired"`
	MFABindRequired       string `position:"Query" name:"MFABindRequired"`
	UserName              string `position:"Query" name:"UserName"`
}

func (req *CreateLoginProfileRequest) Invoke(client *sdk.Client) (resp *CreateLoginProfileResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateLoginProfile", "", "")
	resp = &CreateLoginProfileResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateLoginProfileResponse struct {
	responses.BaseResponse
	RequestId    string
	LoginProfile CreateLoginProfileLoginProfile
}

type CreateLoginProfileLoginProfile struct {
	UserName              string
	PasswordResetRequired bool
	MFABindRequired       bool
	CreateDate            string
}
