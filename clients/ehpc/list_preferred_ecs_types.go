package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListPreferredEcsTypesRequest struct {
	requests.RpcRequest
	ZoneId string `position:"Query" name:"ZoneId"`
}

func (req *ListPreferredEcsTypesRequest) Invoke(client *sdk.Client) (resp *ListPreferredEcsTypesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListPreferredEcsTypes", "ehs", "")
	resp = &ListPreferredEcsTypesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPreferredEcsTypesResponse struct {
	responses.BaseResponse
	RequestId common.String
	Series    ListPreferredEcsTypesSeriesInfoList
}

type ListPreferredEcsTypesSeriesInfo struct {
	SeriesId   common.String
	SeriesName common.String
	Roles      ListPreferredEcsTypesRoles
}

type ListPreferredEcsTypesRoles struct {
	Manager ListPreferredEcsTypesManagerList
	Compute ListPreferredEcsTypesComputeList
	Login   ListPreferredEcsTypesLoginList
}

type ListPreferredEcsTypesSeriesInfoList []ListPreferredEcsTypesSeriesInfo

func (list *ListPreferredEcsTypesSeriesInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPreferredEcsTypesSeriesInfo)
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

type ListPreferredEcsTypesManagerList []common.String

func (list *ListPreferredEcsTypesManagerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type ListPreferredEcsTypesComputeList []common.String

func (list *ListPreferredEcsTypesComputeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type ListPreferredEcsTypesLoginList []common.String

func (list *ListPreferredEcsTypesLoginList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
