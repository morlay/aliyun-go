package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CheckAccountNameAvailableRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CheckAccountNameAvailableRequest) Invoke(client *sdk.Client) (resp *CheckAccountNameAvailableResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CheckAccountNameAvailable", "rds", "")
	resp = &CheckAccountNameAvailableResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckAccountNameAvailableResponse struct {
	responses.BaseResponse
	RequestId common.String
}
