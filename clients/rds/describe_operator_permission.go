package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeOperatorPermissionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeOperatorPermissionRequest) Invoke(client *sdk.Client) (resp *DescribeOperatorPermissionResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOperatorPermission", "rds", "")
	resp = &DescribeOperatorPermissionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeOperatorPermissionResponse struct {
	responses.BaseResponse
	RequestId   string
	Privileges  string
	CreatedTime string
	ExpiredTime string
}
