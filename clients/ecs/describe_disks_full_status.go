package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDisksFullStatusRequest struct {
	requests.RpcRequest
	EventIds             *DescribeDisksFullStatusEventIdList `position:"Query" type:"Repeated" name:"EventId"`
	ResourceOwnerId      int64                               `position:"Query" name:"ResourceOwnerId"`
	PageNumber           int                                 `position:"Query" name:"PageNumber"`
	EventTimeStart       string                              `position:"Query" name:"EventTimeStart"`
	PageSize             int                                 `position:"Query" name:"PageSize"`
	DiskIds              *DescribeDisksFullStatusDiskIdList  `position:"Query" type:"Repeated" name:"DiskId"`
	ResourceOwnerAccount string                              `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                              `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                               `position:"Query" name:"OwnerId"`
	EventTimeEnd         string                              `position:"Query" name:"EventTimeEnd"`
	HealthStatus         string                              `position:"Query" name:"HealthStatus"`
	EventType            string                              `position:"Query" name:"EventType"`
	Status               string                              `position:"Query" name:"Status"`
}

func (req *DescribeDisksFullStatusRequest) Invoke(client *sdk.Client) (resp *DescribeDisksFullStatusResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDisksFullStatus", "ecs", "")
	resp = &DescribeDisksFullStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDisksFullStatusResponse struct {
	responses.BaseResponse
	RequestId         common.String
	TotalCount        common.Integer
	PageNumber        common.Integer
	PageSize          common.Integer
	DiskFullStatusSet DescribeDisksFullStatusDiskFullStatusTypeList
}

type DescribeDisksFullStatusDiskFullStatusType struct {
	DiskId       common.String
	DiskEventSet DescribeDisksFullStatusDiskEventTypeList
	Status       DescribeDisksFullStatusStatus
	HealthStatus DescribeDisksFullStatusHealthStatus
}

type DescribeDisksFullStatusDiskEventType struct {
	EventId   common.String
	EventTime common.String
	EventType DescribeDisksFullStatusEventType
}

type DescribeDisksFullStatusEventType struct {
	Code common.Integer
	Name common.String
}

type DescribeDisksFullStatusStatus struct {
	Code common.Integer
	Name common.String
}

type DescribeDisksFullStatusHealthStatus struct {
	Code common.Integer
	Name common.String
}

type DescribeDisksFullStatusEventIdList []string

func (list *DescribeDisksFullStatusEventIdList) UnmarshalJSON(data []byte) error {
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

type DescribeDisksFullStatusDiskIdList []string

func (list *DescribeDisksFullStatusDiskIdList) UnmarshalJSON(data []byte) error {
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

type DescribeDisksFullStatusDiskFullStatusTypeList []DescribeDisksFullStatusDiskFullStatusType

func (list *DescribeDisksFullStatusDiskFullStatusTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksFullStatusDiskFullStatusType)
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

type DescribeDisksFullStatusDiskEventTypeList []DescribeDisksFullStatusDiskEventType

func (list *DescribeDisksFullStatusDiskEventTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksFullStatusDiskEventType)
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
