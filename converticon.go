package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

var (
	iconSizes = map[string]int{
		"drawable-xxhdpi": 144,
		"drawable-xhdpi":  96,
		"drawable-hdpi":   72,
		"drawable-mdpi":   48,
	}
	androidLauncher = "ic_launcher.png"
	innerPath       = "app/src/main/res"
)

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isDir(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.IsDir()
}

func main() {
	app := cli.NewApp()
	app.Name = "converticon"
	app.Version = "0.1"
	app.Usage = "convert android icon sizes for playstore"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path, p",
			Usage: "specify project path",
		},
		cli.StringFlag{
			Name:  "image, i",
			Usage: "specify image icon path",
		},
	}

	app.Action = func(c *cli.Context) {
		projectPath := c.String("path")
		if projectPath == "" {
			projectPath, _ = os.Getwd()
		}
		if ok := isExist(projectPath); !ok {
			log.Fatalf("not found: %s", projectPath)
		}

		imagePath := c.String("image")
		if imagePath == "" {
			imagePath = androidLauncher
		}
		if ok := isDir(imagePath); ok {
			imagePath = filepath.Join(imagePath, androidLauncher)
		}
		if ok := isExist(imagePath); !ok {
			log.Fatalf("not found: %s", imagePath)
		}

		for p, size := range iconSizes {
			path := filepath.Join(projectPath, innerPath, p, androidLauncher)
			err := resizeCmd(imagePath, path, size)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("generate %s\n", path)
		}
	}

	app.Run(os.Args)
}
