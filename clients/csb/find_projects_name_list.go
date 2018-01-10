package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindProjectsNameListRequest struct {
	requests.RpcRequest
	OperationFlag string `position:"Query" name:"OperationFlag"`
	CsbId         int64  `position:"Query" name:"CsbId"`
}

func (req *FindProjectsNameListRequest) Invoke(client *sdk.Client) (resp *FindProjectsNameListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "FindProjectsNameList", "CSB", "")
	resp = &FindProjectsNameListResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindProjectsNameListResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      FindProjectsNameListData
}

type FindProjectsNameListData struct {
	ProjectNameList FindProjectsNameListProjectNameListList
}

type FindProjectsNameListProjectNameListList []string

func (list *FindProjectsNameListProjectNameListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
