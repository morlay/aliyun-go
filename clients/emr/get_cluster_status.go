package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetClusterStatusRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ItemType        string `position:"Query" name:"ItemType"`
	Interval        string `position:"Query" name:"Interval"`
	Id              string `position:"Query" name:"Id"`
}

func (req *GetClusterStatusRequest) Invoke(client *sdk.Client) (resp *GetClusterStatusResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetClusterStatus", "", "")
	resp = &GetClusterStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetClusterStatusResponse struct {
	responses.BaseResponse
	RequestId  common.String
	StatusList GetClusterStatusStatusList
}

type GetClusterStatusStatus struct {
	Name      common.String
	Legend    common.String
	Unit      common.String
	Series    GetClusterStatusSeriesInfoList
	LineNames GetClusterStatusLineNameList
	Times     GetClusterStatusTimeList
}

type GetClusterStatusSeriesInfo struct {
	SeriesItems GetClusterStatusSeriesItemList
}

type GetClusterStatusSeriesItem struct {
	SeriesValue common.Float
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

type GetClusterStatusLineNameList []common.String

func (list *GetClusterStatusLineNameList) UnmarshalJSON(data []byte) error {
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

type GetClusterStatusTimeList []common.String

func (list *GetClusterStatusTimeList) UnmarshalJSON(data []byte) error {
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
