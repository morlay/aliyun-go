package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListEmrVersionsForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
	PageSize        int   `position:"Query" name:"PageSize"`
	PageNumber      int   `position:"Query" name:"PageNumber"`
}

func (req *ListEmrVersionsForAdminRequest) Invoke(client *sdk.Client) (resp *ListEmrVersionsForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListEmrVersionsForAdmin", "", "")
	resp = &ListEmrVersionsForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListEmrVersionsForAdminResponse struct {
	responses.BaseResponse
	RequestId               string
	PageNumber              int
	PageSize                int
	TotalCount              int
	MainVersionKeyValueList ListEmrVersionsForAdminMainVersionKeyValueList
}

type ListEmrVersionsForAdminMainVersionKeyValue struct {
	Id                   int
	Key                  string
	Value                string
	Status               int
	UtcCreateTimestamp   int64
	UtcModifiedTimestamp int64
}

type ListEmrVersionsForAdminMainVersionKeyValueList []ListEmrVersionsForAdminMainVersionKeyValue

func (list *ListEmrVersionsForAdminMainVersionKeyValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEmrVersionsForAdminMainVersionKeyValue)
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
