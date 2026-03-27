package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"go.ngs.io/dropbox-mcp-server/internal/config"
)

func main() {
	os.Setenv("DROPBOX_MCP_CONFIG_PATH", "/Users/roytemple/.dropbox-mcp-server/config-gps.json")
	cfg, err := config.Load()
	if err != nil { fmt.Println("load error:", err); return }
	fmt.Println("token prefix:", cfg.AccessToken[:20])
	
	dbxConfig := dropbox.Config{Token: cfg.AccessToken}
	client := files.New(dbxConfig)
	
	data := []byte("SDK upload test 2026-03-27")
	arg := files.NewUploadArg("/GPS-Media/test-sdk-upload.txt")
	arg.Mode = &files.WriteMode{Tagged: dropbox.Tagged{Tag: "add"}}
	arg.Autorename = true
	
	result, err := client.Upload(arg, bytes.NewReader(data))
	if err != nil { fmt.Println("upload error:", err); return }
	fmt.Println("uploaded:", result.PathDisplay)
}
