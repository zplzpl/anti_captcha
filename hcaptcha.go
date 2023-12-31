package anti_captcha

import (
	"context"
	"net/url"
)

// HCaptchaResult is the api response from a hcaptcha task
type HCaptchaResult struct {
	GRecaptchaResponse string `json:"gRecaptchaResponse"`
}

// HCaptcha submits and retrieves a hcaptcha task
func (c *Client) HCaptcha(ctx context.Context, siteURL, siteKey, userAgent string, proxy *url.URL, opts ...OptionalValue) (result HCaptchaResult, err error) {
	var taskId int64
	data := map[string]interface{}{
		"type":       "HCaptchaTask",
		"websiteURL": siteURL,
		"websiteKey": siteKey,
		"userAgent":  userAgent,
	}

	for _, v := range opts {
		v(data)
	}

	err = addProxyInfo(proxy, data)
	if err != nil {
		return
	}

	taskId, err = c.createTask(ctx, data)
	if err != nil {
		return
	}

	err = c.fetchTask(ctx, taskId, &result)
	if err != nil {
		return
	}

	return
}

// HCaptchaProxyless submits and retrieves a hcaptcha task
func (c *Client) HCaptchaProxyless(ctx context.Context, siteURL, siteKey string, opts ...OptionalValue) (result HCaptchaResult, err error) {
	var taskId int64
	data := map[string]interface{}{
		"type":       "HCaptchaTaskProxyless",
		"websiteURL": siteURL,
		"websiteKey": siteKey,
	}

	for _, v := range opts {
		v(data)
	}

	taskId, err = c.createTask(ctx, data)
	if err != nil {
		return
	}

	err = c.fetchTask(ctx, taskId, &result)
	if err != nil {
		return
	}

	return
}
