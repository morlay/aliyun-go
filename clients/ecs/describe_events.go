package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeEventsRequest struct {
	requests.RpcRequest
	EventId              string `position:"Query" name:"EventId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	PageSize             int    `position:"Query" name:"PageSize"`
	PlanTime             string `position:"Query" name:"PlanTime"`
	ExpireTime           string `position:"Query" name:"ExpireTime"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	EventType            string `position:"Query" name:"EventType"`
	Status               string `position:"Query" name:"Status"`
}

func (req *DescribeEventsRequest) Invoke(client *sdk.Client) (resp *DescribeEventsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeEvents", "ecs", "")
	resp = &DescribeEventsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeEventsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageNumber common.Integer
	PageSize   common.Integer
	TotalCount common.Integer
	Events     DescribeEventsEventList
}

type DescribeEventsEvent struct {
	ResourceId    common.String
	EventType     common.String
	EventCategory common.String
	Status        common.String
	SupportModify common.String
	PlanTime      common.String
	ExpireTime    common.String
	EventId       common.String
}

type DescribeEventsEventList []DescribeEventsEvent

func (list *DescribeEventsEventList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEventsEvent)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
