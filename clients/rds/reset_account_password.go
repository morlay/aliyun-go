package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetAccountPasswordRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ResetAccountPasswordRequest) Invoke(client *sdk.Client) (response *ResetAccountPasswordResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ResetAccountPasswordRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ResetAccountPassword", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ResetAccountPasswordResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ResetAccountPasswordResponse

	err = client.DoAction(&req, &resp)
	return
}

type ResetAccountPasswordResponse struct {
	RequestId string
}
