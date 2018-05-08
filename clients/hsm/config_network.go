package hsm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ConfigNetworkRequest struct {
	requests.RpcRequest
	VSwitchId       string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	VpcId           string `position:"Query" name:"VpcId"`
	Ip              string `position:"Query" name:"Ip"`
}

func (req *ConfigNetworkRequest) Invoke(client *sdk.Client) (resp *ConfigNetworkResponse, err error) {
	req.InitWithApiInfo("hsm", "2018-01-11", "ConfigNetwork", "hsm", "")
	resp = &ConfigNetworkResponse{}
	err = client.DoAction(req, resp)
	return
}

type ConfigNetworkResponse struct {
	responses.BaseResponse
	RequestId common.String
}
