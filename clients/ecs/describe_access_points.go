package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAccessPointsRequest struct {
	requests.RpcRequest
	Filters              *DescribeAccessPointsFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                           `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                          `position:"Query" name:"ResourceOwnerAccount"`
	PageSize             int                             `position:"Query" name:"PageSize"`
	OwnerId              int64                           `position:"Query" name:"OwnerId"`
	Type                 string                          `position:"Query" name:"Type"`
	PageNumber           int                             `position:"Query" name:"PageNumber"`
}

func (req *DescribeAccessPointsRequest) Invoke(client *sdk.Client) (resp *DescribeAccessPointsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeAccessPoints", "ecs", "")
	resp = &DescribeAccessPointsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAccessPointsFilter struct {
	Key    string                         `name:"Key"`
	Values *DescribeAccessPointsValueList `type:"Repeated" name:"Value"`
}

type DescribeAccessPointsResponse struct {
	responses.BaseResponse
	RequestId      string
	PageNumber     int
	PageSize       int
	TotalCount     int
	AccessPointSet DescribeAccessPointsAccessPointTypeList
}

type DescribeAccessPointsAccessPointType struct {
	AccessPointId    string
	Status           string
	Type             string
	AttachedRegionNo string
	Location         string
	HostOperator     string
	Name             string
	Description      string
}

type DescribeAccessPointsFilterList []DescribeAccessPointsFilter

func (list *DescribeAccessPointsFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessPointsFilter)
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

type DescribeAccessPointsValueList []string

func (list *DescribeAccessPointsValueList) UnmarshalJSON(data []byte) error {
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

type DescribeAccessPointsAccessPointTypeList []DescribeAccessPointsAccessPointType

func (list *DescribeAccessPointsAccessPointTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessPointsAccessPointType)
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
