package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RefreshCdnDomainConfigsCacheRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Domains              string `position:"Query" name:"Domains"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *RefreshCdnDomainConfigsCacheRequest) Invoke(client *sdk.Client) (resp *RefreshCdnDomainConfigsCacheResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "RefreshCdnDomainConfigsCache", "mts", "")
	resp = &RefreshCdnDomainConfigsCacheResponse{}
	err = client.DoAction(req, resp)
	return
}

type RefreshCdnDomainConfigsCacheResponse struct {
	responses.BaseResponse
	RequestId     common.String
	SucessDomains RefreshCdnDomainConfigsCacheSucessDomainList
	FailedDomains RefreshCdnDomainConfigsCacheFailedDomainList
}

type RefreshCdnDomainConfigsCacheSucessDomainList []common.String

func (list *RefreshCdnDomainConfigsCacheSucessDomainList) UnmarshalJSON(data []byte) error {
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

type RefreshCdnDomainConfigsCacheFailedDomainList []common.String

func (list *RefreshCdnDomainConfigsCacheFailedDomainList) UnmarshalJSON(data []byte) error {
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
