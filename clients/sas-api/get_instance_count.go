package sas_api

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetInstanceCountRequest struct {
	requests.RpcRequest
	OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (req *GetInstanceCountRequest) Invoke(client *sdk.Client) (resp *GetInstanceCountResponse, err error) {
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetInstanceCount", "", "")
	resp = &GetInstanceCountResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetInstanceCountResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	Success   bool
	RequestId common.String
	Data      common.Integer
}
