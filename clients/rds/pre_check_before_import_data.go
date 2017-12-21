package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PreCheckBeforeImportDataRequest struct {
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

func (r PreCheckBeforeImportDataRequest) Invoke(client *sdk.Client) (response *PreCheckBeforeImportDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PreCheckBeforeImportDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "PreCheckBeforeImportData", "rds", "")

	resp := struct {
		*responses.BaseResponse
		PreCheckBeforeImportDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PreCheckBeforeImportDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type PreCheckBeforeImportDataResponse struct {
	RequestId  string
	PreCheckId string
}
