package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	LoginProfile GetLoginProfileLoginProfile
}

type GetLoginProfileLoginProfile struct {
	UserName              common.String
	PasswordResetRequired bool
	MFABindRequired       bool
	CreateDate            common.String
}
