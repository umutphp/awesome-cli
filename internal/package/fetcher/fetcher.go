package fetcher

import (
	"fmt"
	"net/http"
	"io/ioutil"
    "os"
    "time"
    "crypto/md5"
    "encoding/hex"
)

const AWESOMEREADME = "https://raw.githubusercontent.com/sindresorhus/awesome/master/readme.md"
const AWESOMECACHEFOLDER = ".awsomecache"

func FetchAwsomeRootRepo() (string, error) {
	return FetchAwsomeRepo(AWESOMEREADME)
}

func FetchAwsomeRepo(repourl string) (string, error) {
	if !CacheFolderExists() {
		CreateCacheFolder()
	}

    cacheFile := GetCachePath(repourl)

    if CacheFileExists(cacheFile) && CacheFileUptoDate(cacheFile) {
        content, err := ioutil.ReadFile(cacheFile)
        
        if err != nil {
            fmt.Println(err)
            return "", err
        } else {
            return string(content), nil
        }
    }

    response, err := http.Get(repourl)
    
    if err != nil {
        fmt.Println(err)
    }

    defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)

    if err != nil {
        fmt.Println(err)
        return "", err
    } 

    responseString := string(responseData)

    return SaveCache(cacheFile, responseString), nil
}

func GetCachePath(url string) string {
    return GetCacheFolderPath() + string(os.PathSeparator) + CacheFileName(url)
}

func GetCacheFolderPath() string {
	home, err := os.UserHomeDir()

    if err != nil {
        fmt.Println(err)
        return ""
    }

    return home + string(os.PathSeparator) + AWESOMECACHEFOLDER
}

func SaveCache(filename string, text string) string {
	file, err := os.Create(filename)

    if err != nil {
        fmt.Println(err)
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
    info, _      := os.Stat(filename)
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
