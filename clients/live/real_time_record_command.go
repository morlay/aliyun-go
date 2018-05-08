package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RealTimeRecordCommandRequest struct {
	requests.RpcRequest
	AppName    string `position:"Query" name:"AppName"`
	DomainName string `position:"Query" name:"DomainName"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Command    string `position:"Query" name:"Command"`
	StreamName string `position:"Query" name:"StreamName"`
}

func (req *RealTimeRecordCommandRequest) Invoke(client *sdk.Client) (resp *RealTimeRecordCommandResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "RealTimeRecordCommand", "live", "")
	resp = &RealTimeRecordCommandResponse{}
	err = client.DoAction(req, resp)
	return
}

type RealTimeRecordCommandResponse struct {
	responses.BaseResponse
	RequestId common.String
}
