package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInstanceHistoryEventsRequest struct {
	requests.RpcRequest
	EventIds                  *DescribeInstanceHistoryEventsEventIdList                  `position:"Query" type:"Repeated" name:"EventId"`
	ResourceOwnerId           int64                                                      `position:"Query" name:"ResourceOwnerId"`
	EventCycleStatus          string                                                     `position:"Query" name:"EventCycleStatus"`
	PageNumber                int                                                        `position:"Query" name:"PageNumber"`
	PageSize                  int                                                        `position:"Query" name:"PageSize"`
	InstanceEventCycleStatuss *DescribeInstanceHistoryEventsInstanceEventCycleStatusList `position:"Query" type:"Repeated" name:"InstanceEventCycleStatus"`
	EventPublishTimeEnd       string                                                     `position:"Query" name:"EventPublishTimeEnd"`
	InstanceEventTypes        *DescribeInstanceHistoryEventsInstanceEventTypeList        `position:"Query" type:"Repeated" name:"InstanceEventType"`
	ResourceOwnerAccount      string                                                     `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string                                                     `position:"Query" name:"OwnerAccount"`
	NotBeforeStart            string                                                     `position:"Query" name:"NotBeforeStart"`
	OwnerId                   int64                                                      `position:"Query" name:"OwnerId"`
	EventPublishTimeStart     string                                                     `position:"Query" name:"EventPublishTimeStart"`
	InstanceId                string                                                     `position:"Query" name:"InstanceId"`
	NotBeforeEnd              string                                                     `position:"Query" name:"NotBeforeEnd"`
	EventType                 string                                                     `position:"Query" name:"EventType"`
}

func (req *DescribeInstanceHistoryEventsRequest) Invoke(client *sdk.Client) (resp *DescribeInstanceHistoryEventsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceHistoryEvents", "ecs", "")
	resp = &DescribeInstanceHistoryEventsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstanceHistoryEventsResponse struct {
	responses.BaseResponse
	RequestId              common.String
	TotalCount             common.Integer
	PageNumber             common.Integer
	PageSize               common.Integer
	InstanceSystemEventSet DescribeInstanceHistoryEventsInstanceSystemEventTypeList
}

type DescribeInstanceHistoryEventsInstanceSystemEventType struct {
	InstanceId       common.String
	EventId          common.String
	EventPublishTime common.String
	NotBefore        common.String
	EventFinishTime  common.String
	EventType        DescribeInstanceHistoryEventsEventType
	EventCycleStatus DescribeInstanceHistoryEventsEventCycleStatus
}

type DescribeInstanceHistoryEventsEventType struct {
	Code common.Integer
	Name common.String
}

type DescribeInstanceHistoryEventsEventCycleStatus struct {
	Code common.Integer
	Name common.String
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

type DescribeInstanceHistoryEventsInstanceEventCycleStatusList []string

func (list *DescribeInstanceHistoryEventsInstanceEventCycleStatusList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceHistoryEventsInstanceEventTypeList []string

func (list *DescribeInstanceHistoryEventsInstanceEventTypeList) UnmarshalJSON(data []byte) error {
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
