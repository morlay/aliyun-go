package ocs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeMonitorValuesRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	MonitorKeys          string `position:"Query" name:"MonitorKeys"`
}

func (req *DescribeMonitorValuesRequest) Invoke(client *sdk.Client) (resp *DescribeMonitorValuesResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "DescribeMonitorValues", "", "")
	resp = &DescribeMonitorValuesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeMonitorValuesResponse struct {
	responses.BaseResponse
	RequestId                   string
	GetOcsMonitorValuesResponse DescribeMonitorValuesGetOcsMonitorValuesResponse
}

type DescribeMonitorValuesGetOcsMonitorValuesResponse struct {
	Date           string
	OcsInstanceIds DescribeMonitorValuesOcsMonitorDTOList
}

type DescribeMonitorValuesOcsMonitorDTO struct {
	OcsInstanceId string
	MonitorKeys   DescribeMonitorValuesOcsMonitorKeyDTOList
}

type DescribeMonitorValuesOcsMonitorKeyDTO struct {
	MonitorKey string
	Value      string
	Unit       string
	Date       string
}

type DescribeMonitorValuesOcsMonitorDTOList []DescribeMonitorValuesOcsMonitorDTO

func (list *DescribeMonitorValuesOcsMonitorDTOList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorValuesOcsMonitorDTO)
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

type DescribeMonitorValuesOcsMonitorKeyDTOList []DescribeMonitorValuesOcsMonitorKeyDTO

func (list *DescribeMonitorValuesOcsMonitorKeyDTOList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorValuesOcsMonitorKeyDTO)
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
