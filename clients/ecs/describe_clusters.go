package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeClustersRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeClustersRequest) Invoke(client *sdk.Client) (resp *DescribeClustersResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeClusters", "ecs", "")
	resp = &DescribeClustersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClustersResponse struct {
	responses.BaseResponse
	RequestId common.String
	Clusters  DescribeClustersClusterList
}

type DescribeClustersCluster struct {
	ClusterId common.String
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
