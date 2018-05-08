package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeEventDetailRequest struct {
	requests.RpcRequest
	EventId              string `position:"Query" name:"EventId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeEventDetailRequest) Invoke(client *sdk.Client) (resp *DescribeEventDetailResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeEventDetail", "ecs", "")
	resp = &DescribeEventDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeEventDetailResponse struct {
	responses.BaseResponse
	RequestId     common.String
	ResourceId    common.String
	EventType     common.String
	EventCategory common.String
	Status        common.String
	SupportModify common.String
	PlanTime      common.String
	ExpireTime    common.String
	EventId       common.String
	StartTime     common.String
	EndTime       common.String
	EffectTime    common.String
	LimitTime     common.String
	Mark          common.String
}
