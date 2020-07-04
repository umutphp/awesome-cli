package fetcher

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const AWESOMEREPOURL = "https://github.com/sindresorhus/awesome/"
const AWESOMECACHEFOLDER = ".awesomecache"
const RAWGITHUBUSERCONTENT = "https://raw.githubusercontent.com"

func FetchAwsomeRootRepo() (string, error) {
	return FetchAwsomeRepo(AWESOMEREPOURL)
}

func FetchAwsomeRepo(repourl string) (string, error) {
	if !CacheFolderExists() {
		CreateCacheFolder()
	}

	cacheFile := GetCachePath(repourl)

	if CacheFileExists(cacheFile) && CacheFileUptoDate(cacheFile) {
		content, err := ioutil.ReadFile(cacheFile)

		if err != nil {
			log.Println(err)
			return "", err
		} else {
			return string(content), nil
		}
	}

	readmes := GetPossibleReadmeFileURLs(repourl)

	for _, rurl := range readmes {
		response, err := http.Get(rurl)

		if err != nil {
			log.Println(err)
			continue
		}

		if response.StatusCode == http.StatusNotFound {
			//log.Println(rurl, "gives 404.")
			continue
		}

		defer response.Body.Close()

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Println(err)
			return "", err
		}

		responseString := string(responseData)

		return SaveCache(cacheFile, responseString), nil
	}

	return "", nil
}

func GetPossibleReadmeFileURLs(repourl string) []string {
	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(repourl)
	if err != nil {
		log.Println(err)
	}

	if strings.Count(u.Host, "github.com") == 0 {
		return []string{}
	}

	prefix := RAWGITHUBUSERCONTENT + u.Path + "/master/"

	return []string{
		prefix + "README",
		prefix + "README.MD",
		prefix + "README.md",
		prefix + "readme",
		prefix + "readme.md",
		prefix + "readme.MD",
	}
}

func GetCachePath(url string) string {
	return GetCacheFolderPath() + string(os.PathSeparator) + CacheFileName(url)
}

func GetCacheFolderPath() string {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Println(err)
		return ""
	}

	return home + string(os.PathSeparator) + AWESOMECACHEFOLDER
}

func SaveCache(filename string, text string) string {
	file, err := os.Create(filename)

	if err != nil {
		log.Println(err)
		return ""
	}

	file.WriteString(text)

	return text
}

func CacheFileName(url string) string {
	hasher := md5.New()

	hasher.Write([]byte(url))

	return hex.EncodeToString(hasher.Sum(nil))
}

func CacheFolderExists() bool {
	info, err := os.Stat(GetCacheFolderPath())

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func CreateCacheFolder() bool {
	err := os.MkdirAll(GetCacheFolderPath(), 0755)

	if err != nil {
		return false
	}

	return true
}

func CacheFileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func CacheFileUptoDate(filename string) bool {
	info, _ := os.Stat(filename)
	modifiedtime := info.ModTime()

	return !IsOlderThanOneDay(modifiedtime)
}

func IsOlderThanOneDay(t time.Time) bool {
	return time.Now().Sub(t) > 24*time.Hour
}

func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) >= lenght {
			return str[0:lenght]
		}
	}
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)

	return err == nil && u.Scheme != "" && u.Host != ""
}
