package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateUploadPathForSQLServerRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateUploadPathForSQLServerRequest) Invoke(client *sdk.Client) (resp *CreateUploadPathForSQLServerResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateUploadPathForSQLServer", "rds", "")
	resp = &CreateUploadPathForSQLServerResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateUploadPathForSQLServerResponse struct {
	responses.BaseResponse
	RequestId         string
	InternetFtpServer string
	InternetPort      int
	IntranetFtpserver string
	Intranetport      int
	UserName          string
	Password          string
	FileName          string
}
