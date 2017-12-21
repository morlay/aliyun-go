package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SwitchDBInstanceChargeTypeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r SwitchDBInstanceChargeTypeRequest) Invoke(client *sdk.Client) (response *SwitchDBInstanceChargeTypeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SwitchDBInstanceChargeTypeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "SwitchDBInstanceChargeType", "rds", "")

	resp := struct {
		*responses.BaseResponse
		SwitchDBInstanceChargeTypeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SwitchDBInstanceChargeTypeResponse

	err = client.DoAction(&req, &resp)
	return
}

type SwitchDBInstanceChargeTypeResponse struct {
	RequestId string
}
