package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type BatchRegisterDeviceWithApplyIdRequest struct {
	requests.RpcRequest
	ApplyId    int64  `position:"Query" name:"ApplyId"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *BatchRegisterDeviceWithApplyIdRequest) Invoke(client *sdk.Client) (resp *BatchRegisterDeviceWithApplyIdResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "BatchRegisterDeviceWithApplyId", "", "")
	resp = &BatchRegisterDeviceWithApplyIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type BatchRegisterDeviceWithApplyIdResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         BatchRegisterDeviceWithApplyIdData
}

type BatchRegisterDeviceWithApplyIdData struct {
	ApplyId common.Long
}
