package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDBInstanceNetworkTypeRequest struct {
	requests.RpcRequest
	ResourceOwnerId                      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount                 string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                         string `position:"Query" name:"OwnerAccount"`
	OwnerId                              int64  `position:"Query" name:"OwnerId"`
	VSwitchId                            string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress                     string `position:"Query" name:"PrivateIpAddress"`
	RetainClassic                        string `position:"Query" name:"RetainClassic"`
	ClassicExpiredDays                   string `position:"Query" name:"ClassicExpiredDays"`
	VPCId                                string `position:"Query" name:"VPCId"`
	DBInstanceId                         string `position:"Query" name:"DBInstanceId"`
	ReadWriteSplittingPrivateIpAddress   string `position:"Query" name:"ReadWriteSplittingPrivateIpAddress"`
	InstanceNetworkType                  string `position:"Query" name:"InstanceNetworkType"`
	ReadWriteSplittingClassicExpiredDays int    `position:"Query" name:"ReadWriteSplittingClassicExpiredDays"`
}

func (req *ModifyDBInstanceNetworkTypeRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceNetworkTypeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceNetworkType", "rds", "")
	resp = &ModifyDBInstanceNetworkTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceNetworkTypeResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskId    common.String
}
