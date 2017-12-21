package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeHpcClustersRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	HpcClusterIds        string `position:"Query" name:"HpcClusterIds"`
}

func (r DescribeHpcClustersRequest) Invoke(client *sdk.Client) (response *DescribeHpcClustersResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeHpcClustersRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeHpcClusters", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeHpcClustersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeHpcClustersResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeHpcClustersResponse struct {
	RequestId   string
	TotalCount  int
	PageNumber  int
	PageSize    int
	HpcClusters DescribeHpcClustersHpcClusterList
}

type DescribeHpcClustersHpcCluster struct {
	HpcClusterId string
	Name         string
	Description  string
}

type DescribeHpcClustersHpcClusterList []DescribeHpcClustersHpcCluster

func (list *DescribeHpcClustersHpcClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHpcClustersHpcCluster)
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
