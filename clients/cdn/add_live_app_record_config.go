package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveAppRecordConfigRequest struct {
	requests.RpcRequest
	OssBucket       string `position:"Query" name:"OssBucket"`
	AppName         string `position:"Query" name:"AppName"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	DomainName      string `position:"Query" name:"DomainName"`
	OssEndpoint     string `position:"Query" name:"OssEndpoint"`
	OssObjectPrefix string `position:"Query" name:"OssObjectPrefix"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
}

func (req *AddLiveAppRecordConfigRequest) Invoke(client *sdk.Client) (resp *AddLiveAppRecordConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "AddLiveAppRecordConfig", "", "")
	resp = &AddLiveAppRecordConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddLiveAppRecordConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
