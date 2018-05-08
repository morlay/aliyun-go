package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveApPortalConfigRequest struct {
	requests.RpcRequest
	AuthKey      string `position:"Query" name:"AuthKey"`
	PortalUrl    string `position:"Query" name:"PortalUrl"`
	PortalStatus string `position:"Query" name:"PortalStatus"`
	Whitelist    string `position:"Query" name:"Whitelist"`
	CheckUrl     string `position:"Query" name:"CheckUrl"`
	ApConfigId   int64  `position:"Query" name:"ApConfigId"`
	AuthSecret   string `position:"Query" name:"AuthSecret"`
	WebAuthUrl   string `position:"Query" name:"WebAuthUrl"`
	Network      int    `position:"Query" name:"Network"`
}

func (req *SaveApPortalConfigRequest) Invoke(client *sdk.Client) (resp *SaveApPortalConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApPortalConfig", "", "")
	resp = &SaveApPortalConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveApPortalConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
