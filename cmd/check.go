package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check site is available",
	Long:  `Checks site is available via proxy`,
	Run: func(cmd *cobra.Command, args []string) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				go check()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func check() {
    // Load CA cert
    caCert, err := os.ReadFile("CA-BrightData.crt")
    if err != nil {
        fmt.Printf("Error reading CA certificate: %s\n", err)
        os.Exit(1)
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    // Create TLS Config with CA
    tlsConfig := &tls.Config{
        RootCAs: caCertPool,
    }
    tlsConfig.BuildNameToCertificate()

    // Create a Transport with TLS config
    transport := &http.Transport{
        TLSClientConfig: tlsConfig,
    }

    // Proxy setup
    proxyStr := "http://{your_username}:{your_password}@brd.superproxy.io:22225"
    proxyURL, err := url.Parse(proxyStr)
    if err != nil {
        fmt.Println(err)
        return
    }
    transport.Proxy = http.ProxyURL(proxyURL)

    // Create client
    client := &http.Client{
        Transport: transport,
    }

    // Make request
    request, err := http.NewRequest("GET", "https://lumtest.com/myip.json", nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    response, err := client.Do(request)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(body))
}
