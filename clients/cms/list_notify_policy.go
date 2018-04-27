package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListNotifyPolicyRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	AlertName  string `position:"Query" name:"AlertName"`
	PageSize   int    `position:"Query" name:"PageSize"`
	Id         string `position:"Query" name:"Id"`
	Dimensions string `position:"Query" name:"Dimensions"`
}

func (req *ListNotifyPolicyRequest) Invoke(client *sdk.Client) (resp *ListNotifyPolicyResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ListNotifyPolicy", "cms", "")
	resp = &ListNotifyPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListNotifyPolicyResponse struct {
	responses.BaseResponse
	Code             string
	Message          string
	Success          string
	TraceId          string
	Total            int
	NotifyPolicyList ListNotifyPolicyNotifyPolicyList
}

type ListNotifyPolicyNotifyPolicy struct {
	AlertName  string
	Dimensions string
	Type       string
	Id         string
	StartTime  int64
	EndTime    int64
}

type ListNotifyPolicyNotifyPolicyList []ListNotifyPolicyNotifyPolicy

func (list *ListNotifyPolicyNotifyPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListNotifyPolicyNotifyPolicy)
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
