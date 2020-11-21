/*
 * Tencent is pleased to support the open source community by making TKEStack available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package prom

import (
	"context"
	"tkestack.io/kvass/pkg/api"

	promapi "github.com/prometheus/client_golang/api"
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// Client is a client to do prometheus API request
type Client struct {
	url string
	c   promapi.Client
}

// NewClient return an cli with url
func NewClient(url string) (*Client, error) {
	c, err := promapi.NewClient(promapi.Config{Address: url})
	if err != nil {
		return nil, err
	}
	return &Client{
		url: url,
		c:   c,
	}, nil
}

// RuntimeInfo return the current status of this shard, only return tManager targets if scrapingOnly is true,
// otherwise ,all target this cli discovered will be returned
func (c *Client) RuntimeInfo() (promv1.RuntimeinfoResult, error) {
	return promv1.NewAPI(c.c).Runtimeinfo(context.TODO())
}

// ConfigReload do Config reloading
func (c *Client) ConfigReload() error {
	url := c.url + "/-/reload"
	return api.Post(url, nil, nil)
}
