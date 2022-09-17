package apm

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func getLatestReleaseUrl(repoUrl string) string {
	// in: https://github.com/gocolly/colly
	// out: https://api.github.com/repos/gocolly/colly/releases/latest

	s := strings.Replace(repoUrl, "https://github.com/", "", -1)
	// gocolly/colly
	ss := strings.Split(s, "/")
	// gocolly colly

	repoAuthor := ss[0]
	repoName := ss[1]

	return fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoAuthor, repoName)
}

func getTarGzUrl(releaseLatestUrl string) (string, error) {
	resp, err := http.Get(releaseLatestUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var repoJson GithubReleaseLatestJson

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &repoJson)
	if err != nil {
		return "", err
	}

	return repoJson.TarballUrl, nil
}

func extractTarGz(gzipStream io.Reader) error {
	uncompressedStream := gzipStream

	tarReader := tar.NewReader(uncompressedStream)
	var header *tar.Header
	var err error
	var originamDirPath string
	var pkgName string
	for {
		header, err = tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("ExtractTarGz: Next() failed: %v", err)
		}

		nameHasPrefix := fmt.Sprintf("%s/packages/%s", os.Getenv("ARRIETTY_PATH"), header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(nameHasPrefix, 0755); err != nil {
				return fmt.Errorf("ExtractTarGz: Mkdir() failed: %v", err)
			}
			originamDirPath = nameHasPrefix
		case tar.TypeReg:
			outFile, err := os.Create(nameHasPrefix)
			if err != nil {
				return fmt.Errorf("ExtractTarGz: Create() failed: %v", err)
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				// outFile.Close error omitted as Copy error is more interesting at this point
				outFile.Close()
				return fmt.Errorf("ExtractTarGz: Copy() failed: %v", err)
			}

			if err := outFile.Close(); err != nil {
				return fmt.Errorf("ExtractTarGz: Close() failed: %v", err)
			}

			if strings.Contains(header.Name, "pkg.json") {
				j, err := os.ReadFile(nameHasPrefix)
				if err != nil {
					return fmt.Errorf("failed to read file: %v", err)
				}
				pkgInfo, err := UnmarshalPkgJson(j)
				if err != nil {
					return err
				}
				pkgName = pkgInfo.Name
			}
		case tar.TypeXGlobalHeader:
			continue

		default:
			return fmt.Errorf("ExtractTarGz: uknown type: %b in %s", header.Typeflag, header.Name)
		}
	}

	err = os.Rename(originamDirPath, fmt.Sprintf("%s/packages/%s", os.Getenv("ARRIETTY_PATH"), pkgName))
	if err != nil {
		return err
	}

	return nil
}

func InstallPkgFromGithub(repoUrl string) error {
	releaseUrl := getLatestReleaseUrl(repoUrl)
	targzUrl, err := getTarGzUrl(releaseUrl)
	if err != nil {
		return err
	}

	resp, err := http.Get(targzUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	err = extractTarGz(gzipReader)

	return err
}
