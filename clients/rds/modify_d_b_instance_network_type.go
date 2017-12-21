package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceNetworkTypeRequest struct {
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

func (r ModifyDBInstanceNetworkTypeRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceNetworkTypeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceNetworkTypeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceNetworkType", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceNetworkTypeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceNetworkTypeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceNetworkTypeResponse struct {
	RequestId string
	TaskId    string
}
