package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstanceSSLRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBInstanceSSLRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceSSLResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceSSL", "rds", "")
	resp = &DescribeDBInstanceSSLResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceSSLResponse struct {
	responses.BaseResponse
	RequestId           string
	ConnectionString    string
	SSLExpireTime       string
	RequireUpdate       string
	RequireUpdateReason string
}
