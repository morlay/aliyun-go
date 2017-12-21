package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstanceNetInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDBInstanceNetInfoRequest) Invoke(client *sdk.Client) (response *DescribeDBInstanceNetInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBInstanceNetInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBInstanceNetInfo", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBInstanceNetInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDBInstanceNetInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBInstanceNetInfoResponse struct {
	RequestId           string
	InstanceNetworkType string
	DBInstanceNetInfos  DescribeDBInstanceNetInfoDBInstanceNetInfoList
}

type DescribeDBInstanceNetInfoDBInstanceNetInfo struct {
	ConnectionString string
	IPAddress        string
	IPType           string
	Port             string
	VPCId            string
	VSwitchId        string
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
