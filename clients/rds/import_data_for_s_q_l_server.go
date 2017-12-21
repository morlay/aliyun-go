package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImportDataForSQLServerRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	FileName             string `position:"Query" name:"FileName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ImportDataForSQLServerRequest) Invoke(client *sdk.Client) (response *ImportDataForSQLServerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ImportDataForSQLServerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ImportDataForSQLServer", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ImportDataForSQLServerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ImportDataForSQLServerResponse

	err = client.DoAction(&req, &resp)
	return
}

type ImportDataForSQLServerResponse struct {
	RequestId string
	ImportID  int
}
