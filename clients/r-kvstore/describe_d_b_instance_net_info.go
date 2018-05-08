package r_kvstore

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
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBInstanceNetInfoRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceNetInfoResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeDBInstanceNetInfo", "redisa", "")
	resp = &DescribeDBInstanceNetInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceNetInfoResponse struct {
	responses.BaseResponse
	RequestId           common.String
	InstanceNetworkType common.String
	NetInfoItems        DescribeDBInstanceNetInfoInstanceNetInfoList
}

type DescribeDBInstanceNetInfoInstanceNetInfo struct {
	ConnectionString  common.String
	IPAddress         common.String
	Port              common.String
	VPCId             common.String
	VSwitchId         common.String
	DBInstanceNetType common.String
	ExpiredTime       common.String
	Upgradeable       common.String
}

type DescribeDBInstanceNetInfoInstanceNetInfoList []DescribeDBInstanceNetInfoInstanceNetInfo

func (list *DescribeDBInstanceNetInfoInstanceNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoInstanceNetInfo)
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
