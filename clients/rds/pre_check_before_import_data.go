package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PreCheckBeforeImportDataRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ImportDataType         string `position:"Query" name:"ImportDataType"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken            string `position:"Query" name:"ClientToken"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	SourceDatabaseDBNames  string `position:"Query" name:"SourceDatabaseDBNames"`
	SourceDatabaseIp       string `position:"Query" name:"SourceDatabaseIp"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	SourceDatabasePassword string `position:"Query" name:"SourceDatabasePassword"`
	SourceDatabasePort     string `position:"Query" name:"SourceDatabasePort"`
	SourceDatabaseUserName string `position:"Query" name:"SourceDatabaseUserName"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
}

func (req *PreCheckBeforeImportDataRequest) Invoke(client *sdk.Client) (resp *PreCheckBeforeImportDataResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "PreCheckBeforeImportData", "rds", "")
	resp = &PreCheckBeforeImportDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type PreCheckBeforeImportDataResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PreCheckId common.String
}
