package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateDatabaseRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	DBDescription        string `position:"Query" name:"DBDescription"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CharacterSetName     string `position:"Query" name:"CharacterSetName"`
}

func (req *CreateDatabaseRequest) Invoke(client *sdk.Client) (resp *CreateDatabaseResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateDatabase", "rds", "")
	resp = &CreateDatabaseResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateDatabaseResponse struct {
	responses.BaseResponse
	RequestId common.String
}
