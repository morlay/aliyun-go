package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ResetAccountRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ResetAccountRequest) Invoke(client *sdk.Client) (resp *ResetAccountResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ResetAccount", "rds", "")
	resp = &ResetAccountResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetAccountResponse struct {
	responses.BaseResponse
	RequestId common.String
}
