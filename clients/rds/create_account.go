package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateAccountRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AccountType          string `position:"Query" name:"AccountType"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountDescription   string `position:"Query" name:"AccountDescription"`
}

func (req *CreateAccountRequest) Invoke(client *sdk.Client) (resp *CreateAccountResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateAccount", "rds", "")
	resp = &CreateAccountResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateAccountResponse struct {
	responses.BaseResponse
	RequestId common.String
}
