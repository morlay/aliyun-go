package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBatchResultRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	TraceId      string `position:"Query" name:"TraceId"`
}

func (req *DescribeBatchResultRequest) Invoke(client *sdk.Client) (resp *DescribeBatchResultResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeBatchResult", "", "")
	resp = &DescribeBatchResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBatchResultResponse struct {
	responses.BaseResponse
	RequestId     string
	TraceId       string
	Status        int64
	BatchCount    int64
	SuccessNumber int64
	FailResults   DescribeBatchResultFailResultList
}

type DescribeBatchResultFailResult struct {
	BatchIndex string
	ErrorCode  string
}

type DescribeBatchResultFailResultList []DescribeBatchResultFailResult

func (list *DescribeBatchResultFailResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBatchResultFailResult)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
