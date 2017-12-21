package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBClusterNetInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBClusterNetInfoRequest) Invoke(client *sdk.Client) (resp *DescribeDBClusterNetInfoResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterNetInfo", "polardb", "")
	resp = &DescribeDBClusterNetInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBClusterNetInfoResponse struct {
	responses.BaseResponse
	RequestId          string
	ClusterNetworkType string
	DBClusterNetInfos  DescribeDBClusterNetInfoDBClusterNetInfoList
}

type DescribeDBClusterNetInfoDBClusterNetInfo struct {
	ConnectionString string
	IPAddress        string
	IPType           string
	Port             string
	VPCId            string
	VSwitchId        string
}

type DescribeDBClusterNetInfoDBClusterNetInfoList []DescribeDBClusterNetInfoDBClusterNetInfo

func (list *DescribeDBClusterNetInfoDBClusterNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterNetInfoDBClusterNetInfo)
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
