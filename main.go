package main

import (
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	cli, err := client.NewEnvClient()
	ctx := context.Background()
	if err != nil {
		panic(err)
	}

	// list all images
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	// loop over images and pull the image with the corresponding RepoTags
	for _, image := range images {
		go doPull(image)
	}
}

func doPull(image cli.Image) {
	imageString := image.RepoTags[0]
	out, err := cli.ImagePull(ctx, imageString, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
}

