package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetUserMFAInfoRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

func (req *GetUserMFAInfoRequest) Invoke(client *sdk.Client) (resp *GetUserMFAInfoResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "GetUserMFAInfo", "", "")
	resp = &GetUserMFAInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserMFAInfoResponse struct {
	responses.BaseResponse
	RequestId common.String
	MFADevice GetUserMFAInfoMFADevice
}

type GetUserMFAInfoMFADevice struct {
	SerialNumber common.String
}
