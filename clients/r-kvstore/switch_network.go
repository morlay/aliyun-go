package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SwitchNetworkRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	TargetNetworkType    string `position:"Query" name:"TargetNetworkType"`
	RetainClassic        string `position:"Query" name:"RetainClassic"`
	ClassicExpiredDays   string `position:"Query" name:"ClassicExpiredDays"`
	VpcId                string `position:"Query" name:"VpcId"`
}

func (req *SwitchNetworkRequest) Invoke(client *sdk.Client) (resp *SwitchNetworkResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "SwitchNetwork", "redisa", "")
	resp = &SwitchNetworkResponse{}
	err = client.DoAction(req, resp)
	return
}

type SwitchNetworkResponse struct {
	responses.BaseResponse
	RequestId string
	TaskId    string
}
