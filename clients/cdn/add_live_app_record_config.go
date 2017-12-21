package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveAppRecordConfigRequest struct {
	OssBucket       string `position:"Query" name:"OssBucket"`
	AppName         string `position:"Query" name:"AppName"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	DomainName      string `position:"Query" name:"DomainName"`
	OssEndpoint     string `position:"Query" name:"OssEndpoint"`
	OssObjectPrefix string `position:"Query" name:"OssObjectPrefix"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
}

func (r AddLiveAppRecordConfigRequest) Invoke(client *sdk.Client) (response *AddLiveAppRecordConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLiveAppRecordConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "AddLiveAppRecordConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLiveAppRecordConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLiveAppRecordConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLiveAppRecordConfigResponse struct {
	RequestId string
}
