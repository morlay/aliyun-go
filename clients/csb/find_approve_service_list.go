package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FindApproveServiceListRequest struct {
	requests.RpcRequest
	ApproveLevel   string `position:"Query" name:"ApproveLevel"`
	ProjectName    string `position:"Query" name:"ProjectName"`
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
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      FindApproveServiceListData
}

type FindApproveServiceListData struct {
	Total       common.Integer
	PageNumber  common.Integer
	CurrentPage common.Integer
	ServiceList FindApproveServiceListServiceList
}

type FindApproveServiceListService struct {
	AllVisiable    bool
	CasTargets     common.String
	CreateTime     common.Long
	CsbId          common.Long
	Id             common.Long
	InterfaceName  common.String
	ModifiedTime   common.Long
	OwnerId        common.String
	PrincipalName  common.String
	ProjectId      common.Long
	ProjectName    common.String
	Qps            common.Integer
	Scope          common.String
	ServiceName    common.String
	ServiceVersion common.String
	SkipAuth       bool
	StatisticName  common.String
	Status         common.Integer
	UserId         common.String
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
