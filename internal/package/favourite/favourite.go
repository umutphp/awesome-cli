package favourite

import (
	"math/rand"
    "time"
    "os"
    "encoding/gob"

	"github.com/umutphp/awesome-cli/internal/package/fetcher"
)

type Favourite struct {
	Name string
	Children map[string]Favourite
}

func New(name string) Favourite {
	return Favourite{
		Name: name,
		Children: map[string]Favourite{},
	}
}

func (f *Favourite) Add(fav Favourite) {
	_, ok := f.Children[fav.Name]

	if ok == false {
		f.Children[fav.Name] = fav
	}
}

func (f *Favourite) GetName() string {
	return f.Name
}

func (f *Favourite) GetChild(name string) Favourite {
	return f.Children[name]
}

func (f *Favourite) GetRandom() Favourite {
	rand.Seed(time.Now().UTC().UnixNano())
	rint := rand.Intn(len(f.Children))

	i := 0
	for _,fav := range f.Children  {
		if i == rint {
			return fav
		}
		i++
	}

	return f.GetRandom()
}

func (f *Favourite) SaveCache() {
	filename := fetcher.GetCachePath(f.Name)

	encodeFile, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	// Since this is a binary format large parts of it will be unreadable
	encoder := gob.NewEncoder(encodeFile)

	// Write to the file
	if err := encoder.Encode(f); err != nil {
		panic(err)
	}

	encodeFile.Close()
}

func NewFromCache(cachename string) Favourite {
	filename   := fetcher.GetCachePath(cachename)
	_, err     := os.Stat(filename)
	favourites := Favourite{
		Name: cachename,
		Children: map[string]Favourite{},
	}

    if os.IsNotExist(err) {
        return favourites
    }
	
	decodeFile, err := os.Open(filename)
	
	if err != nil {
		panic(err)
	}
	
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)

	decoder.Decode(&favourites)

	return favourites
}
