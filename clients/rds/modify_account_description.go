package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyAccountDescriptionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountDescription   string `position:"Query" name:"AccountDescription"`
}

func (req *ModifyAccountDescriptionRequest) Invoke(client *sdk.Client) (resp *ModifyAccountDescriptionResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyAccountDescription", "rds", "")
	resp = &ModifyAccountDescriptionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyAccountDescriptionResponse struct {
	responses.BaseResponse
	RequestId string
}
