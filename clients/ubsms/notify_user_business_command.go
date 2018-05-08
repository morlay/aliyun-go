package ubsms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type NotifyUserBusinessCommandRequest struct {
	requests.RpcRequest
	Uid         string `position:"Query" name:"Uid"`
	ServiceCode string `position:"Query" name:"ServiceCode"`
	Cmd         string `position:"Query" name:"Cmd"`
	Region      string `position:"Query" name:"Region"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	Password    string `position:"Query" name:"Password"`
}

func (req *NotifyUserBusinessCommandRequest) Invoke(client *sdk.Client) (resp *NotifyUserBusinessCommandResponse, err error) {
	req.InitWithApiInfo("Ubsms", "2015-06-23", "NotifyUserBusinessCommand", "", "")
	resp = &NotifyUserBusinessCommandResponse{}
	err = client.DoAction(req, resp)
	return
}

type NotifyUserBusinessCommandResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
