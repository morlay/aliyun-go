package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveAppRecordConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveAppRecordConfigRequest) Invoke(client *sdk.Client) (response *DescribeLiveAppRecordConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveAppRecordConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveAppRecordConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveAppRecordConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveAppRecordConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveAppRecordConfigResponse struct {
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
