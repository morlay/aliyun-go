package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetAccountForPGRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ResetAccountForPGRequest) Invoke(client *sdk.Client) (response *ResetAccountForPGResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ResetAccountForPGRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ResetAccountForPG", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ResetAccountForPGResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ResetAccountForPGResponse

	err = client.DoAction(&req, &resp)
	return
}

type ResetAccountForPGResponse struct {
	RequestId string
}
