package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamsBlockListRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveStreamsBlockListRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamsBlockListResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamsBlockList", "", "")
	resp = &DescribeLiveStreamsBlockListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamsBlockListResponse struct {
	responses.BaseResponse
	RequestId  common.String
	DomainName common.String
	StreamUrls DescribeLiveStreamsBlockListStreamUrlList
}

type DescribeLiveStreamsBlockListStreamUrlList []common.String

func (list *DescribeLiveStreamsBlockListStreamUrlList) UnmarshalJSON(data []byte) error {
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
