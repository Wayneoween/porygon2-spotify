package spotify

import (
	"testing"
)

func TestResolveTrackInfo(t *testing.T) {
	trackID := "77hmyDdKIn5JmAJmWp6264"

	expected := trackInfo{
		Name: "Back to the Middle",
		Artists: []artistInfo{
			artistInfo{Name: "Deerhunter"},
		},
		Album: albumInfo{Name: "Monomania"},
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
