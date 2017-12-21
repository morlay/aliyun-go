package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImportDatabaseBetweenInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SourceDBInstanceId   string `position:"Query" name:"SourceDBInstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInfo               string `position:"Query" name:"DBInfo"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ImportDatabaseBetweenInstancesRequest) Invoke(client *sdk.Client) (response *ImportDatabaseBetweenInstancesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ImportDatabaseBetweenInstancesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ImportDatabaseBetweenInstances", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ImportDatabaseBetweenInstancesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ImportDatabaseBetweenInstancesResponse

	err = client.DoAction(&req, &resp)
	return
}

type ImportDatabaseBetweenInstancesResponse struct {
	RequestId string
	ImportId  string
}
