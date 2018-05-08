package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type MigrateToOtherZoneRequest struct {
	requests.RpcRequest
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	EffectiveTime        string `position:"Query" name:"EffectiveTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *MigrateToOtherZoneRequest) Invoke(client *sdk.Client) (resp *MigrateToOtherZoneResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "MigrateToOtherZone", "rds", "")
	resp = &MigrateToOtherZoneResponse{}
	err = client.DoAction(req, resp)
	return
}

type MigrateToOtherZoneResponse struct {
	responses.BaseResponse
	RequestId common.String
}
