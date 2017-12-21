package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddIPSegmentRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Ipsegment   string `position:"Query" name:"Ipsegment"`
	Memo        string `position:"Query" name:"Memo"`
}

func (req *AddIPSegmentRequest) Invoke(client *sdk.Client) (resp *AddIPSegmentResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "AddIPSegment", "", "")
	resp = &AddIPSegmentResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddIPSegmentResponse struct {
	responses.BaseResponse
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
