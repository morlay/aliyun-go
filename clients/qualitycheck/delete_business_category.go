package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteBusinessCategoryRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *DeleteBusinessCategoryRequest) Invoke(client *sdk.Client) (resp *DeleteBusinessCategoryResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "DeleteBusinessCategory", "", "")
	resp = &DeleteBusinessCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteBusinessCategoryResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      common.String
}
