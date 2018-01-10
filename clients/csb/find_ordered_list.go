package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindOrderedListRequest struct {
	requests.RpcRequest
	ProjectName  string `position:"Query" name:"ProjectName"`
	ShowDelOrder string `position:"Query" name:"ShowDelOrder"`
	CsbId        int64  `position:"Query" name:"CsbId"`
	Alias        string `position:"Query" name:"Alias"`
	ServiceName  string `position:"Query" name:"ServiceName"`
	PageNum      int    `position:"Query" name:"PageNum"`
	ServiceId    int64  `position:"Query" name:"ServiceId"`
	Status       string `position:"Query" name:"Status"`
}

func (req *FindOrderedListRequest) Invoke(client *sdk.Client) (resp *FindOrderedListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "FindOrderedList", "CSB", "")
	resp = &FindOrderedListResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindOrderedListResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      FindOrderedListData
}

type FindOrderedListData struct {
	CurrentPage int
	PageNumber  int
	OrderList   FindOrderedListOrderList
}

type FindOrderedListOrder struct {
	Alias                 string
	ProjectName           string
	ServiceName           string
	ServiceVersion        string
	OrderStatus           int
	AliveOrderCount       int
	GmtCreate             int64
	MaxRT                 int
	MinRT                 int
	ServiceId             string
	ServiceStatus         int
	ErrorTypeCatagoryList FindOrderedListErrorTypeCatagoryList
	Orders                FindOrderedListOrder1List
	Total                 FindOrderedListTotal
}

type FindOrderedListErrorTypeCatagory struct {
	Name     string
	Total    int64
	ErrorNum int64
}

type FindOrderedListOrder1 struct {
	Alias           string
	ApproveComments string
	CsbId           int64
	GmtCreate       int64
	GmtModified     int64
	GroupName       string
	Id              int64
	ProjectName     string
	ServiceId       int64
	ServiceName     string
	ServiceStatus   int
	ServiceVersion  string
	StatisticName   string
	Status          int
	UserId          string
	SlaInfo         FindOrderedListSlaInfo
	Total2          FindOrderedListTotal2
}

type FindOrderedListSlaInfo struct {
	Qph string
	Qps string
}

type FindOrderedListTotal2 struct {
	ErrorNum int
	Total    int
}

type FindOrderedListTotal struct {
	ErrorNum int
	Total    int
}

type FindOrderedListOrderList []FindOrderedListOrder

func (list *FindOrderedListOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindOrderedListOrder)
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

type FindOrderedListErrorTypeCatagoryList []FindOrderedListErrorTypeCatagory

func (list *FindOrderedListErrorTypeCatagoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindOrderedListErrorTypeCatagory)
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

type FindOrderedListOrder1List []FindOrderedListOrder1

func (list *FindOrderedListOrder1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindOrderedListOrder1)
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
