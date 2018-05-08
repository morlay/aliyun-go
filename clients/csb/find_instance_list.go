package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FindInstanceListRequest struct {
	requests.RpcRequest
	SearchTxt string `position:"Query" name:"SearchTxt"`
	CsbId     int64  `position:"Query" name:"CsbId"`
	PageNum   int    `position:"Query" name:"PageNum"`
	Status    int    `position:"Query" name:"Status"`
}

func (req *FindInstanceListRequest) Invoke(client *sdk.Client) (resp *FindInstanceListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "FindInstanceList", "CSB", "")
	resp = &FindInstanceListResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindInstanceListResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      FindInstanceListData
}

type FindInstanceListData struct {
	CurrentPage common.Integer
	PageNumber  common.Integer
	ItemList    FindInstanceListItemList
}

type FindInstanceListItem struct {
	Description      common.String
	FrontStatus      common.String
	GmtCreate        common.Long
	GmtModified      common.Long
	Id               common.Long
	InstanceCategory common.Integer
	Name             common.String
	StatusCode       common.Integer
	Visible          bool
	VpcName          common.String
}

type FindInstanceListItemList []FindInstanceListItem

func (list *FindInstanceListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindInstanceListItem)
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
