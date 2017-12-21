package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateAccountRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DatabaseName         string `position:"Query" name:"DatabaseName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountDescription   string `position:"Query" name:"AccountDescription"`
}

func (r CreateAccountRequest) Invoke(client *sdk.Client) (response *CreateAccountResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateAccountRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "CreateAccount", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		CreateAccountResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateAccountResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateAccountResponse struct {
	RequestId string
}
