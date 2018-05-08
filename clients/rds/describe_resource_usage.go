package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeResourceUsageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeResourceUsageRequest) Invoke(client *sdk.Client) (resp *DescribeResourceUsageResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeResourceUsage", "rds", "")
	resp = &DescribeResourceUsageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeResourceUsageResponse struct {
	responses.BaseResponse
	RequestId      common.String
	DBInstanceId   common.String
	Engine         common.String
	DiskUsed       common.Long
	DataSize       common.Long
	LogSize        common.Long
	BackupSize     common.Long
	SQLSize        common.Long
	ColdBackupSize common.Long
}
