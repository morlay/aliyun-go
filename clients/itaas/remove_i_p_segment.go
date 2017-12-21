package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveIPSegmentRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Uuid        string `position:"Query" name:"Uuid"`
}

func (req *RemoveIPSegmentRequest) Invoke(client *sdk.Client) (resp *RemoveIPSegmentResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "RemoveIPSegment", "", "")
	resp = &RemoveIPSegmentResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveIPSegmentResponse struct {
	responses.BaseResponse
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList RemoveIPSegmentErrorMessageList
}

type RemoveIPSegmentErrorMessage struct {
	ErrorMessage string
}

type RemoveIPSegmentErrorMessageList []RemoveIPSegmentErrorMessage

func (list *RemoveIPSegmentErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveIPSegmentErrorMessage)
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
