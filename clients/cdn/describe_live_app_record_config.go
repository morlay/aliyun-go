package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveAppRecordConfigRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveAppRecordConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLiveAppRecordConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveAppRecordConfig", "", "")
	resp = &DescribeLiveAppRecordConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveAppRecordConfigResponse struct {
	responses.BaseResponse
	RequestId     string
	LiveAppRecord DescribeLiveAppRecordConfigLiveAppRecord
}

type DescribeLiveAppRecordConfigLiveAppRecord struct {
	DomainName      string
	AppName         string
	OssEndpoint     string
	OssBucket       string
	OssObjectPrefix string
	CreateTime      string
}
