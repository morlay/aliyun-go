package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVRoutersRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	VRouterId            string `position:"Query" name:"VRouterId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeVRoutersRequest) Invoke(client *sdk.Client) (response *DescribeVRoutersResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeVRoutersRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeVRouters", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeVRoutersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeVRoutersResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeVRoutersResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	VRouters   DescribeVRoutersVRouterList
}

type DescribeVRoutersVRouter struct {
	RegionId      string
	VpcId         string
	VRouterName   string
	Description   string
	VRouterId     string
	CreationTime  string
	RouteTableIds DescribeVRoutersRouteTableIdList
}

type DescribeVRoutersVRouterList []DescribeVRoutersVRouter

func (list *DescribeVRoutersVRouterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVRoutersVRouter)
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

type DescribeVRoutersRouteTableIdList []string

func (list *DescribeVRoutersRouteTableIdList) UnmarshalJSON(data []byte) error {
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
