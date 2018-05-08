package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SdkGenerateByAppRequest struct {
	requests.RpcRequest
	AppId    int64  `position:"Query" name:"AppId"`
	Language string `position:"Query" name:"Language"`
}

func (req *SdkGenerateByAppRequest) Invoke(client *sdk.Client) (resp *SdkGenerateByAppResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SdkGenerateByApp", "apigateway", "")
	resp = &SdkGenerateByAppResponse{}
	err = client.DoAction(req, resp)
	return
}

type SdkGenerateByAppResponse struct {
	responses.BaseResponse
	RequestId    common.String
	DownloadLink common.String
}
