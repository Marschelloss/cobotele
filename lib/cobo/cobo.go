package cobo

import (
	"time"
)

// ComicItem is a struct which defines a fetched comic.
type ComicItem struct {
	Title      string
	PictureURL string
	SiteName   string
	Date       time.Time
}

// comicSiteTraits defines a comic-site. The number of
// fetched comics in one rune is declared in ComicNum.
// This is importand for calculating the channel size.
type comicSiteTraits struct {
	SiteName string
	SiteURL  string
	ComicNum int
}

// Start launches the fetching process.
func Start() {
	comicSitesList := createComicSiteList()
	comicNumSum := sumComicNum(comicSitesList)

	comicQueue := make(chan ComicItem, comicNumSum)
}

func createComicSiteList() []*comicSiteTraits {
	comicSiteList := []*comicSiteTraits{}

	poorlyDrawnTraits := &comicSiteTraits{
		SiteName: "PoorlyDrawnLines",
		SiteURL:  "http://poorlydrawnlines.com",
		ComicNum: 10,
	}
	comicSiteList = append(comicSiteList, poorlyDrawnTraits)

	// Repeat for other comic-Sites. Way to automate this better?

	return comicSiteList
}

func sumComicNum(comicSitesTraits []*comicSiteTraits) int {
	var comicNumSum int

	for _, comicSiteTraits := range comicSitesTraits {
		comicNumSum += comicSiteTraits.ComicNum
	}

	return comicNumSum
}
