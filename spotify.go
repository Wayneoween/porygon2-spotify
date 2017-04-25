package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/0x263b/porygon2"
)

const (
	apiEndpoint = "https://api.spotify.com/v1/tracks/"
)

type artistInfo struct{ Name string }

type albumInfo struct{ Name string }

type trackInfo struct {
	ID           string
	Name         string
	Artists      []artistInfo
	Album        albumInfo
	DurationMS   int
	ExternalURLs struct {
		OpenSpotifyURL string `json:"spotify"`
	} `json:"external_urls"`
}

func (track trackInfo) String() string {
	artists := make([]string, len(track.Artists))

	for i, artist := range track.Artists {
		artists[i] = artist.Name
	}

	duration, _ := time.ParseDuration(strconv.Itoa(track.DurationMS) + "ms")

	ret := fmt.Sprintf("%s - %s (%s) %s", strings.Join(artists, ", "), track.Name, track.Album.Name, duration)
	return ret
}

func resolveTrackInfo(trackID string) (*trackInfo, error) {
	resp, err := http.Get(apiEndpoint + trackID)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var track trackInfo
	if err = dec.Decode(&track); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &track, nil
}

func spotify(command *bot.Cmd, matches []string) (msg string, err error) {
	return strings.Join(matches, ", "), nil
}

func init() {
	bot.RegisterCommand("^spotify ([A-Za-z0-9]+)", spotify)
}