# Tailscale Device Manager Web Service

This web service provides a user interface for viewing information about Linux-based devices registered on Tailscale and fetching the list of installed applications on each device.

## Technologies Used

- **Language & Framework**: The backend is developed in Go, utilizing the Gin web framework.
- **User Interface (UI)**: The front end uses HTMX and Tailwind CSS for a dynamic and responsive design.
- **Deployment**: The application is deployed on AWS EC2 instances.
- **Authentication**: Utilizes AWS Cognito User Pools for authentication.

## Current Issues

1. While the service successfully retrieves information from devices connected via Tailscale, it fails to fetch the list of installed applications from each device.

2. The authentication system is not fully implemented. While login functionality works, the restriction of access to authenticated users only is incomplete.

## Installation and Running

The program requires Go version 1.22.

To set up the system for deployment:
1. Install Go on the deployment system.
2. Clone the repository using `git clone`.
3. Execute `go run main.go` to run the program.

After registering your devices with Tailscale, update the `TailscaleDevicesHandler` URL in the format:
`https://api.tailscale.com/api/v2/tailnet/{user_id}/devices`, where `user_id` is your Tailscale registered ID.

## Getting Started

You can login from the home page either by registering a new account or using the provided demo credentials.

- **Demo ID**: vekeso9433@flexvio.com
- **Demo Password**: abcd1234!

This service aims to simplify managing Tailscale-registered Linux devices and enhance the visibility of installed applications across your network.

## Work Process

This [development process document](https://docs.google.com/document/d/1goaHwZMRKISYeXH7IYmhDYPXu867JopLfzqVepu04jc/edit) shows current development process and the problem