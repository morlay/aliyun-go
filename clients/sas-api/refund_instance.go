package sas_api

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RefundInstanceRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *RefundInstanceRequest) Invoke(client *sdk.Client) (resp *RefundInstanceResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "RefundInstance", "", "")
	resp = &RefundInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RefundInstanceResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
}
