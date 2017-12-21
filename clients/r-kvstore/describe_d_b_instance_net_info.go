package r_kvstore

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId           string
	InstanceNetworkType string
	NetInfoItems        DescribeDBInstanceNetInfoInstanceNetInfoList
}

type DescribeDBInstanceNetInfoInstanceNetInfo struct {
	ConnectionString  string
	IPAddress         string
	Port              string
	VPCId             string
	VSwitchId         string
	DBInstanceNetType string
	ExpiredTime       string
	Upgradeable       string
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
