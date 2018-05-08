package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ImportDataForSQLServerRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	FileName             string `position:"Query" name:"FileName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ImportDataForSQLServerRequest) Invoke(client *sdk.Client) (resp *ImportDataForSQLServerResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ImportDataForSQLServer", "rds", "")
	resp = &ImportDataForSQLServerResponse{}
	err = client.DoAction(req, resp)
	return
}

type ImportDataForSQLServerResponse struct {
	responses.BaseResponse
	RequestId common.String
	ImportID  common.Integer
}
