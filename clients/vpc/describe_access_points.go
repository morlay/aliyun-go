package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAccessPointsRequest struct {
	requests.RpcRequest
	Filters              *DescribeAccessPointsFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                           `position:"Query" name:"ResourceOwnerId"`
	HostOperator         string                          `position:"Query" name:"HostOperator"`
	ResourceOwnerAccount string                          `position:"Query" name:"ResourceOwnerAccount"`
	Name                 string                          `position:"Query" name:"Name"`
	PageSize             int                             `position:"Query" name:"PageSize"`
	OwnerId              int64                           `position:"Query" name:"OwnerId"`
	Type                 string                          `position:"Query" name:"Type"`
	PageNumber           int                             `position:"Query" name:"PageNumber"`
}

func (req *DescribeAccessPointsRequest) Invoke(client *sdk.Client) (resp *DescribeAccessPointsResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeAccessPoints", "vpc", "")
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
	RequestId      common.String
	PageNumber     common.Integer
	PageSize       common.Integer
	TotalCount     common.Integer
	AccessPointSet DescribeAccessPointsAccessPointTypeList
}

type DescribeAccessPointsAccessPointType struct {
	AccessPointId    common.String
	Status           common.String
	Type             common.String
	AttachedRegionNo common.String
	Location         common.String
	HostOperator     common.String
	Name             common.String
	Description      common.String
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
