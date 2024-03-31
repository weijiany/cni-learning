package etcdutils

import (
	"context"
	"go.etcd.io/etcd/client/pkg/v3/transport"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

type Client struct {
	etcdClient *clientv3.Client
	basePath   string
}

func GetClient() *Client {
	tlsInfo := transport.TLSInfo{
		CertFile:      "/etc/kubernetes/pki/etcd/healthcheck-client.crt",
		KeyFile:       "/etc/kubernetes/pki/etcd/healthcheck-client.key",
		TrustedCAFile: "/etc/kubernetes/pki/etcd/ca.crt",
	}
	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		panic(err)
	}

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		TLS:         tlsConfig,
		DialTimeout: 1 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	return &Client{
		etcdClient: etcdClient,
		basePath:   "/my-cni/ipam/",
	}
}

func (c *Client) key(suffix []string) string {
	return c.basePath + strings.Join(suffix, "/")
}

func (c *Client) Get(key ...string) string {
	resp, err := c.etcdClient.Get(context.TODO(), c.key(key))
	if err != nil {
		panic(err)
	}
	if len(resp.Kvs) == 0 {
		return ""
	}
	return string(resp.Kvs[0].Value)
}

func (c *Client) Put(value string, key ...string) error {
	_, err := c.etcdClient.Put(context.TODO(), c.key(key), value)
	return err
}
