package httpclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type HttpClient interface {
	DiscoverAccounts(ctx context.Context, phone string, bankCode string) ([]string, error)
	SetMpin(ctx context.Context, accountId string, bankCode string, mpinEn string ) error 
	ChangeMpin(ctx context.Context, accountId string, bankCode string, oldMpinEn string, newMpinEn string ) error 
	GetBalance(ctx context.Context, accountId string, bankCode string, mpinEn string) (int64, error)
}

type BankServiceClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewBankServiceClient(url string) HttpClient {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, //must set to false in production
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &BankServiceClient{
		BaseURL: url,
		HTTPClient: &http.Client{
			Timeout:   10 * time.Second,
			Transport: transport,
		},
	}
}

func (c *BankServiceClient) DiscoverAccounts(ctx context.Context, phone string, bankCode string) ([]string, error) {
	body, _ := json.Marshal(map[string]interface{}{
		"phone":     phone,
		"bank_code": bankCode,
	})

	url := fmt.Sprintf("%s/account/discover", c.BaseURL)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-INTERNAL-API-KEY", os.Getenv("INTERNAL_API_KEY"))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []string{}, fmt.Errorf("bank returned error :%d", resp.StatusCode)
	}

	var result []string
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func (c *BankServiceClient) SetMpin(ctx context.Context, accountId string, bankCode string, mpinEn string ) error {
	body, _ := json.Marshal(map[string]interface{}{
		"account_id":accountId,
		"bank_code":bankCode,
		"mpin_encrypted":mpinEn,
	})

	url := fmt.Sprintf("%s/account/mpin", c.BaseURL)
	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-INTERNAL-API-KEY", os.Getenv("INTERNAL_API_KEY"))

	_, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	return nil 
}

func (c *BankServiceClient) ChangeMpin(ctx context.Context, accountId string, bankCode string, oldMpinEn string, newMpinEn string ) error {
	body, _ := json.Marshal(map[string]interface{}{
		"account_id":accountId,
		"bank_code":bankCode,
		"old_mpin_encrypted":oldMpinEn,
		"new_mpin_encrypted":newMpinEn,
	})

	url := fmt.Sprintf("%s/account/mpin", c.BaseURL)
	req, _ := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-INTERNAL-API-KEY", os.Getenv("INTERNAL_API_KEY"))

	_, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	return nil 
}

func (c *BankServiceClient) GetBalance(ctx context.Context, accountId string, bankCode string, mpinEn string) (int64, error){
	body, _ := json.Marshal(map[string]interface{}{
		"account_id":accountId,
		"bank_code":bankCode,
		"mpin_encrypted":mpinEn,
	})

	url := fmt.Sprintf("%s/account/balance", c.BaseURL)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-INTERNAL-API-KEY", os.Getenv("INTERNAL_API_KEY"))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0,  err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bank returned error :%d", resp.StatusCode)
	}

	var result int64
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}