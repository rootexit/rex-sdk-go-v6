package rexClient

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexConfig"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexSigner"
	"github.com/rootexit/rexLib/rexCtx"
	"github.com/rootexit/rexLib/rexHeaders"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	emptyStringSHA256 = `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
)

type Client struct {
	*http.Client
	conf   *rexConfig.Config
	signer *rexSigner.Signer
}

func NewClient(c *rexConfig.Config) *Client {
	httpClient := &http.Client{
		Timeout: time.Duration(c.Timeout) * time.Millisecond,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return &Client{
		Client: httpClient,
		conf:   c,
		signer: rexSigner.NewSigner(c),
	}
}

func (cli *Client) EasyNewRequest(ctx context.Context, svc string, relativePath string, method string, sendBody interface{}) ([]byte, error) {
	apiUrl := fmt.Sprintf("%s://%s%s%s", cli.conf.Protocol, cli.conf.Endpoint, "/rex/v5/apis", relativePath)
	if cli.conf.Debug {
		logx.Infof("打印一下请求的url :%s", apiUrl)
	}
	fn := cli.NewRequest(ctx, svc, apiUrl, method, nil, sendBody)
	res, err := fn()
	if err != nil {
		logx.Infof("打印一下请求错误 :%s", err)
		return nil, err
	}
	// 读取响应体
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logx.Infof("打印一下请求结果 :%v", res.Body)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (cli *Client) NewRequest(
	ctx context.Context, // 新增 context 参数
	svc string,
	url string, // URL
	method string, // HTTP 方法
	headers *map[string]string, // 请求头
	sendBody interface{}) func() (*http.Response, error) { // 返回闭包函数

	var (
		res *http.Response
		err error
	)

	// 创建一个 channel 来控制请求完成或超时
	c := make(chan struct{})
	go func() {
		defer close(c) // 保证 goroutine 退出时关闭 channel

		sendBodyJson := ""
		var payload []byte

		if sendBody != nil {
			// 将发送体序列化为 JSON
			sendBodyBt, marshalErr := json.Marshal(sendBody)
			if marshalErr != nil {
				err = marshalErr
				return
			}
			logx.Infof("send body: %s", string(sendBodyBt))
			sendBodyJson = string(sendBodyBt)
			payload = sendBodyBt
		}
		logx.Infof("payload: %s", string(payload))

		sendBodyIo := strings.NewReader(sendBodyJson)

		// 使用 context 控制请求
		req, reqErr := http.NewRequestWithContext(ctx, method, url, sendBodyIo)
		if reqErr != nil {
			err = reqErr
			return
		}

		// 设置请求头
		if headers != nil {
			for k, v := range *headers {
				req.Header.Set(k, v)
			}
		}

		if ctx.Value(rexCtx.CtxRequestId{}) != nil {
			logx.Infof("CtxRequestId: %s", ctx.Value(rexCtx.CtxRequestId{}))
			req.Header.Set(rexHeaders.HeaderXRequestIdFor, fmt.Sprintf("%v", ctx.Value(rexCtx.CtxRequestId{})))
		}

		req.Header.Set("Content-Type", "application/json")
		cli.signer.Sign(req, payload, svc, time.Now())
		res, err = cli.Client.Do(req)
		if err != nil {
			return
		}
	}()

	return func() (*http.Response, error) {
		select {
		case <-c: // 请求完成
			return res, err
		case <-ctx.Done(): // 请求超时或取消
			return nil, ctx.Err()
		}
	}
}
