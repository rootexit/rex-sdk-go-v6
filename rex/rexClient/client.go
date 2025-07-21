package rexClient

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/rootexit/rex-sdk-go-v6/rex/rexConfig"
	"github.com/rootexit/rexLib/rexCrypto"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	emptyStringSHA256 = `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
)

type QxClient struct {
	*http.Client
	conf        *rexConfig.Config
	credentials *credentials.Credentials
}

func NewQxClient(c *rexConfig.Config) *QxClient {
	httpClient := &http.Client{
		Timeout: time.Duration(c.Timeout) * time.Millisecond,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	credentials := credentials.NewStaticCredentials(c.AccessKeyID, c.AccessKeySecret, "")
	return &QxClient{
		Client:      httpClient,
		conf:        c,
		credentials: credentials,
	}
}

func (cli *QxClient) EasyNewRequest(ctx context.Context, svc string, relativePath string, method string, sendBody interface{}) ([]byte, error) {
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

func (cli *QxClient) NewRequest(
	ctx context.Context, // 新增 context 参数
	svc string,
	url string, // URL
	method string, // HTTP 方法
	headers *map[string]string, // 请求头
	sendBody interface{}) func() (*http.Response, error) { // 返回闭包函数

	signer := v4.NewSigner(cli.credentials)

	var (
		res *http.Response
		err error
	)

	// 创建一个 channel 来控制请求完成或超时
	c := make(chan struct{})
	go func() {
		defer close(c) // 保证 goroutine 退出时关闭 channel

		sendBodyJson := ""

		if sendBody != nil {
			// 将发送体序列化为 JSON
			sendBodyBt, marshalErr := json.Marshal(sendBody)
			if marshalErr != nil {
				err = marshalErr
				return
			}
			sendBodyJson = string(sendBodyBt)
		}

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

		req.Header.Set("Content-Type", "application/json")
		if sendBody != nil {
			hexContentSha256, _ := rexCrypto.Sha256(sendBodyJson)
			req.Header.Add("X-Amz-Content-Sha256", hexContentSha256)
		} else {
			req.Header.Add("X-Amz-Content-Sha256", emptyStringSHA256)
		}
		signer.Sign(req, sendBodyIo, svc, cli.conf.Region, time.Now())
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
