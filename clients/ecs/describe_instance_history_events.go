package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceHistoryEventsRequest struct {
	EventIds              *DescribeInstanceHistoryEventsEventIdList `position:"Query" type:"Repeated" name:"EventId"`
	ResourceOwnerId       int64                                     `position:"Query" name:"ResourceOwnerId"`
	EventCycleStatus      string                                    `position:"Query" name:"EventCycleStatus"`
	PageNumber            int                                       `position:"Query" name:"PageNumber"`
	PageSize              int                                       `position:"Query" name:"PageSize"`
	EventPublishTimeEnd   string                                    `position:"Query" name:"EventPublishTimeEnd"`
	ResourceOwnerAccount  string                                    `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string                                    `position:"Query" name:"OwnerAccount"`
	NotBeforeStart        string                                    `position:"Query" name:"NotBeforeStart"`
	OwnerId               int64                                     `position:"Query" name:"OwnerId"`
	EventPublishTimeStart string                                    `position:"Query" name:"EventPublishTimeStart"`
	InstanceId            string                                    `position:"Query" name:"InstanceId"`
	NotBeforeEnd          string                                    `position:"Query" name:"NotBeforeEnd"`
	EventType             string                                    `position:"Query" name:"EventType"`
}

func (r DescribeInstanceHistoryEventsRequest) Invoke(client *sdk.Client) (response *DescribeInstanceHistoryEventsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceHistoryEventsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceHistoryEvents", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceHistoryEventsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstanceHistoryEventsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceHistoryEventsResponse struct {
	RequestId              string
	TotalCount             int
	PageNumber             int
	PageSize               int
	InstanceSystemEventSet DescribeInstanceHistoryEventsInstanceSystemEventTypeList
}

type DescribeInstanceHistoryEventsInstanceSystemEventType struct {
	InstanceId       string
	EventId          string
	EventPublishTime string
	NotBefore        string
	EventType        DescribeInstanceHistoryEventsEventType
}

type DescribeInstanceHistoryEventsEventType struct {
	Code int
	Name string
}

type DescribeInstanceHistoryEventsEventIdList []string

func (list *DescribeInstanceHistoryEventsEventIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeInstanceHistoryEventsInstanceSystemEventTypeList []DescribeInstanceHistoryEventsInstanceSystemEventType

func (list *DescribeInstanceHistoryEventsInstanceSystemEventTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceHistoryEventsInstanceSystemEventType)
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
