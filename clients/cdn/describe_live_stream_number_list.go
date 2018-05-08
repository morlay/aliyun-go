package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamNumberListRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveStreamNumberListRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamNumberListResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamNumberList", "", "")
	resp = &DescribeLiveStreamNumberListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamNumberListResponse struct {
	responses.BaseResponse
	RequestId         common.String
	DomainName        common.String
	StreamNumberInfos DescribeLiveStreamNumberListStreamNumberInfoList
}

type DescribeLiveStreamNumberListStreamNumberInfo struct {
	Number common.Integer
	Time   common.String
}

type DescribeLiveStreamNumberListStreamNumberInfoList []DescribeLiveStreamNumberListStreamNumberInfo

func (list *DescribeLiveStreamNumberListStreamNumberInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamNumberListStreamNumberInfo)
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
