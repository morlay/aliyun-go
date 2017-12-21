package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SwitchTempInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	TempInstanceId       string `position:"Domain" name:"TempInstanceId"`
	InstanceId           string `position:"Domain" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *SwitchTempInstanceRequest) Invoke(client *sdk.Client) (resp *SwitchTempInstanceResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "SwitchTempInstance", "redisa", "")
	resp = &SwitchTempInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type SwitchTempInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
