package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryPageByApplyIdRequest struct {
	ApplyId     int64 `position:"Query" name:"ApplyId"`
	PageSize    int   `position:"Query" name:"PageSize"`
	CurrentPage int   `position:"Query" name:"CurrentPage"`
}

func (r QueryPageByApplyIdRequest) Invoke(client *sdk.Client) (response *QueryPageByApplyIdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryPageByApplyIdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "QueryPageByApplyId", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryPageByApplyIdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryPageByApplyIdResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryPageByApplyIdResponse struct {
	RequestId       string
	Success         bool
	ErrorMessage    string
	PageSize        int
	Page            int
	PageCount       int
	Total           int
	ApplyDeviceList QueryPageByApplyIdApplyDeviceInfoList
}

type QueryPageByApplyIdApplyDeviceInfo struct {
	DeviceId     string
	DeviceName   string
	DeviceSecret string
}

type QueryPageByApplyIdApplyDeviceInfoList []QueryPageByApplyIdApplyDeviceInfo

func (list *QueryPageByApplyIdApplyDeviceInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPageByApplyIdApplyDeviceInfo)
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
