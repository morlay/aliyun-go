package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ChangePasswordRequest struct {
	requests.RpcRequest
	OldPassword string `position:"Query" name:"OldPassword"`
	NewPassword string `position:"Query" name:"NewPassword"`
}

func (req *ChangePasswordRequest) Invoke(client *sdk.Client) (resp *ChangePasswordResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ChangePassword", "", "")
	resp = &ChangePasswordResponse{}
	err = client.DoAction(req, resp)
	return
}

type ChangePasswordResponse struct {
	responses.BaseResponse
	RequestId common.String
}
