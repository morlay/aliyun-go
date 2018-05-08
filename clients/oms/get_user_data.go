package oms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetUserDataRequest struct {
	requests.RpcRequest
	OwnerId      int64  `position:"Query" name:"OwnerId"`
	OwnerAccount string `position:"Query" name:"OwnerAccount"`
	ProductName  string `position:"Query" name:"ProductName"`
	DataType     string `position:"Query" name:"DataType"`
	StartTime    string `position:"Query" name:"StartTime"`
	EndTime      string `position:"Query" name:"EndTime"`
	NextToken    string `position:"Query" name:"NextToken"`
	MaxResult    int    `position:"Query" name:"MaxResult"`
}

func (req *GetUserDataRequest) Invoke(client *sdk.Client) (resp *GetUserDataResponse, err error) {
	req.InitWithApiInfo("Oms", "2015-02-12", "GetUserData", "", "")
	resp = &GetUserDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserDataResponse struct {
	responses.BaseResponse
	RequestId   common.String
	ProductName common.String
	DataType    common.String
	NextToken   common.String
	DataList    GetUserDataDataList
}

type GetUserDataData struct {
	DataItems GetUserDataDataItemList
}

type GetUserDataDataItem struct {
	Name  common.String
	Value common.String
}

type GetUserDataDataList []GetUserDataData

func (list *GetUserDataDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserDataData)
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

type GetUserDataDataItemList []GetUserDataDataItem

func (list *GetUserDataDataItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserDataDataItem)
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
