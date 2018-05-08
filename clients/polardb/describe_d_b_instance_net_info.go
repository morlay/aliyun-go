package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBInstanceNetInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBInstanceNetInfoRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceNetInfoResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBInstanceNetInfo", "polardb", "")
	resp = &DescribeDBInstanceNetInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceNetInfoResponse struct {
	responses.BaseResponse
	RequestId           common.String
	InstanceNetworkType common.String
	DBInstanceNetInfos  DescribeDBInstanceNetInfoDBInstanceNetInfoList
}

type DescribeDBInstanceNetInfoDBInstanceNetInfo struct {
	ConnectionString common.String
	IPAddress        common.String
	IPType           common.String
	Port             common.String
	VPCId            common.String
	VSwitchId        common.String
}

type DescribeDBInstanceNetInfoDBInstanceNetInfoList []DescribeDBInstanceNetInfoDBInstanceNetInfo

func (list *DescribeDBInstanceNetInfoDBInstanceNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoDBInstanceNetInfo)
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
