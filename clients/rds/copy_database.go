package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CopyDatabaseRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CopyDatabaseRequest) Invoke(client *sdk.Client) (resp *CopyDatabaseResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CopyDatabase", "rds", "")
	resp = &CopyDatabaseResponse{}
	err = client.DoAction(req, resp)
	return
}

type CopyDatabaseResponse struct {
	responses.BaseResponse
	DBName   common.String
	DBStatus common.String
	TaskId   common.String
}
