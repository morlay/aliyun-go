package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateAsrVocabRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *CreateAsrVocabRequest) Invoke(client *sdk.Client) (resp *CreateAsrVocabResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "CreateAsrVocab", "", "")
	resp = &CreateAsrVocabResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateAsrVocabResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
