package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddBusinessCategoryRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *AddBusinessCategoryRequest) Invoke(client *sdk.Client) (resp *AddBusinessCategoryResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "AddBusinessCategory", "", "")
	resp = &AddBusinessCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddBusinessCategoryResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
