package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateLoginProfileRequest struct {
	requests.RpcRequest
	Password              string `position:"Query" name:"Password"`
	PasswordResetRequired string `position:"Query" name:"PasswordResetRequired"`
	MFABindRequired       string `position:"Query" name:"MFABindRequired"`
	UserName              string `position:"Query" name:"UserName"`
}

func (req *UpdateLoginProfileRequest) Invoke(client *sdk.Client) (resp *UpdateLoginProfileResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdateLoginProfile", "", "")
	resp = &UpdateLoginProfileResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateLoginProfileResponse struct {
	responses.BaseResponse
	RequestId string
}
