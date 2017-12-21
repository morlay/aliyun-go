package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RevokeAccountPrivilegeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r RevokeAccountPrivilegeRequest) Invoke(client *sdk.Client) (response *RevokeAccountPrivilegeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RevokeAccountPrivilegeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "RevokeAccountPrivilege", "rds", "")

	resp := struct {
		*responses.BaseResponse
		RevokeAccountPrivilegeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RevokeAccountPrivilegeResponse

	err = client.DoAction(&req, &resp)
	return
}

type RevokeAccountPrivilegeResponse struct {
	RequestId string
}
