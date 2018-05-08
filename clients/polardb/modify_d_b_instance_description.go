package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDBInstanceDescriptionRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	DBInstanceId          string `position:"Query" name:"DBInstanceId"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyDBInstanceDescriptionRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceDescriptionResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBInstanceDescription", "polardb", "")
	resp = &ModifyDBInstanceDescriptionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceDescriptionResponse struct {
	responses.BaseResponse
	RequestId common.String
}
