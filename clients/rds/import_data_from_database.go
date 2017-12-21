package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImportDataFromDatabaseRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ImportDataType         string `position:"Query" name:"ImportDataType"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	IsLockTable            string `position:"Query" name:"IsLockTable"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	SourceDatabaseDBNames  string `position:"Query" name:"SourceDatabaseDBNames"`
	SourceDatabaseIp       string `position:"Query" name:"SourceDatabaseIp"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	SourceDatabasePassword string `position:"Query" name:"SourceDatabasePassword"`
	SourceDatabasePort     string `position:"Query" name:"SourceDatabasePort"`
	SourceDatabaseUserName string `position:"Query" name:"SourceDatabaseUserName"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
}

func (req *ImportDataFromDatabaseRequest) Invoke(client *sdk.Client) (resp *ImportDataFromDatabaseResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ImportDataFromDatabase", "rds", "")
	resp = &ImportDataFromDatabaseResponse{}
	err = client.DoAction(req, resp)
	return
}

type ImportDataFromDatabaseResponse struct {
	responses.BaseResponse
	RequestId string
	ImportId  int
}
