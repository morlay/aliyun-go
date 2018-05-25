package pvtz

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRequestGraphRequest struct {
	requests.RpcRequest
	VpcId          string `position:"Query" name:"VpcId"`
	UserClientIp   string `position:"Query" name:"UserClientIp"`
	ZoneId         string `position:"Query" name:"ZoneId"`
	Lang           string `position:"Query" name:"Lang"`
	StartTimestamp int64  `position:"Query" name:"StartTimestamp"`
	EndTimestamp   int64  `position:"Query" name:"EndTimestamp"`
}

func (req *DescribeRequestGraphRequest) Invoke(client *sdk.Client) (resp *DescribeRequestGraphResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "DescribeRequestGraph", "pvtz", "")
	resp = &DescribeRequestGraphResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRequestGraphResponse struct {
	responses.BaseResponse
	RequestId      common.String
	RequestDetails DescribeRequestGraphZoneRequestTopList
}

type DescribeRequestGraphZoneRequestTop struct {
	Time         common.String
	Timestamp    common.Long
	RequestCount common.Long
}

type DescribeRequestGraphZoneRequestTopList []DescribeRequestGraphZoneRequestTop

func (list *DescribeRequestGraphZoneRequestTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRequestGraphZoneRequestTop)
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
