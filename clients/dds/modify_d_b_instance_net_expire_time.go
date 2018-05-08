package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDBInstanceNetExpireTimeRequest struct {
	requests.RpcRequest
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken            string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString         string `position:"Query" name:"ConnectionString"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	DBInstanceId             string `position:"Query" name:"DBInstanceId"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	ClassicExpendExpiredDays int    `position:"Query" name:"ClassicExpendExpiredDays"`
}

func (req *ModifyDBInstanceNetExpireTimeRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceNetExpireTimeResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyDBInstanceNetExpireTime", "dds", "")
	resp = &ModifyDBInstanceNetExpireTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceNetExpireTimeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
