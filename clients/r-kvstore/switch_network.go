
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type SwitchNetworkRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
VSwitchId string `position:"Query" name:"VSwitchId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
TargetNetworkType string `position:"Query" name:"TargetNetworkType"`
RetainClassic string `position:"Query" name:"RetainClassic"`
ClassicExpiredDays string `position:"Query" name:"ClassicExpiredDays"`
VpcId string `position:"Query" name:"VpcId"`
}

func (r SwitchNetworkRequest) Invoke(client *sdk.Client) (response *SwitchNetworkResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SwitchNetworkRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "SwitchNetwork", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		SwitchNetworkResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.SwitchNetworkResponse

	err = client.DoAction(&req, &resp)
	return
}

type SwitchNetworkResponse struct {
RequestId string
TaskId string
}

