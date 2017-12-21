package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GrantOperatorPermissionRequest struct {
	Privileges           string `position:"Query" name:"Privileges"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ExpiredTime          string `position:"Query" name:"ExpiredTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r GrantOperatorPermissionRequest) Invoke(client *sdk.Client) (response *GrantOperatorPermissionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GrantOperatorPermissionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "GrantOperatorPermission", "rds", "")

	resp := struct {
		*responses.BaseResponse
		GrantOperatorPermissionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GrantOperatorPermissionResponse

	err = client.DoAction(&req, &resp)
	return
}

type GrantOperatorPermissionResponse struct {
	RequestId string
}
