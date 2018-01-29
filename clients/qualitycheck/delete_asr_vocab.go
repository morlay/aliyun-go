package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAsrVocabRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *DeleteAsrVocabRequest) Invoke(client *sdk.Client) (resp *DeleteAsrVocabResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "DeleteAsrVocab", "", "")
	resp = &DeleteAsrVocabResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAsrVocabResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
