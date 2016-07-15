package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
)

var (
	LASTFM_API_KEY = "e4f57ad61d672ed443210dbfb82c55ed"
)

type Track struct {
	Name string `xml:"name"`
}

type TrackList struct {
	Tracks []Track `xml:"track"`
}

func stand(genre string) string {
	res, _ := http.Get("http://powerlisting.wikia.com/wiki/Special:Random")
	b, _ := ioutil.ReadAll(res.Body)
	html := string(b)
	desc_regex := regexp.MustCompile("<meta name=\"description\" content=\".*\" />")
	desc := desc_regex.FindString(html)
	desc = desc[34:]
	for i := 0; i < len(desc); i++ {
		if desc[i] == '.' {
			desc = desc[:i+1]
			break
		}
	}
	name_regex := regexp.MustCompile("<title>.* - Superpower Wiki - Wikia</title>")
	name := name_regex.FindString(html)
	name = name[7 : len(name)-34]

	res = nil
	track := ""
	stand := ""
	if genre != "" && genre != "all" {
		res, _ = http.Get(fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=tag.gettoptracks&tag=%v&api_key=%v", genre, LASTFM_API_KEY))
		b, _ = ioutil.ReadAll(res.Body)
		b = b[57 : len(b)-7] //Please don't ask
		var tl TrackList
		xml.Unmarshal(b, &tl)
		if len(tl.Tracks) == 0 {
			stand += "I couldn't find anything for the tag \"" + genre + "\"; the default song list will be used.\n\n"
			genre = "all"
		} else {
			track = tl.Tracks[rand.Intn(len(tl.Tracks))].Name
		}
	}
	if genre == "" || genre == "all" {
		res, _ = http.Get("http://ws.audioscrobbler.com/2.0/?method=chart.gettoptracks&api_key=" + LASTFM_API_KEY)
		b, _ = ioutil.ReadAll(res.Body)
		b = b[57 : len(b)-7]
		var tl TrackList
		xml.Unmarshal(b, &tl)
		track = tl.Tracks[rand.Intn(len(tl.Tracks))].Name
	}
	stand += "Stand Name: 「" + track + "」\n\n"
	stand += "Destructive Power: " + string(rune(rand.Intn(5)+'A')) + "\n"
	stand += "Speed: " + string(rune(rand.Intn(5)+'A')) + "\n"
	stand += "Range: " + string(rune(rand.Intn(5)+'A')) + "\n"
	stand += "Durability: " + string(rune(rand.Intn(5)+'A')) + "\n"
	stand += "Precision: " + string(rune(rand.Intn(5)+'A')) + "\n"
	stand += "Development Potential: " + string(rune(rand.Intn(5)+'A')) + "\n"
	stand += "\nPower Name: 「" + name + "」" + "\n"
	stand += "Power Description: " + desc + "\n"
	return stand
}
