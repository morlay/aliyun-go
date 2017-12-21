package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetAccountPasswordRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ResetAccountPasswordRequest) Invoke(client *sdk.Client) (resp *ResetAccountPasswordResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "ResetAccountPassword", "polardb", "")
	resp = &ResetAccountPasswordResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetAccountPasswordResponse struct {
	responses.BaseResponse
	RequestId string
}
