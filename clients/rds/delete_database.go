package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDatabaseRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteDatabaseRequest) Invoke(client *sdk.Client) (resp *DeleteDatabaseResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DeleteDatabase", "rds", "")
	resp = &DeleteDatabaseResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDatabaseResponse struct {
	responses.BaseResponse
	RequestId string
}
