package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ResetAccountPasswordRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ResetAccountPasswordRequest) Invoke(client *sdk.Client) (resp *ResetAccountPasswordResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "ResetAccountPassword", "dds", "")
	resp = &ResetAccountPasswordResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetAccountPasswordResponse struct {
	responses.BaseResponse
	RequestId common.String
}
