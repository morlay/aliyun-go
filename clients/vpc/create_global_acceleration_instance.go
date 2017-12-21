package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateGlobalAccelerationInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ServiceLocation      string `position:"Query" name:"ServiceLocation"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateGlobalAccelerationInstanceRequest) Invoke(client *sdk.Client) (resp *CreateGlobalAccelerationInstanceResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateGlobalAccelerationInstance", "vpc", "")
	resp = &CreateGlobalAccelerationInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateGlobalAccelerationInstanceResponse struct {
	responses.BaseResponse
	RequestId                    string
	GlobalAccelerationInstanceId string
	IpAddress                    string
}
