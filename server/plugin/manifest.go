// This file is automatically generated. Do not modify it manually.

package plugin

import (
	"encoding/json"
	"strings"

	"github.com/mattermost/mattermost/server/public/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "com.darklord.plugin-google-drive",
  "name": "Google Drive Plugin",
  "description": "This plugin allows you to integrate Google Drive to your Mattermost instance.",
  "homepage_url": "https://github.com/darkLord/mattermost-plugin-google-drive",
  "support_url": "https://github.com/mattermost/mattermost-plugin-google-drive/issues",
  "icon_path": "assets/icon.svg",
  "version": "0.0.0+eebdc03",
  "min_server_version": "6.2.1",
  "server": {
    "executables": {
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "darwin-arm64": "server/dist/plugin-darwin-arm64",
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "linux-arm64": "server/dist/plugin-linux-arm64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "settings_schema": {
    "header": "The Google Drive plugin for Mattermost allows users to create, share files in Google drive and receive notifications for shared files and comments on files to stay up-to-date. \n \n Instructions for setup are [available here](https://github.com/mattermost-community/mattermost-plugin-google-drive#configuration).",
    "footer": "* To report an issue, make a suggestion or a contribution, [check the repository](https://github.com/darkLord19/mattermost-plugin-google-drive).",
    "settings": [
      {
        "key": "GoogleOAuthClientID",
        "display_name": "Google OAuth Client ID:",
        "type": "text",
        "help_text": "The client ID for the OAuth app registered with Google.",
        "placeholder": "",
        "default": null,
        "hosting": ""
      },
      {
        "key": "GoogleOAuthClientSecret",
        "display_name": "Google OAuth Client Secret:",
        "type": "text",
        "help_text": "The client secret for the OAuth app registered with Google.",
        "placeholder": "",
        "default": null,
        "hosting": ""
      },
      {
        "key": "EncryptionKey",
        "display_name": "At Rest Encryption Key:",
        "type": "generated",
        "help_text": "The AES encryption key used to encrypt stored access tokens.",
        "placeholder": "",
        "default": null,
        "hosting": ""
      }
    ]
  }
}
`

func init() {
	_ = json.NewDecoder(strings.NewReader(manifestStr)).Decode(&manifest)
}
