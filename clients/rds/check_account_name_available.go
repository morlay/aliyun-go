package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckAccountNameAvailableRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CheckAccountNameAvailableRequest) Invoke(client *sdk.Client) (response *CheckAccountNameAvailableResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CheckAccountNameAvailableRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CheckAccountNameAvailable", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CheckAccountNameAvailableResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CheckAccountNameAvailableResponse

	err = client.DoAction(&req, &resp)
	return
}

type CheckAccountNameAvailableResponse struct {
	RequestId string
}
