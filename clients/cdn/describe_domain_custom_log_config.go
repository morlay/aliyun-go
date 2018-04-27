package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainCustomLogConfigRequest struct {
	requests.RpcRequest
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

func (req *DescribeDomainCustomLogConfigRequest) Invoke(client *sdk.Client) (resp *DescribeDomainCustomLogConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainCustomLogConfig", "", "")
	resp = &DescribeDomainCustomLogConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainCustomLogConfigResponse struct {
	responses.BaseResponse
	RequestId string
	ConfigId  string
	Remark    string
	Sample    string
	Tag       string
}
