package polardb

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
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DatabaseName         string `position:"Query" name:"DatabaseName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountDescription   string `position:"Query" name:"AccountDescription"`
}

func (req *CreateAccountRequest) Invoke(client *sdk.Client) (resp *CreateAccountResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "CreateAccount", "polardb", "")
	resp = &CreateAccountResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateAccountResponse struct {
	responses.BaseResponse
	RequestId common.String
}
