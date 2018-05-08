package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SdkGenerateRequest struct {
	requests.RpcRequest
	GroupId  string `position:"Query" name:"GroupId"`
	AppId    int64  `position:"Query" name:"AppId"`
	Language string `position:"Query" name:"Language"`
}

func (req *SdkGenerateRequest) Invoke(client *sdk.Client) (resp *SdkGenerateResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SdkGenerate", "apigateway", "")
	resp = &SdkGenerateResponse{}
	err = client.DoAction(req, resp)
	return
}

type SdkGenerateResponse struct {
	responses.BaseResponse
	RequestId    common.String
	DownloadLink common.String
}
