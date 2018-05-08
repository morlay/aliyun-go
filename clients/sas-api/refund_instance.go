package sas_api

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.String
	Message   common.String
	RequestId common.String
}
