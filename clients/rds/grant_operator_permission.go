package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GrantOperatorPermissionRequest struct {
	requests.RpcRequest
	Privileges           string `position:"Query" name:"Privileges"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ExpiredTime          string `position:"Query" name:"ExpiredTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *GrantOperatorPermissionRequest) Invoke(client *sdk.Client) (resp *GrantOperatorPermissionResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "GrantOperatorPermission", "rds", "")
	resp = &GrantOperatorPermissionResponse{}
	err = client.DoAction(req, resp)
	return
}

type GrantOperatorPermissionResponse struct {
	responses.BaseResponse
	RequestId string
}
