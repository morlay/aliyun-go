package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDatabaseRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	DBDescription        string `position:"Query" name:"DBDescription"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CharacterSetName     string `position:"Query" name:"CharacterSetName"`
}

func (r CreateDatabaseRequest) Invoke(client *sdk.Client) (response *CreateDatabaseResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateDatabaseRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateDatabase", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CreateDatabaseResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateDatabaseResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateDatabaseResponse struct {
	RequestId string
}
