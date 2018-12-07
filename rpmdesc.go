package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/acrap/rpmdesc/src/rpmdesc"
	"github.com/urfave/cli"
)

func main() {
	rpmName := ""
	arch := ""
	os_arg := ""
	var licenseOutput, homepageOutput, noFoundPacketName, noFileListOutput bool = false, false, false, false

	app := cli.NewApp()
	app.Version = "0.1"
	app.Description = "rpmdesc is a tool that gives you full package name, homepage and list of files in RPM packet"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "nofilelist",
			Usage: "remove file list from the output",
		},

		cli.BoolFlag{
			Name:  "nopname",
			Usage: "remove picked packet name from the output",
		},
		cli.BoolFlag{
			Name:  "license",
			Usage: "add license info to output",
		},

		cli.BoolFlag{
			Name:  "homepage",
			Usage: "add homepage info to output",
		},

		cli.StringFlag{
			Name:  "name",
			Value: "",
			Usage: "name of rpm packet",
		},
		cli.StringFlag{
			Name:  "os",
			Value: "",
			Usage: "OS of rpm packet",
		},
		cli.StringFlag{
			Name:  "arch",
			Value: "",
			Usage: "architecture of rpm packet",
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(c.String("name")) == 0 {
			rpmName = enterName()
		} else {
			rpmName = c.String("name")
		}
		arch = c.String("arch")
		os_arg = c.String("os")
		licenseOutput = c.Bool("license")
		homepageOutput = c.Bool("homepage")
		noFoundPacketName = c.Bool("nopname")
		noFileListOutput = c.Bool("nofilelist")
		return nil
	}

	app.Run(os.Args)

	if len(rpmName) == 0 {
		// no rpm name specified
		return
	}

	res := local.GetSearchUrl(rpmName, os_arg, arch)
	descUrl := local.GetDescUrl(res)
	if len(descUrl) == 0 {
		log.Fatal("No rpm's found")
		return
	}

	if !noFileListOutput || noFoundPacketName {
		fmt.Println("Found packet:", descUrl[strings.LastIndex(descUrl, "/")+1:strings.LastIndex(descUrl, ".")])
	}

	doc, err := goquery.NewDocument("https://rpmfind.net/" + descUrl)

	if err != nil {
		log.Fatal(err)
	}

	if homepageOutput {
		fmt.Println(local.GetHomepage(doc))
	}

	if !noFileListOutput {
		fmt.Println(local.GetObjectsFromRpm(doc))
	}

	if licenseOutput {
		fmt.Print("License: ", local.GetLicenseFromRpm(doc))
	}

}

func enterName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rpm name: ")
	rpmName, _ := reader.ReadString('\n')
	rpmName = strings.Replace(rpmName, "\n", "", 1)
	fmt.Println(rpmName)
	return rpmName
}
