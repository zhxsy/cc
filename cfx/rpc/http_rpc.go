package rpc

import (
	"cfx_server/cfx/util"
	"cfx_server/cfx/util/net"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type CfxRpcResult struct {
	JsonRpc string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Result  interface{} `json:"result"`
}

func Do(method, url string, body interface{}, options ...net.Option) (*CfxRpcResult, error) {
	req := ""
	switch body.(type) {
	case nil:
	case string:
		req = body.(string)
	default:
		req = util.MarshalToStringWithoutErr(body)
	}

	var httpResp *http.Response
	var bs []byte
	httpReq, err := http.NewRequest(method, url, strings.NewReader(req))
	if err != nil {
		log.Default().Printf("NewRequest err:[%v]\n", err)
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	for _, o := range options {
		o.ApplyOption(httpReq)
	}
	client := http.Client{}
	httpResp, err = client.Do(httpReq)
	if err != nil {
		log.Default().Printf("client.Do err:[%v]\n", err)
		return nil, err
	}
	if httpResp.Body != nil {
		if bs, err = ioutil.ReadAll(httpResp.Body); err != nil {
			return nil, err
		}
	}
	if httpResp.StatusCode >= 400 {
		return nil, errors.New(httpResp.Status)
	}

	if err != nil {
		return nil, err
	}

	result := &CfxRpcResult{}
	err = json.Unmarshal(bs, result)

	return result, err
}
