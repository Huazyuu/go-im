package open_login

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type GiteeConf struct {
	ClientID     string
	ClientSecret string
	Redirect     string
}

type GiteeInfo struct {
	Login     string `json:"login"`      // 用户名
	Name      string `json:"name"`       // 昵称
	AvatarURL string `json:"avatar_url"` // 头像
	ID        int    `json:"id"`         // 用户唯一标识（Gitee返回的是数字）
}

type GiteeLogin struct {
	clientID     string
	clientSecret string
	redirect     string
	code         string
	accessToken  string
}

func NewGiteeLogin(code string, conf GiteeConf) (giteeInfo GiteeInfo, err error) {
	if code == "" {
		return GiteeInfo{}, errors.New("code is empty")
	}
	giteeLogin := &GiteeLogin{
		clientID:     conf.ClientID,
		clientSecret: conf.ClientSecret,
		redirect:     conf.Redirect,
		code:         code,
	}
	err = giteeLogin.GetAccessToken()
	if err != nil {
		return giteeInfo, err
	}
	giteeInfo, err = giteeLogin.GetUserInfo()
	if err != nil {
		return giteeInfo, err
	}
	logx.Info(giteeInfo)
	return giteeInfo, nil
}

// GetAccessToken 获取访问令牌
func (g *GiteeLogin) GetAccessToken() error {
	client := &http.Client{Timeout: 30 * time.Second}

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", g.clientID)
	data.Set("client_secret", g.clientSecret)
	data.Set("code", g.code)
	data.Set("redirect_uri", g.redirect)

	req, err := http.NewRequest(
		"POST",
		"https://gitee.com/oauth/token",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return fmt.Errorf("create request failed: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		logx.Error(err)
		return fmt.Errorf("request failed: %v", errors.New("系统错误"))
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %v", err)
	}

	var result struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Error        string `json:"error"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("parse json failed: %v", err)
	}

	if result.AccessToken == "" {
		return fmt.Errorf("empty access token: %s", result.Error)
	}

	g.accessToken = result.AccessToken
	return nil
}

// GetUserInfo 获取用户信息
func (g *GiteeLogin) GetUserInfo() (GiteeInfo, error) {
	var info GiteeInfo
	client := &http.Client{Timeout: 10 * time.Second}

	// 构建带access_token的请求URL
	reqUrl := fmt.Sprintf("https://gitee.com/api/v5/user?access_token=%s", g.accessToken)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return info, fmt.Errorf("create request failed: %v", err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return info, fmt.Errorf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return info, fmt.Errorf("unexpected status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return info, fmt.Errorf("read body failed: %v", err)
	}

	if err := json.Unmarshal(body, &info); err != nil {
		return info, fmt.Errorf("parse json failed: %v", err)
	}

	// 如果需要字符串类型的ID可以转换
	// info.IDStr = strconv.FormatInt(info.ID, 10)

	return info, nil
}
