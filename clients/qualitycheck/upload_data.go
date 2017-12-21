package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadDataRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *UploadDataRequest) Invoke(client *sdk.Client) (resp *UploadDataResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadData", "", "")
	resp = &UploadDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type UploadDataResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
