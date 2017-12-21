package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelImportRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImportId             int    `position:"Query" name:"ImportId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CancelImportRequest) Invoke(client *sdk.Client) (resp *CancelImportResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CancelImport", "rds", "")
	resp = &CancelImportResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelImportResponse struct {
	responses.BaseResponse
	RequestId string
}
