package ots

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetInstanceRequest struct {
	requests.RpcRequest
	Accept      string `position:"Header" name:"Accept"`
	VersionName string `position:"Header" name:"VersionName"`
}

func (req *GetInstanceRequest) Invoke(client *sdk.Client) (resp *GetInstanceResponse, err error) {
	req.InitWithApiInfo("Ots", "2013-09-12", "GetInstance", "", "")
	resp = &GetInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
