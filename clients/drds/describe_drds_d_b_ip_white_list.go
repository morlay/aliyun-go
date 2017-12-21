package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDrdsDBIpWhiteListRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	GroupName      string `position:"Query" name:"GroupName"`
}

func (r DescribeDrdsDBIpWhiteListRequest) Invoke(client *sdk.Client) (response *DescribeDrdsDBIpWhiteListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDrdsDBIpWhiteListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsDBIpWhiteList", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDrdsDBIpWhiteListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDrdsDBIpWhiteListResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDrdsDBIpWhiteListResponse struct {
	RequestId string
	Success   bool
	Data      DescribeDrdsDBIpWhiteListData
}

type DescribeDrdsDBIpWhiteListData struct {
	IpWhiteList DescribeDrdsDBIpWhiteListIpWhiteListList
}

type DescribeDrdsDBIpWhiteListIpWhiteListList []string

func (list *DescribeDrdsDBIpWhiteListIpWhiteListList) UnmarshalJSON(data []byte) error {
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
