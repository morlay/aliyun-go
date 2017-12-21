package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetLoginProfileRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

func (req *GetLoginProfileRequest) Invoke(client *sdk.Client) (resp *GetLoginProfileResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "GetLoginProfile", "", "")
	resp = &GetLoginProfileResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetLoginProfileResponse struct {
	responses.BaseResponse
	RequestId    string
	LoginProfile GetLoginProfileLoginProfile
}

type GetLoginProfileLoginProfile struct {
	UserName              string
	PasswordResetRequired bool
	MFABindRequired       bool
	CreateDate            string
}
