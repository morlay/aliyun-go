package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetUserOutputStatisticInfoRequest struct {
	requests.RpcRequest
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
}

func (req *GetUserOutputStatisticInfoRequest) Invoke(client *sdk.Client) (resp *GetUserOutputStatisticInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetUserOutputStatisticInfo", "", "")
	resp = &GetUserOutputStatisticInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserOutputStatisticInfoResponse struct {
	responses.BaseResponse
	RequestId      common.String
	UserOutputList GetUserOutputStatisticInfoClusterStatUserOutputList
}

type GetUserOutputStatisticInfoClusterStatUserOutput struct {
	User        common.String
	BytesOutput common.Long
}

type GetUserOutputStatisticInfoClusterStatUserOutputList []GetUserOutputStatisticInfoClusterStatUserOutput

func (list *GetUserOutputStatisticInfoClusterStatUserOutputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserOutputStatisticInfoClusterStatUserOutput)
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
