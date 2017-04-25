package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/0x263b/porygon2"
)

const (
	apiEndpoint    = "https://api.spotify.com/v1/tracks/"
	formatTemplate = "%s - %s (%s) %s - %s"
)

var patterns = []*regexp.Regexp{
	regexp.MustCompile("spotify:track:([A-Za-z0-9]+)"),
	regexp.MustCompile("https://play.spotify.com/track/([A-Za-z0-9]+)"),
	regexp.MustCompile("https://open.spotify.com/track/([A-Za-z0-9]+)"),
}

type artistInfo struct{ Name string }

type albumInfo struct{ Name string }

type trackInfo struct {
	ID           string
	Name         string
	Artists      []artistInfo
	Album        albumInfo
	DurationMS   int `json:"duration_ms"`
	ExternalURLs struct {
		OpenSpotifyURL string `json:"spotify"`
	} `json:"external_urls"`
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

func spotify(cmd *bot.PassiveCmd) (string, error) {
	var matches []string
	for _, pattern := range patterns {
		if matches = pattern.FindStringSubmatch(cmd.Raw); len(matches) > 0 {
			break
		}
	}

	if len(matches) == 0 {
		return "", nil
	}

	trackInfo, err := resolveTrackInfo(matches[1])
	if err != nil {
		return "", err
	}

	artists := make([]string, len(trackInfo.Artists))
	for i, artist := range trackInfo.Artists {
		artists[i] = artist.Name
	}

	duration, _ := time.ParseDuration(strconv.Itoa(trackInfo.DurationMS) + "ms")

	msg := fmt.Sprintf(formatTemplate,
		strings.Join(artists, ", "),
		trackInfo.Name,
		trackInfo.Album.Name,
		duration,
		trackInfo.ExternalURLs.OpenSpotifyURL)

	return msg, nil
}

func init() {
	bot.RegisterPassiveCommand("spotify", spotify)
}
