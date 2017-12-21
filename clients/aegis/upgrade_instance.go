package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpgradeInstanceRequest struct {
	InstanceId  string `position:"Query" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	VmNumber    int    `position:"Query" name:"VmNumber"`
	OwnerId     int64  `position:"Query" name:"OwnerId"`
	VersionCode int    `position:"Query" name:"VersionCode"`
}

func (r UpgradeInstanceRequest) Invoke(client *sdk.Client) (response *UpgradeInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpgradeInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("aegis", "2016-11-11", "UpgradeInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		UpgradeInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpgradeInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpgradeInstanceResponse struct {
	OrderId   string
	RequestId string
}
