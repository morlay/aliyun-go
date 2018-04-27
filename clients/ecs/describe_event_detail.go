package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId     string
	ResourceId    string
	EventType     string
	EventCategory string
	Status        string
	SupportModify string
	PlanTime      string
	ExpireTime    string
	EventId       string
	StartTime     string
	EndTime       string
	EffectTime    string
	LimitTime     string
	Mark          string
}
