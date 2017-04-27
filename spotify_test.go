package spotify

import (
	"fmt"
	"testing"
	"time"

	"github.com/0x263b/porygon2"
)

func TestResolveTrackInfo(t *testing.T) {
	trackID := "77hmyDdKIn5JmAJmWp6264"

	expected := trackInfo{
		Name: "Back to the Middle",
		Artists: []artistInfo{
			artistInfo{Name: "Deerhunter"},
		},
		Album:      albumInfo{Name: "Monomania"},
		DurationMS: 156533,
	}

	trackInfo, err := resolveTrackInfo(trackID)
	if err != nil {
		t.Errorf("%v", err)
	}

	if trackInfo.Name != expected.Name {
		t.Errorf("expected track name to be %v, but got %v", expected.Name, trackInfo.Name)
	}

	if trackInfo.Album.Name != expected.Album.Name {
		t.Errorf("expected album name to be %v, but got %v", expected.Album.Name, trackInfo.Album.Name)
	}

	if len(trackInfo.Artists) != len(expected.Artists) {
		t.Errorf("expected artist length to be %v, but got %v", len(expected.Artists), len(trackInfo.Artists))
	}

	if trackInfo.DurationMS != expected.DurationMS {
		t.Errorf("expected duration to be %v, but got %v", expected.DurationMS, trackInfo.DurationMS)
	}
}

func TestOpenSpotifyURL(t *testing.T) {
	trackID := "77hmyDdKIn5JmAJmWp6264"
	expected := "https://open.spotify.com/track/77hmyDdKIn5JmAJmWp6264"

	trackInfo, err := resolveTrackInfo(trackID)
	if err != nil {
		t.Errorf("%v", err)
	}

	openSpotifyURL := trackInfo.ExternalURLs.OpenSpotifyURL

	if openSpotifyURL != expected {
		t.Errorf("expected url to be %v, got %v", expected, openSpotifyURL)
	}
}

func TestSpotify(t *testing.T) {
	cmd := bot.PassiveCmd{
		Raw:     "lol random spotify:track:3n75gL3WU5tAwwAgssRI9j",
		Channel: "foo",
		Nick:    "randomdude",
	}

	duration := time.Duration(118960/1000) * time.Second
	expected := fmt.Sprintf(formatTemplate, "Misfits", "Skulls", "Collection", duration,
		"https://open.spotify.com/track/3n75gL3WU5tAwwAgssRI9j")

	got, err := spotify(&cmd)
	if err != nil {
		t.Errorf("%v", err)
	}

	if got != expected {
		t.Errorf("expected output to be %v, got %v", expected, got)
	}
}
