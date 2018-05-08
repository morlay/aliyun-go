package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateVSwitchRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpcId                string `position:"Query" name:"VpcId"`
	VSwitchName          string `position:"Query" name:"VSwitchName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CidrBlock            string `position:"Query" name:"CidrBlock"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateVSwitchRequest) Invoke(client *sdk.Client) (resp *CreateVSwitchResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateVSwitch", "vpc", "")
	resp = &CreateVSwitchResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateVSwitchResponse struct {
	responses.BaseResponse
	RequestId common.String
	VSwitchId common.String
}
