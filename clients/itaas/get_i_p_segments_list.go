package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
	Success   bool
	Data      GetIPSegmentsListIpsegmentInfoList
	ErrorList GetIPSegmentsListErrorMessageList
}

type GetIPSegmentsListIpsegmentInfo struct {
	CreateTimeL common.Long
	Ipsegment   common.String
	Memo        common.String
	ModifyTimeL common.Long
	Uuid        common.String
}

type GetIPSegmentsListErrorMessage struct {
	ErrorMessage common.String
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
