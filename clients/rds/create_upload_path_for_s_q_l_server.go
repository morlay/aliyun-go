package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateUploadPathForSQLServerRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CreateUploadPathForSQLServerRequest) Invoke(client *sdk.Client) (response *CreateUploadPathForSQLServerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateUploadPathForSQLServerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateUploadPathForSQLServer", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CreateUploadPathForSQLServerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateUploadPathForSQLServerResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateUploadPathForSQLServerResponse struct {
	RequestId         string
	InternetFtpServer string
	InternetPort      int
	IntranetFtpserver string
	Intranetport      int
	UserName          string
	Password          string
	FileName          string
}
