package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeUserDataRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeUserDataRequest) Invoke(client *sdk.Client) (resp *DescribeUserDataResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeUserData", "ecs", "")
	resp = &DescribeUserDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeUserDataResponse struct {
	responses.BaseResponse
	RequestId  common.String
	RegionId   common.String
	InstanceId common.String
	UserData   common.String
}
