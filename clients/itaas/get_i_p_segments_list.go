package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetIPSegmentsListRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
}

func (req *GetIPSegmentsListRequest) Invoke(client *sdk.Client) (resp *GetIPSegmentsListResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "GetIPSegmentsList", "", "")
	resp = &GetIPSegmentsListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetIPSegmentsListResponse struct {
	responses.BaseResponse
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	Data      GetIPSegmentsListIpsegmentInfoList
	ErrorList GetIPSegmentsListErrorMessageList
}

type GetIPSegmentsListIpsegmentInfo struct {
	CreateTimeL int64
	Ipsegment   string
	Memo        string
	ModifyTimeL int64
	Uuid        string
}

type GetIPSegmentsListErrorMessage struct {
	ErrorMessage string
}

type GetIPSegmentsListIpsegmentInfoList []GetIPSegmentsListIpsegmentInfo

func (list *GetIPSegmentsListIpsegmentInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetIPSegmentsListIpsegmentInfo)
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

type GetIPSegmentsListErrorMessageList []GetIPSegmentsListErrorMessage

func (list *GetIPSegmentsListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetIPSegmentsListErrorMessage)
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
