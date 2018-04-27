package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindApprovalOrderListRequest struct {
	requests.RpcRequest
	ProjectName string `position:"Query" name:"ProjectName"`
	Alias       string `position:"Query" name:"Alias"`
	ServiceName string `position:"Query" name:"ServiceName"`
	ServiceId   int64  `position:"Query" name:"ServiceId"`
	PageNum     int    `position:"Query" name:"PageNum"`
	OnlyPending string `position:"Query" name:"OnlyPending"`
}

func (req *FindApprovalOrderListRequest) Invoke(client *sdk.Client) (resp *FindApprovalOrderListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "FindApprovalOrderList", "CSB", "")
	resp = &FindApprovalOrderListResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindApprovalOrderListResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      FindApprovalOrderListData
}

type FindApprovalOrderListData struct {
	CurrentPage int
	PageNumber  int
	OrderList   FindApprovalOrderListOrderList
}

type FindApprovalOrderListOrder struct {
	Alias               string
	CredentialGroupId   int64
	CsbId               int64
	GmtCreate           int64
	GmtModified         int64
	GroupName           string
	Id                  int64
	ProjectName         string
	ServiceId           int64
	ServiceName         string
	ServiceStatus       int
	ServiceVersion      string
	StatisticName       string
	Status              int
	StrictWhiteListJson string
	UserId              string
	UserName            string
	SlaInfo             FindApprovalOrderListSlaInfo
	Total               FindApprovalOrderListTotal
}

type FindApprovalOrderListSlaInfo struct {
	Qph int
	Qps int
}

type FindApprovalOrderListTotal struct {
	ErrorNum int
	Total    int
}

type FindApprovalOrderListOrderList []FindApprovalOrderListOrder

func (list *FindApprovalOrderListOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindApprovalOrderListOrder)
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
