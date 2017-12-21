package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveApPortalConfigRequest struct {
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

func (r SaveApPortalConfigRequest) Invoke(client *sdk.Client) (response *SaveApPortalConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveApPortalConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApPortalConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveApPortalConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveApPortalConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveApPortalConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
