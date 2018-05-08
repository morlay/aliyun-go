package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      FindApprovalOrderListData
}

type FindApprovalOrderListData struct {
	CurrentPage common.Integer
	PageNumber  common.Integer
	OrderList   FindApprovalOrderListOrderList
}

type FindApprovalOrderListOrder struct {
	Alias               common.String
	CredentialGroupId   common.Long
	CsbId               common.Long
	GmtCreate           common.Long
	GmtModified         common.Long
	GroupName           common.String
	Id                  common.Long
	ProjectName         common.String
	ServiceId           common.Long
	ServiceName         common.String
	ServiceStatus       common.Integer
	ServiceVersion      common.String
	StatisticName       common.String
	Status              common.Integer
	StrictWhiteListJson common.String
	UserId              common.String
	UserName            common.String
	SlaInfo             FindApprovalOrderListSlaInfo
	Total               FindApprovalOrderListTotal
}

type FindApprovalOrderListSlaInfo struct {
	Qph common.Integer
	Qps common.Integer
}

type FindApprovalOrderListTotal struct {
	ErrorNum common.Integer
	Total    common.Integer
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
