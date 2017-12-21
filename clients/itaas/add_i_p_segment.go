package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddIPSegmentRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Ipsegment   string `position:"Query" name:"Ipsegment"`
	Memo        string `position:"Query" name:"Memo"`
}

func (r AddIPSegmentRequest) Invoke(client *sdk.Client) (response *AddIPSegmentResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddIPSegmentRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("ITaaS", "2017-05-05", "AddIPSegment", "", "")

	resp := struct {
		*responses.BaseResponse
		AddIPSegmentResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddIPSegmentResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddIPSegmentResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList AddIPSegmentErrorMessageList
}

type AddIPSegmentErrorMessage struct {
	ErrorMessage string
}

type AddIPSegmentErrorMessageList []AddIPSegmentErrorMessage

func (list *AddIPSegmentErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddIPSegmentErrorMessage)
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
