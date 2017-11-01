package util

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

const ETCD_TIMEOUT = 3

var eventType = map[int32]string{
	0: "PUT",
	1: "DELETE",
}

type EtcdConn struct {
	Host string
	Port int
}

func (c *EtcdConn) Conn() (*clientv3.Client, error) {
	connString := fmt.Sprintf("%s:%d", c.Host, c.Port)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{connString},
		DialTimeout: ETCD_TIMEOUT * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// Get value of a key
func EtcdGet(clt *clientv3.Client, keyString string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ETCD_TIMEOUT * time.Second)
	resp, err := clt.Get(ctx, keyString)
	cancel()
	if err != nil {
		return "", err
	}

	var v string
	for _, kv := range resp.Kvs {
		v = string(kv.Value)
	}
	return v, nil
}

// Get values of all keys with prefix
func EtcdGetPrefix(clt *clientv3.Client, prefixString string) ([]map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ETCD_TIMEOUT * time.Second)
	resp, err := clt.Get(
		ctx,
		prefixString,
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
	)
	cancel()
	if err != nil {
		return nil, err
	}

	var result []map[string]string
	for _, kv := range resp.Kvs {
		mkv := make(map[string]string)
		mkv["key"] = string(kv.Key)
		mkv["value"] = string(kv.Value)
		result = append(result, mkv)
	}
	return result, nil
}

// Get all keys with prefix
func EtcdGetKeysOnly(clt *clientv3.Client, prefixString string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ETCD_TIMEOUT * time.Second)
	resp, err := clt.Get(
		ctx,
		prefixString,
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
	)
	cancel()
	if err != nil {
		return nil, err
	}

	var result []string
	for _, kv := range resp.Kvs {
		mk := string(kv.Key)
		result = append(result, mk)
	}
	return result, nil
}

// Put value to a key
func EtcdPut(clt *clientv3.Client, keyString string, valueString string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ETCD_TIMEOUT * time.Second)
	_, err := clt.Put(ctx, keyString, valueString)
	cancel()
	if err != nil {
		return false, err
	}
	return true, nil
}

// Put value to a key with TTL
func EtcdPutWithLeast(clt *clientv3.Client, keyString string, valueString string, t int64) (bool, error) {
	resp, e := clt.Grant(context.TODO(), t)
	if e != nil {
		return false, e
	}
	_, err := clt.Put(context.TODO(), keyString, valueString, clientv3.WithLease(resp.ID))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Watch a key once
func EtcdWatchOnce(clt *clientv3.Client, keyString string) (string, string, string) {
	var event string
	var key string
	var value string

	result := clt.Watch(context.Background(), keyString)

	for wresp := range result {
		for _, ev := range wresp.Events {
			event = eventType[int32(ev.Type)]
			key = string(ev.Kv.Key)
			value = string(ev.Kv.Value)
			break
		}
		break
	}
	return event, key, value
}

// Watch all keys with prefix
func EtcdWatchPrefixOnce(clt *clientv3.Client, prefixString string) (string, string, string) {
	var event string
	var key string
	var value string
	result := clt.Watch(context.Background(), prefixString, clientv3.WithPrefix())
	for wresp := range result {
		for _, ev := range wresp.Events {
			event = eventType[int32(ev.Type)]
			key = string(ev.Kv.Key)
			value = string(ev.Kv.Value)
			break
		}
		break
	}
	return event, key, value
}

// Delete a key
func EtcdDelete(clt *clientv3.Client, keyString string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ETCD_TIMEOUT * time.Second)
	_, err := clt.Delete(ctx, keyString)
	cancel()
	if err != nil {
		return false, err
	}
	return true, nil
}

// Delelte all keys with prefix
func EtcdDeletePrefix(clt *clientv3.Client, prefixString string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ETCD_TIMEOUT * time.Second)
	_, err := clt.Delete(ctx, prefixString, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return false, err
	}
	return true, nil
}
