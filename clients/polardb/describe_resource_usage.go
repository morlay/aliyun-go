package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourceUsageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeResourceUsageRequest) Invoke(client *sdk.Client) (resp *DescribeResourceUsageResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeResourceUsage", "polardb", "")
	resp = &DescribeResourceUsageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeResourceUsageResponse struct {
	responses.BaseResponse
	RequestId    string
	DBInstanceId string
	Engine       string
	DBType       string
	DBVersion    string
	DiskUsed     int64
	DataSize     int64
	LogSize      int64
	BackupSize   int64
}
