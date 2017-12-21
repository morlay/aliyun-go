package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamStreamStatusRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamStreamStatusRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamStreamStatusResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamStreamStatus", "", "")
	resp = &DescribeLiveStreamStreamStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamStreamStatusResponse struct {
	responses.BaseResponse
	RequestId    string
	StreamStatus string
}
