package dds

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
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountDescription   string `position:"Query" name:"AccountDescription"`
}

func (req *ModifyAccountDescriptionRequest) Invoke(client *sdk.Client) (resp *ModifyAccountDescriptionResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyAccountDescription", "dds", "")
	resp = &ModifyAccountDescriptionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyAccountDescriptionResponse struct {
	responses.BaseResponse
	RequestId common.String
}
