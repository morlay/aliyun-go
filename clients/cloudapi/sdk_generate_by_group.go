package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SdkGenerateByGroupRequest struct {
	requests.RpcRequest
	GroupId  string `position:"Query" name:"GroupId"`
	Language string `position:"Query" name:"Language"`
}

func (req *SdkGenerateByGroupRequest) Invoke(client *sdk.Client) (resp *SdkGenerateByGroupResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SdkGenerateByGroup", "apigateway", "")
	resp = &SdkGenerateByGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type SdkGenerateByGroupResponse struct {
	responses.BaseResponse
	RequestId    string
	DownloadLink string
}
