package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceDescriptionRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	DBInstanceId          string `position:"Query" name:"DBInstanceId"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	NodeId                string `position:"Query" name:"NodeId"`
}

func (req *ModifyDBInstanceDescriptionRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceDescriptionResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyDBInstanceDescription", "dds", "")
	resp = &ModifyDBInstanceDescriptionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceDescriptionResponse struct {
	responses.BaseResponse
	RequestId string
}
