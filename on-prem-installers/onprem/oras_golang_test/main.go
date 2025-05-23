package main

import (
	"context"
	"fmt"

	oras "oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/file"
	"oras.land/oras-go/v2/registry/remote"
)

const (
	RS_URL             = "registry-rs.edgeorchestration.intel.com"
	INSTALLERS_RS_PATH = "edge-orch/common/files"
	ARCHIVES_RS_PATH   = "edge-orch/common/files/orchestrator"
	ORCH_VERSION       = "v3.0.0"
	INSTALLERS_DIR     = "./installers"
	ARCHIVES_DIR       = "./repo_archives"
)

var (
	installerList = []string{
		"onprem-config-installer",
		"onprem-ke-installer",
		"onprem-argocd-installer",
		"onprem-gitea-installer",
		"onprem-orch-installer",
	}

	archiveList = []string{
		"onpremfull",
	}
)

func main() {
	downloadArtifacts(RS_URL, INSTALLERS_RS_PATH, ORCH_VERSION, INSTALLERS_DIR, installerList)
	downloadArtifacts(RS_URL, ARCHIVES_RS_PATH, ORCH_VERSION, ARCHIVES_DIR, archiveList)
}

func downloadArtifacts(registryUrl, registryPath, orchVersion, artifactDir string, artifactList []string) {
	ctx := context.Background()

	fileStore, err := file.New(artifactDir)
	if err != nil {
		panic(err) // TODO
	}
	defer fileStore.Close()

	for _, artifact := range artifactList {
		fmt.Println("downloading artifact: " + artifact)
		repo, err := remote.NewRepository(registryUrl + "/" + registryPath + "/" + artifact)
		if err != nil {
			panic(err) // TODO
		}

		manifestDescriptor, err := oras.Copy(ctx, repo, orchVersion, fileStore, orchVersion, oras.DefaultCopyOptions)
		if err != nil {
			panic(err) // TODO
		}
		fmt.Println("manifest descriptor:", manifestDescriptor)
	}
}
