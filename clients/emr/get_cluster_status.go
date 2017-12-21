package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetClusterStatusRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ItemType        string `position:"Query" name:"ItemType"`
	Interval        string `position:"Query" name:"Interval"`
}

func (r GetClusterStatusRequest) Invoke(client *sdk.Client) (response *GetClusterStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetClusterStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "GetClusterStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		GetClusterStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetClusterStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetClusterStatusResponse struct {
	RequestId  string
	StatusList GetClusterStatusStatusList
}

type GetClusterStatusStatus struct {
	Name      string
	Legend    string
	Unit      string
	Series    GetClusterStatusSeriesInfoList
	LineNames GetClusterStatusLineNameList
	Times     GetClusterStatusTimeList
}

type GetClusterStatusSeriesInfo struct {
	SeriesItems GetClusterStatusSeriesItemList
}

type GetClusterStatusSeriesItem struct {
	SeriesValue float32
}

type GetClusterStatusStatusList []GetClusterStatusStatus

func (list *GetClusterStatusStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterStatusStatus)
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

type GetClusterStatusSeriesInfoList []GetClusterStatusSeriesInfo

func (list *GetClusterStatusSeriesInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterStatusSeriesInfo)
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

type GetClusterStatusLineNameList []string

func (list *GetClusterStatusLineNameList) UnmarshalJSON(data []byte) error {
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

type GetClusterStatusTimeList []string

func (list *GetClusterStatusTimeList) UnmarshalJSON(data []byte) error {
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

type GetClusterStatusSeriesItemList []GetClusterStatusSeriesItem

func (list *GetClusterStatusSeriesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterStatusSeriesItem)
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
