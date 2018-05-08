package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FindOrderableListRequest struct {
	requests.RpcRequest
	ProjectName string `position:"Query" name:"ProjectName"`
	CsbId       int64  `position:"Query" name:"CsbId"`
	Alias       string `position:"Query" name:"Alias"`
	ServiceName string `position:"Query" name:"ServiceName"`
	PageNum     int    `position:"Query" name:"PageNum"`
}

func (req *FindOrderableListRequest) Invoke(client *sdk.Client) (resp *FindOrderableListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "FindOrderableList", "CSB", "")
	resp = &FindOrderableListResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindOrderableListResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      FindOrderableListData
}

type FindOrderableListData struct {
	CurrentPage common.Integer
	PageNumber  common.Integer
	ServiceList FindOrderableListServiceList
}

type FindOrderableListService struct {
	Alias          common.String
	AllVisiable    bool
	ApproveUserId  common.String
	CasTargets     common.String
	CreateTime     common.Long
	CsbId          common.Long
	Id             common.Long
	InterfaceName  common.String
	ModifiedTime   common.Long
	OwnerId        common.String
	PrincipalName  common.String
	ProjectId      common.String
	ProjectName    common.String
	Scope          common.String
	ServiceName    common.String
	ServiceVersion common.String
	SkipAuth       bool
	StatisticName  common.String
	Status         common.Integer
	UserId         common.String
}

type FindOrderableListServiceList []FindOrderableListService

func (list *FindOrderableListServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindOrderableListService)
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
