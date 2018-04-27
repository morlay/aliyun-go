package cbn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRouteConflictRequest struct {
	requests.RpcRequest
	ChildInstanceId           string `position:"Query" name:"ChildInstanceId"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	DestinationCidrBlock      string `position:"Query" name:"DestinationCidrBlock"`
	PageSize                  int    `position:"Query" name:"PageSize"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	ChildInstanceType         string `position:"Query" name:"ChildInstanceType"`
	ChildInstanceRouteTableId string `position:"Query" name:"ChildInstanceRouteTableId"`
	PageNumber                int    `position:"Query" name:"PageNumber"`
	ChildInstanceRegionId     string `position:"Query" name:"ChildInstanceRegionId"`
}

func (req *DescribeRouteConflictRequest) Invoke(client *sdk.Client) (resp *DescribeRouteConflictResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DescribeRouteConflict", "cbn", "")
	resp = &DescribeRouteConflictResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRouteConflictResponse struct {
	responses.BaseResponse
	RequestId      string
	PageNumber     int
	TotalCount     int
	PageSize       int
	RouteConflicts DescribeRouteConflictRouteConflictList
}

type DescribeRouteConflictRouteConflict struct {
	DestinationCidrBlock string
	RegionId             string
	InstanceId           string
	InstanceType         string
	Status               string
}

type DescribeRouteConflictRouteConflictList []DescribeRouteConflictRouteConflict

func (list *DescribeRouteConflictRouteConflictList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteConflictRouteConflict)
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
