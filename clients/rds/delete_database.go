package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDatabaseRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteDatabaseRequest) Invoke(client *sdk.Client) (response *DeleteDatabaseResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDatabaseRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DeleteDatabase", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DeleteDatabaseResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteDatabaseResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteDatabaseResponse struct {
	RequestId string
}
