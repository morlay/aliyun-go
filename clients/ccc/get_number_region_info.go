package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetNumberRegionInfoRequest struct {
	requests.RpcRequest
	Number     string `position:"Query" name:"Number"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *GetNumberRegionInfoRequest) Invoke(client *sdk.Client) (resp *GetNumberRegionInfoResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "GetNumberRegionInfo", "ccc", "")
	resp = &GetNumberRegionInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetNumberRegionInfoResponse struct {
	responses.BaseResponse
	RequestId   string
	Success     bool
	Code        string
	Message     string
	PhoneNumber GetNumberRegionInfoPhoneNumber
}

type GetNumberRegionInfoPhoneNumber struct {
	Number   string
	Province string
	City     string
}
