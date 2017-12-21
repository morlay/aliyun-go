package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourceUsageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeResourceUsageRequest) Invoke(client *sdk.Client) (response *DescribeResourceUsageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeResourceUsageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeResourceUsage", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeResourceUsageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeResourceUsageResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeResourceUsageResponse struct {
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
