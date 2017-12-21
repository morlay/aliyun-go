package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceNetExpireTimeRequest struct {
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken            string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString         string `position:"Query" name:"ConnectionString"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	DBInstanceId             string `position:"Query" name:"DBInstanceId"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	ClassicExpendExpiredDays int    `position:"Query" name:"ClassicExpendExpiredDays"`
}

func (r ModifyDBInstanceNetExpireTimeRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceNetExpireTimeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceNetExpireTimeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyDBInstanceNetExpireTime", "dds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceNetExpireTimeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceNetExpireTimeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceNetExpireTimeResponse struct {
	RequestId string
}
