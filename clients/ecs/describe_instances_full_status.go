package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstancesFullStatusRequest struct {
	requests.RpcRequest
	EventIds              *DescribeInstancesFullStatusEventIdList    `position:"Query" type:"Repeated" name:"EventId"`
	ResourceOwnerId       int64                                      `position:"Query" name:"ResourceOwnerId"`
	PageNumber            int                                        `position:"Query" name:"PageNumber"`
	PageSize              int                                        `position:"Query" name:"PageSize"`
	EventPublishTimeEnd   string                                     `position:"Query" name:"EventPublishTimeEnd"`
	ResourceOwnerAccount  string                                     `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string                                     `position:"Query" name:"OwnerAccount"`
	NotBeforeStart        string                                     `position:"Query" name:"NotBeforeStart"`
	OwnerId               int64                                      `position:"Query" name:"OwnerId"`
	EventPublishTimeStart string                                     `position:"Query" name:"EventPublishTimeStart"`
	InstanceIds           *DescribeInstancesFullStatusInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
	NotBeforeEnd          string                                     `position:"Query" name:"NotBeforeEnd"`
	HealthStatus          string                                     `position:"Query" name:"HealthStatus"`
	EventType             string                                     `position:"Query" name:"EventType"`
	Status                string                                     `position:"Query" name:"Status"`
}

func (req *DescribeInstancesFullStatusRequest) Invoke(client *sdk.Client) (resp *DescribeInstancesFullStatusResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstancesFullStatus", "ecs", "")
	resp = &DescribeInstancesFullStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstancesFullStatusResponse struct {
	responses.BaseResponse
	RequestId             string
	TotalCount            int
	PageNumber            int
	PageSize              int
	InstanceFullStatusSet DescribeInstancesFullStatusInstanceFullStatusTypeList
}

type DescribeInstancesFullStatusInstanceFullStatusType struct {
	InstanceId              string
	ScheduledSystemEventSet DescribeInstancesFullStatusScheduledSystemEventTypeList
	Status                  DescribeInstancesFullStatusStatus
	HealthStatus            DescribeInstancesFullStatusHealthStatus
}

type DescribeInstancesFullStatusScheduledSystemEventType struct {
	EventId          string
	EventPublishTime string
	NotBefore        string
	EventCycleStatus DescribeInstancesFullStatusEventCycleStatus
	EventType        DescribeInstancesFullStatusEventType
}

type DescribeInstancesFullStatusEventCycleStatus struct {
	Code int
	Name string
}

type DescribeInstancesFullStatusEventType struct {
	Code int
	Name string
}

type DescribeInstancesFullStatusStatus struct {
	Code int
	Name string
}

type DescribeInstancesFullStatusHealthStatus struct {
	Code int
	Name string
}

type DescribeInstancesFullStatusEventIdList []string

func (list *DescribeInstancesFullStatusEventIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesFullStatusInstanceIdList []string

func (list *DescribeInstancesFullStatusInstanceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesFullStatusInstanceFullStatusTypeList []DescribeInstancesFullStatusInstanceFullStatusType

func (list *DescribeInstancesFullStatusInstanceFullStatusTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesFullStatusInstanceFullStatusType)
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

type DescribeInstancesFullStatusScheduledSystemEventTypeList []DescribeInstancesFullStatusScheduledSystemEventType

func (list *DescribeInstancesFullStatusScheduledSystemEventTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesFullStatusScheduledSystemEventType)
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
