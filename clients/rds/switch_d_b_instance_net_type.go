package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SwitchDBInstanceNetTypeRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ConnectionStringType   string `position:"Query" name:"ConnectionStringType"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken            string `position:"Query" name:"ClientToken"`
	Port                   string `position:"Query" name:"Port"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
}

func (r SwitchDBInstanceNetTypeRequest) Invoke(client *sdk.Client) (response *SwitchDBInstanceNetTypeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SwitchDBInstanceNetTypeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "SwitchDBInstanceNetType", "rds", "")

	resp := struct {
		*responses.BaseResponse
		SwitchDBInstanceNetTypeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SwitchDBInstanceNetTypeResponse

	err = client.DoAction(&req, &resp)
	return
}

type SwitchDBInstanceNetTypeResponse struct {
	RequestId string
}
