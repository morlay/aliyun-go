package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterNetInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeClusterNetInfoRequest) Invoke(client *sdk.Client) (response *DescribeClusterNetInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeClusterNetInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeClusterNetInfo", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeClusterNetInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeClusterNetInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeClusterNetInfoResponse struct {
	RequestId            string
	DBClusterNetworkType string
	DBInstanceNetInfos   DescribeClusterNetInfoDBInstanceNetInfoList
}

type DescribeClusterNetInfoDBInstanceNetInfo struct {
	ConnectionString string
	IPAddress        string
	IPType           string
	Port             string
	VPCId            string
	VSwitchId        string
}

type DescribeClusterNetInfoDBInstanceNetInfoList []DescribeClusterNetInfoDBInstanceNetInfo

func (list *DescribeClusterNetInfoDBInstanceNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterNetInfoDBInstanceNetInfo)
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
