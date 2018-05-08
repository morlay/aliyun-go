package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      FindProjectsNameListData
}

type FindProjectsNameListData struct {
	ProjectNameList FindProjectsNameListProjectNameListList
}

type FindProjectsNameListProjectNameListList []common.String

func (list *FindProjectsNameListProjectNameListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
