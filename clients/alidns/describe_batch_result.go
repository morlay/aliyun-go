package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBatchResultRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	TraceId      string `position:"Query" name:"TraceId"`
}

func (r DescribeBatchResultRequest) Invoke(client *sdk.Client) (response *DescribeBatchResultResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeBatchResultRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeBatchResult", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeBatchResultResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeBatchResultResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeBatchResultResponse struct {
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
