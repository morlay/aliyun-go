package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GrantAccountPrivilegeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountPrivilege     string `position:"Query" name:"AccountPrivilege"`
}

func (r GrantAccountPrivilegeRequest) Invoke(client *sdk.Client) (response *GrantAccountPrivilegeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GrantAccountPrivilegeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "GrantAccountPrivilege", "rds", "")

	resp := struct {
		*responses.BaseResponse
		GrantAccountPrivilegeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GrantAccountPrivilegeResponse

	err = client.DoAction(&req, &resp)
	return
}

type GrantAccountPrivilegeResponse struct {
	RequestId string
}
