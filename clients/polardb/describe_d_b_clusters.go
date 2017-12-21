package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBClustersRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBType               string `position:"Query" name:"DBType"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeDBClustersRequest) Invoke(client *sdk.Client) (resp *DescribeDBClustersResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusters", "polardb", "")
	resp = &DescribeDBClustersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBClustersResponse struct {
	responses.BaseResponse
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeDBClustersDBClusterList
}

type DescribeDBClustersDBCluster struct {
	DBClusterId          string
	DBClusterDescription string
	PayType              string
	DBClusterNetworkType string
	RegionId             string
	ExpireTime           string
	DBClusterStatus      string
	Engine               string
	DBType               string
	DBVersion            string
	LockMode             string
	LockReason           string
	CreateTime           string
	VpcId                string
	VSwitchId            string
}

type DescribeDBClustersDBClusterList []DescribeDBClustersDBCluster

func (list *DescribeDBClustersDBClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClustersDBCluster)
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
