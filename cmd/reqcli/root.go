package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/bryant-rh/reqcli/pkg"
	"github.com/imroc/req/v3"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var (
	version     string
	methodSlice = []string{"GET", "POST", "DELETE", "PUT", "UPDATE"}
	headerSlice []string
	headerMap   map[string]string
	urlStr      string
	request     string
	data        string
	resp        *req.Response
)

// versionString returns the version prefixed by 'v'
// or an empty string if no version has been populated by goreleaser.
// In this case, the --version flag will not be added by cobra.
func versionString() string {
	if len(version) == 0 {
		return ""
	}
	return "v" + version
}

func NewCmd() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:           "reqcli",
		Short:         "reqcli is a tool for detecting interface time-consuming",
		Version:       versionString(),
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Run: func(cmd *cobra.Command, args []string) {
			//校验参数
			err := Validate(cmd, args)
			if err != nil {
				klog.Fatal(err)
			}

			headerMap = make(map[string]string)

			if len(request) != 0 {
				for _, v := range headerSlice {
					s := strings.Split(v, ":")
					headerMap[s[0]] = s[1]

				}
			}

			client := pkg.NewReqClient()

			switch {
			case request == "GET":
				resp, err = client.R().
					SetHeaders(headerMap).
					Get(urlStr)
			case request == "POST":
				resp, err = client.R().
					SetHeaders(headerMap).
					SetBody(pkg.JsonToMap(data)).
					Post(urlStr)
			case request == "PUT":
				resp, err = client.R().
					SetHeaders(headerMap).
					SetBodyJsonString(data).
					Put(urlStr)
			case request == "DELETE":
				resp, err = client.R().
					SetHeaders(headerMap).
					SetBodyJsonString(data).
					Delete(urlStr)
			}

			if err != nil {
				klog.Fatal(err)
			}

			trace := resp.TraceInfo() // Use `resp.Request.TraceInfo()` to avoid unnecessary struct copy in production.

			fmt.Println("\n" + pkg.GreenColor(trace.Blame())) // Print out exactly where the http request is slowing down.
			fmt.Println(pkg.RedColor("-------------------------------"))
			fmt.Println(trace)

		},
	}

	rootCmd.PersistentFlags().StringVarP(&request, "request", "X", "GET", "指定method,支持 GET、POST、DELETE、PUT")
	rootCmd.PersistentFlags().StringSliceVarP(&headerSlice, "header", "H", []string{}, "Pass custom header(s) to server")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "HTTP POST data")
	return rootCmd
}

func Validate(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)

	} else if len(args) == 1 {
		urlStr = args[0]
		_, err := url.ParseRequestURI(urlStr)
		if err != nil {
			return err
		}

		if request != "" {
			if !pkg.ContainsInSlice(methodSlice, request) {
				//klog.Fatalf("不支持方法:[%s],只支持:%v ", request, methodSlice)
				return fmt.Errorf("不支持方法:[%s],只支持:%v ", request, methodSlice)

			}
		}
	} else if len(args) > 1 {
		return fmt.Errorf("url地址只支持一个")
	}
	return nil

}
