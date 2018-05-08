package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      FindOrderedListData
}

type FindOrderedListData struct {
	CurrentPage common.Integer
	PageNumber  common.Integer
	OrderList   FindOrderedListOrderList
}

type FindOrderedListOrder struct {
	Alias                 common.String
	ProjectName           common.String
	ServiceName           common.String
	ServiceVersion        common.String
	OrderStatus           common.Integer
	AliveOrderCount       common.Integer
	GmtCreate             common.Long
	MaxRT                 common.Integer
	MinRT                 common.Integer
	ServiceId             common.String
	ServiceStatus         common.Integer
	ErrorTypeCatagoryList FindOrderedListErrorTypeCatagoryList
	Orders                FindOrderedListOrder1List
	Total                 FindOrderedListTotal
}

type FindOrderedListErrorTypeCatagory struct {
	Name     common.String
	Total    common.Long
	ErrorNum common.Long
}

type FindOrderedListOrder1 struct {
	Alias           common.String
	ApproveComments common.String
	CsbId           common.Long
	GmtCreate       common.Long
	GmtModified     common.Long
	GroupName       common.String
	Id              common.Long
	ProjectName     common.String
	ServiceId       common.Long
	ServiceName     common.String
	ServiceStatus   common.Integer
	ServiceVersion  common.String
	StatisticName   common.String
	Status          common.Integer
	UserId          common.String
	SlaInfo         FindOrderedListSlaInfo
	Total2          FindOrderedListTotal2
}

type FindOrderedListSlaInfo struct {
	Qph common.String
	Qps common.String
}

type FindOrderedListTotal2 struct {
	ErrorNum common.Integer
	Total    common.Integer
}

type FindOrderedListTotal struct {
	ErrorNum common.Integer
	Total    common.Integer
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
