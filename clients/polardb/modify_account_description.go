package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyAccountDescriptionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountDescription   string `position:"Query" name:"AccountDescription"`
}

func (req *ModifyAccountDescriptionRequest) Invoke(client *sdk.Client) (resp *ModifyAccountDescriptionResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyAccountDescription", "polardb", "")
	resp = &ModifyAccountDescriptionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyAccountDescriptionResponse struct {
	responses.BaseResponse
	RequestId common.String
}
