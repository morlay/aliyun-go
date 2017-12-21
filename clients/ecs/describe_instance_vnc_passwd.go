package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceVncPasswdRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeInstanceVncPasswdRequest) Invoke(client *sdk.Client) (resp *DescribeInstanceVncPasswdResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceVncPasswd", "ecs", "")
	resp = &DescribeInstanceVncPasswdResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstanceVncPasswdResponse struct {
	responses.BaseResponse
	RequestId string
	VncPasswd string
}
