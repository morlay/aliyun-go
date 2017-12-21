package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelImportRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImportId             int    `position:"Query" name:"ImportId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CancelImportRequest) Invoke(client *sdk.Client) (response *CancelImportResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CancelImportRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CancelImport", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CancelImportResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CancelImportResponse

	err = client.DoAction(&req, &resp)
	return
}

type CancelImportResponse struct {
	RequestId string
}
