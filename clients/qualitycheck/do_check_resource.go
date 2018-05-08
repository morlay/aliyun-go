package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DoCheckResourceRequest struct {
	requests.RpcRequest
	Country        string `position:"Query" name:"Country"`
	Hid            int64  `position:"Query" name:"Hid"`
	Level          int64  `position:"Query" name:"Level"`
	Success        string `position:"Query" name:"Success"`
	Interrupt      string `position:"Query" name:"Interrupt"`
	GmtWakeup      string `position:"Query" name:"GmtWakeup"`
	Pk             string `position:"Query" name:"Pk"`
	Bid            string `position:"Query" name:"Bid"`
	Message        string `position:"Query" name:"Message"`
	Prompt         string `position:"Query" name:"Prompt"`
	TaskExtraData  string `position:"Query" name:"TaskExtraData"`
	TaskIdentifier string `position:"Query" name:"TaskIdentifier"`
}

func (req *DoCheckResourceRequest) Invoke(client *sdk.Client) (resp *DoCheckResourceResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "DoCheckResource", "", "")
	resp = &DoCheckResourceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DoCheckResourceResponse struct {
	responses.BaseResponse
	Interrupt       common.String
	Invoker         common.Long
	Pk              common.String
	Bid             common.String
	Hid             common.Long
	Country         common.String
	TaskIdentifier  common.String
	TaskIdentifier1 common.String
	GmtWakeup       common.String
	Success         bool
	Message         common.String
	Level           common.Long
	Url             common.String
	Prompt          common.String
}
