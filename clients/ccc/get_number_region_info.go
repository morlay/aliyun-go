package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	Success     bool
	Code        common.String
	Message     common.String
	PhoneNumber GetNumberRegionInfoPhoneNumber
}

type GetNumberRegionInfoPhoneNumber struct {
	Number   common.String
	Province common.String
	City     common.String
}
