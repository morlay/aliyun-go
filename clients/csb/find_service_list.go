package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindServiceListRequest struct {
	requests.RpcRequest
	ProjectName    string `position:"Query" name:"ProjectName"`
	ShowDelService string `position:"Query" name:"ShowDelService"`
	CasShowType    int    `position:"Query" name:"CasShowType"`
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
	Message   string
	Code      int
	RequestId string
	Data      FindServiceListData
}

type FindServiceListData struct {
	CurrentPage int
	PageNumber  int
	Total       int
	ServiceList FindServiceListServiceList
}

type FindServiceListService struct {
	Alias          string
	AllVisiable    bool
	CreateTime     int64
	CsbId          int64
	Description    string
	Id             int64
	InterfaceName  string
	ModifiedTime   int64
	OrderInfo      string
	OwnerId        string
	PrincipalName  string
	ProjectId      int64
	ProjectName    string
	Scope          string
	ServiceName    string
	ServiceVersion string
	SkipAuth       bool
	StatisticName  string
	Status         int
	UserId         string
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
