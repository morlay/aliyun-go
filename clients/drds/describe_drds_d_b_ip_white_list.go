package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDrdsDBIpWhiteListRequest struct {
	requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	GroupName      string `position:"Query" name:"GroupName"`
}

func (req *DescribeDrdsDBIpWhiteListRequest) Invoke(client *sdk.Client) (resp *DescribeDrdsDBIpWhiteListResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsDBIpWhiteList", "", "")
	resp = &DescribeDrdsDBIpWhiteListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDrdsDBIpWhiteListResponse struct {
	responses.BaseResponse
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
