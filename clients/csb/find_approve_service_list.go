package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindApproveServiceListRequest struct {
	requests.RpcRequest
	ProjectName    string `position:"Query" name:"ProjectName"`
	ApproveLevel   string `position:"Query" name:"ApproveLevel"`
	ShowDelService string `position:"Query" name:"ShowDelService"`
	CsbId          int64  `position:"Query" name:"CsbId"`
	Alias          string `position:"Query" name:"Alias"`
	ServiceName    string `position:"Query" name:"ServiceName"`
}

func (req *FindApproveServiceListRequest) Invoke(client *sdk.Client) (resp *FindApproveServiceListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "FindApproveServiceList", "CSB", "")
	resp = &FindApproveServiceListResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindApproveServiceListResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      FindApproveServiceListData
}

type FindApproveServiceListData struct {
	Total       int
	PageNumber  int
	CurrentPage int
	ServiceList FindApproveServiceListServiceList
}

type FindApproveServiceListService struct {
	AllVisiable    bool
	CasTargets     string
	CreateTime     int64
	CsbId          int64
	Id             int64
	InterfaceName  string
	ModifiedTime   int64
	OwnerId        string
	PrincipalName  string
	ProjectId      int64
	ProjectName    string
	Qps            int
	Scope          string
	ServiceName    string
	ServiceVersion string
	SkipAuth       bool
	StatisticName  string
	Status         int
	UserId         string
}

type FindApproveServiceListServiceList []FindApproveServiceListService

func (list *FindApproveServiceListServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindApproveServiceListService)
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
