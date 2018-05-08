package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FindServiceListRequest struct {
	requests.RpcRequest
	ProjectName    string `position:"Query" name:"ProjectName"`
	CasShowType    int    `position:"Query" name:"CasShowType"`
	ShowDelService string `position:"Query" name:"ShowDelService"`
	CsbId          int64  `position:"Query" name:"CsbId"`
	Alias          string `position:"Query" name:"Alias"`
	ServiceName    string `position:"Query" name:"ServiceName"`
	PageNum        int    `position:"Query" name:"PageNum"`
}

func (req *FindServiceListRequest) Invoke(client *sdk.Client) (resp *FindServiceListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "FindServiceList", "CSB", "")
	resp = &FindServiceListResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindServiceListResponse struct {
	responses.BaseResponse
	Message   common.String
	Code      common.Integer
	RequestId common.String
	Data      FindServiceListData
}

type FindServiceListData struct {
	CurrentPage common.Integer
	PageNumber  common.Integer
	Total       common.Integer
	ServiceList FindServiceListServiceList
}

type FindServiceListService struct {
	Alias          common.String
	AllVisiable    bool
	CreateTime     common.Long
	CsbId          common.Long
	Description    common.String
	Id             common.Long
	InterfaceName  common.String
	ModifiedTime   common.Long
	OrderInfo      common.String
	OwnerId        common.String
	PrincipalName  common.String
	ProjectId      common.Long
	ProjectName    common.String
	Scope          common.String
	ServiceName    common.String
	ServiceVersion common.String
	SkipAuth       bool
	StatisticName  common.String
	Status         common.Integer
	UserId         common.String
	CasTargets     common.String
}

type FindServiceListServiceList []FindServiceListService

func (list *FindServiceListServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindServiceListService)
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
