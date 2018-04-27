package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    int    `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	EffectiveTime        string `position:"Query" name:"EffectiveTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PayType              string `position:"Query" name:"PayType"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
}

func (req *ModifyDBInstanceSpecRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceSpecResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceSpec", "rds", "")
	resp = &ModifyDBInstanceSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceSpecResponse struct {
	responses.BaseResponse
	RequestId string
}
