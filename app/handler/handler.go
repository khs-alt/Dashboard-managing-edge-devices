package handler

import (
	"camereye_backend_test/command"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Device struct {
	Addresses []string `json:"addresses"`
	Hostname  string   `json:"hostname"`
	LastSeen  string   `json:"lastSeen"`
}

type DevicesResponse struct {
	Devices []Device `json:"devices"`
}

// MainHandler redirects to the login page
func MainHandler(c *gin.Context) {
	redirectURL := "https://fianl-test.auth.ap-southeast-2.amazoncognito.com/login?client_id=6d0am4r9j19hifgtb2u3tucpt3&response_type=token&scope=aws.cognito.signin.user.admin&redirect_uri=https%3A%2F%2Fcg7jy1e6bi.execute-api.ap-southeast-2.amazonaws.com%2Fdemo%2Ftest_python"
	c.Redirect(http.StatusFound, redirectURL)
}

// LoginHandler handles the login process
func LoginHandler(c *gin.Context) {
	redirectURL := "https://fianl-test.auth.ap-southeast-2.amazoncognito.com/login?client_id=6d0am4r9j19hifgtb2u3tucpt3&response_type=token&scope=aws.cognito.signin.user.admin&redirect_uri=https%3A%2F%2Fcg7jy1e6bi.execute-api.ap-southeast-2.amazonaws.com%2Fdemo%2Ftest_python"
	c.Redirect(http.StatusFound, redirectURL)
}

// HomeHandler serves the home page
func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Home",
	})
}

// TailscaleDevicesHandler fetches devices from Tailscale API and renders them in a table
func TailscaleDevicesHandler(c *gin.Context) {
	url := "https://api.tailscale.com/api/v2/tailnet/hyungsubkim03@gmail.com/devices"

	// Create HTTP client
	client := &http.Client{}

	// Create a new HTTP request to the Tailscale API
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Add necessary headers
	req.Header.Add("Authorization", "Bearer tskey-api-kv7icz5CNTRL-gMgqnxnFgSKZvVvziW1ySKaqkpUMVSYQ7")

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute request"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	var devicesResponse DevicesResponse
	if err := json.Unmarshal(body, &devicesResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal response body"})
		return
	}

	// Construct the HTML response for the devices
	htmlResponse := `
					<table class='table-auto w-full mt-4'>
						<thead>
							<tr class='bg-gray-200'>
							<th class='px-4 py-2 w-1/4'>Hostname</th>
							<th class='px-4 py-2 w-1/4'>IP</th>
							<th class='px-4 py-2 w-1/4'>Status</th>
							<th class='px-4 py-2 w-1/4'>Actions</th>
							</tr>
						</thead>
					<tbody>`
	for _, device := range devicesResponse.Devices {
		// Generate a unique result target ID for each device
		resultTargetID := fmt.Sprintf("result-target-%s", device.Hostname)
		status := determineStatus(device.LastSeen)
		ip := "N/A"
		if len(device.Addresses) > 0 {
			ip = device.Addresses[0]
		}
		htmlResponse += fmt.Sprintf(`
							<tr>
								<td class='px-4 py-2'>%s</td>
								<td class='px-4 py-2'>%s</td>
								<td class='px-4 py-2'>%s</td>
								<td class='px-4 py-2'>
									<button hx-get='http://13.236.4.196/device-install-list?hostname=%s' hx-target='#%s' hx-trigger='click' class='bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded'>Get Package</button>
								</td>
							</tr>
							<tr>
								<td colspan="4" class="px-4 py-2 max-w-xs whitespace-normal overflow-auto" id="%s">package list</td>
							</tr>`, device.Hostname, ip, status, device.Hostname, resultTargetID, resultTargetID)
	}

	htmlResponse += "</tbody></table>"

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, htmlResponse)
	// Sends the response body to the client
}

// determineStatus calculates the online status based on the last seen time
func determineStatus(lastSeen string) string {
	lastSeenTime, err := time.Parse(time.RFC3339, lastSeen)
	if err != nil {
		return "Unknown"
	}
	if time.Since(lastSeenTime) <= 5*time.Minute {
		return "<span style='color: green;'>Online</span>"
	}
	return "<span style='color: red;'>Offline</span>"
}

// InstallListHandler handles the installation list request
func InstallListHandler(c *gin.Context) {
	device := c.Query("hostname")
	fmt.Println(device)
	lsOutput, err := command.ExecuteCommand("ls", []string{"-l"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute command"})
	}

	c.JSON(200, gin.H{
		"message":  lsOutput,
		"hostname": device,
	})
}
