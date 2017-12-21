package sas_api

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      string
	Message   string
	Success   bool
	RequestId string
	Data      int
}
