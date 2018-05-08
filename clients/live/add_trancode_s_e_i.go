package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddTrancodeSEIRequest struct {
	requests.RpcRequest
	Delay      int    `position:"Query" name:"Delay"`
	AppName    string `position:"Query" name:"AppName"`
	Repeat     int    `position:"Query" name:"Repeat"`
	DomainName string `position:"Query" name:"DomainName"`
	Pattern    string `position:"Query" name:"Pattern"`
	Text       string `position:"Query" name:"Text"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	StreamName string `position:"Query" name:"StreamName"`
}

func (req *AddTrancodeSEIRequest) Invoke(client *sdk.Client) (resp *AddTrancodeSEIResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddTrancodeSEI", "live", "")
	resp = &AddTrancodeSEIResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddTrancodeSEIResponse struct {
	responses.BaseResponse
	RequestId common.String
}
