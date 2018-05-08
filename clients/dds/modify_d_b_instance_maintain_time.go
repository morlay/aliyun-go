package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDBInstanceMaintainTimeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	MaintainStartTime    string `position:"Query" name:"MaintainStartTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MaintainEndTime      string `position:"Query" name:"MaintainEndTime"`
}

func (req *ModifyDBInstanceMaintainTimeRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceMaintainTimeResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyDBInstanceMaintainTime", "dds", "")
	resp = &ModifyDBInstanceMaintainTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceMaintainTimeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
