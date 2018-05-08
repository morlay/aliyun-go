package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RevokeAccountPrivilegeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RevokeAccountPrivilegeRequest) Invoke(client *sdk.Client) (resp *RevokeAccountPrivilegeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "RevokeAccountPrivilege", "rds", "")
	resp = &RevokeAccountPrivilegeResponse{}
	err = client.DoAction(req, resp)
	return
}

type RevokeAccountPrivilegeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
