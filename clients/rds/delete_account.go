package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAccountRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteAccountRequest) Invoke(client *sdk.Client) (response *DeleteAccountResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteAccountRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DeleteAccount", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DeleteAccountResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteAccountResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteAccountResponse struct {
	RequestId string
}
