package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ImportDatabaseBetweenInstancesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SourceDBInstanceId   string `position:"Query" name:"SourceDBInstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInfo               string `position:"Query" name:"DBInfo"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ImportDatabaseBetweenInstancesRequest) Invoke(client *sdk.Client) (resp *ImportDatabaseBetweenInstancesResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ImportDatabaseBetweenInstances", "rds", "")
	resp = &ImportDatabaseBetweenInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ImportDatabaseBetweenInstancesResponse struct {
	responses.BaseResponse
	RequestId common.String
	ImportId  common.String
}
