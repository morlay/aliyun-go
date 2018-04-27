package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetInstanceConsoleOutputRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
}

func (req *GetInstanceConsoleOutputRequest) Invoke(client *sdk.Client) (resp *GetInstanceConsoleOutputResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "GetInstanceConsoleOutput", "ecs", "")
	resp = &GetInstanceConsoleOutputResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetInstanceConsoleOutputResponse struct {
	responses.BaseResponse
	RequestId      string
	InstanceId     string
	ConsoleOutput  string
	LastUpdateTime string
}
