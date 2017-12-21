package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RevokeOperatorPermissionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r RevokeOperatorPermissionRequest) Invoke(client *sdk.Client) (response *RevokeOperatorPermissionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RevokeOperatorPermissionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "RevokeOperatorPermission", "rds", "")

	resp := struct {
		*responses.BaseResponse
		RevokeOperatorPermissionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RevokeOperatorPermissionResponse

	err = client.DoAction(&req, &resp)
	return
}

type RevokeOperatorPermissionResponse struct {
	RequestId string
}
