package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamsBlockListRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int    `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int    `position:"Query" name:"PageNum"`
}

func (req *DescribeLiveStreamsBlockListRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamsBlockListResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsBlockList", "live", "")
	resp = &DescribeLiveStreamsBlockListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamsBlockListResponse struct {
	responses.BaseResponse
	RequestId  string
	DomainName string
	PageNum    int
	PageSize   int
	TotalNum   int
	TotalPage  int
	StreamUrls DescribeLiveStreamsBlockListStreamUrlList
}

type DescribeLiveStreamsBlockListStreamUrlList []string

func (list *DescribeLiveStreamsBlockListStreamUrlList) UnmarshalJSON(data []byte) error {
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
