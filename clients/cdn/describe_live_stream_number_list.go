package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamNumberListRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveStreamNumberListRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamNumberListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamNumberListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamNumberList", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamNumberListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamNumberListResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamNumberListResponse struct {
	RequestId         string
	DomainName        string
	StreamNumberInfos DescribeLiveStreamNumberListStreamNumberInfoList
}

type DescribeLiveStreamNumberListStreamNumberInfo struct {
	Number int
	Time   string
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
