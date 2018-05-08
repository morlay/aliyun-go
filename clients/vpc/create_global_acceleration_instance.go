package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateGlobalAccelerationInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthType        string `position:"Query" name:"BandwidthType"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ServiceLocation      string `position:"Query" name:"ServiceLocation"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	Name                 string `position:"Query" name:"Name"`
}

func (req *CreateGlobalAccelerationInstanceRequest) Invoke(client *sdk.Client) (resp *CreateGlobalAccelerationInstanceResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateGlobalAccelerationInstance", "vpc", "")
	resp = &CreateGlobalAccelerationInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateGlobalAccelerationInstanceResponse struct {
	responses.BaseResponse
	RequestId                    common.String
	GlobalAccelerationInstanceId common.String
	IpAddress                    common.String
}
