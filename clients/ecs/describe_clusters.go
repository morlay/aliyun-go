package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClustersRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeClustersRequest) Invoke(client *sdk.Client) (response *DescribeClustersResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeClustersRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeClusters", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeClustersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeClustersResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeClustersResponse struct {
	RequestId string
	Clusters  DescribeClustersClusterList
}

type DescribeClustersCluster struct {
	ClusterId string
}

type DescribeClustersClusterList []DescribeClustersCluster

func (list *DescribeClustersClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClustersCluster)
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
