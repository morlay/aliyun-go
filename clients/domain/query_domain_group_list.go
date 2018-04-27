package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDomainGroupListRequest struct {
	requests.RpcRequest
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	DomainGroupName   string `position:"Query" name:"DomainGroupName"`
	Lang              string `position:"Query" name:"Lang"`
	ShowDeletingGroup string `position:"Query" name:"ShowDeletingGroup"`
}

func (req *QueryDomainGroupListRequest) Invoke(client *sdk.Client) (resp *QueryDomainGroupListResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryDomainGroupList", "", "")
	resp = &QueryDomainGroupListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDomainGroupListResponse struct {
	responses.BaseResponse
	RequestId string
	Data      QueryDomainGroupListDomainGroupList
}

type QueryDomainGroupListDomainGroup struct {
	DomainGroupId     string
	DomainGroupName   string
	TotalNumber       int
	CreationDate      string
	ModificationDate  string
	DomainGroupStatus string
	BeingDeleted      bool
}

type QueryDomainGroupListDomainGroupList []QueryDomainGroupListDomainGroup

func (list *QueryDomainGroupListDomainGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDomainGroupListDomainGroup)
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
