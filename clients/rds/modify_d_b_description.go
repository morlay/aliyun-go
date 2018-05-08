package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDBDescriptionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	DBDescription        string `position:"Query" name:"DBDescription"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyDBDescriptionRequest) Invoke(client *sdk.Client) (resp *ModifyDBDescriptionResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBDescription", "rds", "")
	resp = &ModifyDBDescriptionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBDescriptionResponse struct {
	responses.BaseResponse
	RequestId common.String
}
