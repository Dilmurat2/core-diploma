package ozon

import (
	"context"
	"fmt"
	"github.com/playwright-community/playwright-go"
)

type Client interface {
	GetProducts(context.Context, GetProductsRequest) (string, error)
}

type client struct {
	pw      *playwright.Playwright
	browser playwright.Browser
	context playwright.BrowserContext
	config  Config
}

func NewClient(config Config) (Client, error) {

	if err := setupBrowser(); err != nil {
		return nil, fmt.Errorf("could not setup browser: %v", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("could not start playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
		Args: []string{
			"--disable-gpu",
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"--window-size=800,600",
			"--disable-extensions",
			"--disable-blink-features=AutomationControlled",
			"--disable-software-rasterizer",
			"--disable-default-apps",
			"--no-zygote",
			"--disable-background-networking",
			"--disable-background-timer-throttling",
			"--disable-backgrounding-occluded-windows",
			"--disable-breakpad",
			"--disable-component-extensions-with-background-pages",
			"--disable-features=TranslateUI,BlinkGenPropertyTrees",
			"--disable-ipc-flooding-protection",
			"--disable-renderer-backgrounding",
			"--enable-features=NetworkService,NetworkServiceInProcess",
		},
	})

	if err != nil {
		return nil, fmt.Errorf("could not launch browser: %v", err)
	}

	pwContext, err := browser.NewContext(playwright.BrowserNewContextOptions{
		JavaScriptEnabled: playwright.Bool(true),
		NoViewport:        playwright.Bool(true),
		UserAgent:         playwright.String("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"),
	})

	if err != nil {
		return nil, fmt.Errorf("could not create context: %v", err)
	}

	cookies := make([]playwright.OptionalCookie, 0, len(config.Cookies))

	for key, value := range config.Cookies {
		cookies = append(cookies, playwright.OptionalCookie{
			Name:   key,
			Value:  value,
			Domain: playwright.String(config.Site),
			Path:   playwright.String("/"),
		})
	}

	if err := pwContext.AddCookies(cookies); err != nil {
		return nil, fmt.Errorf("could not add cookies: %v", err)
	}

	return &client{
		pw:      pw,
		browser: browser,
		context: pwContext,
		config:  config,
	}, nil
}

func (s *client) GetProducts(_ context.Context, req GetProductsRequest) (string, error) {
	page, err := s.context.NewPage()
	if err != nil {
		return "", fmt.Errorf("could not create page: %v", err)
	}

	url := s.buildSearchUrl(req.Query, req.Sort)

	if _, err := page.Goto(url); err != nil {
		return "", fmt.Errorf("could not navigate to page: %v", err)
	}

	if _, err := page.WaitForSelector(s.config.Selector, playwright.PageWaitForSelectorOptions{
		State: playwright.WaitForSelectorStateAttached,
	}); err != nil {
		return "", fmt.Errorf("timeout waiting for search results: %v", err)
	}
	content, err := page.Content()
	if err != nil {
		return "", fmt.Errorf("could not get page content: %v", err)
	}

	if err := page.Close(); err != nil {
		return "", fmt.Errorf("could not close page: %v", err)
	}

	return content, nil
}

func (s *client) buildSearchUrl(query, sorting string) string {
	return fmt.Sprintf("%s%s?text=%s&sorting=%s", s.config.BaseUrl, s.config.SearchPath, query, sorting)
}

func setupBrowser() error {
	if err := playwright.Install(); err != nil {
		return err
	}
	return nil
}
