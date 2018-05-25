package dyvmsapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type VoipAddAccountRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DeviceId             string `position:"Query" name:"DeviceId"`
}

func (req *VoipAddAccountRequest) Invoke(client *sdk.Client) (resp *VoipAddAccountResponse, err error) {
	req.InitWithApiInfo("Dyvmsapi", "2017-05-25", "VoipAddAccount", "", "")
	resp = &VoipAddAccountResponse{}
	err = client.DoAction(req, resp)
	return
}

type VoipAddAccountResponse struct {
	responses.BaseResponse
	RequestId common.String
	Module    common.String
	Code      common.String
	Message   common.String
}
