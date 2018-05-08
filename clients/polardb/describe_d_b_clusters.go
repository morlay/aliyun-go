package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId        common.String
	PageNumber       common.Integer
	TotalRecordCount common.Integer
	PageRecordCount  common.Integer
	Items            DescribeDBClustersDBClusterList
}

type DescribeDBClustersDBCluster struct {
	DBClusterId          common.String
	DBClusterDescription common.String
	PayType              common.String
	DBClusterNetworkType common.String
	RegionId             common.String
	ExpireTime           common.String
	DBClusterStatus      common.String
	Engine               common.String
	DBType               common.String
	DBVersion            common.String
	LockMode             common.String
	LockReason           common.String
	CreateTime           common.String
	VpcId                common.String
	VSwitchId            common.String
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
