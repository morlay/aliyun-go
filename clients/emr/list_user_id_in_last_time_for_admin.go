package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListUserIdInLastTimeForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
	PageSize        int   `position:"Query" name:"PageSize"`
	StartTime       int64 `position:"Query" name:"StartTime"`
	PageNumber      int   `position:"Query" name:"PageNumber"`
}

func (req *ListUserIdInLastTimeForAdminRequest) Invoke(client *sdk.Client) (resp *ListUserIdInLastTimeForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListUserIdInLastTimeForAdmin", "", "")
	resp = &ListUserIdInLastTimeForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListUserIdInLastTimeForAdminResponse struct {
	responses.BaseResponse
	RequestId  string
	PageNumber int
	PageSize   int
	TotalCount int
	UserIds    ListUserIdInLastTimeForAdminUserIdList
}

type ListUserIdInLastTimeForAdminUserIdList []string

func (list *ListUserIdInLastTimeForAdminUserIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
