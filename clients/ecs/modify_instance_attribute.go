package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyInstanceAttributeRequest struct {
	requests.RpcRequest
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Password             string `position:"Query" name:"Password"`
	HostName             string `position:"Query" name:"HostName"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	InstanceName         string `position:"Query" name:"InstanceName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Recyclable           string `position:"Query" name:"Recyclable"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyInstanceAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceAttribute", "ecs", "")
	resp = &ModifyInstanceAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
