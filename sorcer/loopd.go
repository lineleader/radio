package sorcer

import (
	"encoding/xml"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codegoalie/bubbletea-test/models"
	"github.com/codegoalie/bubbletea-test/utils"
)

const loopdName = "Loop'd (Sorcer Radio)\t"

// TODO the numeric subdomain is dynamic
// The web requests this URL to get the subdomain in the Servers part of the LiveStreamConfig struct below:
// https://playerservices.streamtheworld.com/api/livestream?mount=SP_R4852369&transports=http,hls&version=1.10&request.preventCache=1730918512938
// request.preventCache look liks a unix timestamp
// The sbmid in the loopdURL looks like a random UUID
const (
	loopdURL       = "https://14223.live.streamtheworld.com/SP_R4852369.mp3?tdsdk=js-2.9&swm=false&pname=TDSdk&pversion=2.9&banners=none&burst-time=15&sbmid=12fcba0b-1e04-4cbc-c109-822b5ce95fcf"
	loopdStationID = "140481"
)

type Loopd struct{}

func (l Loopd) Name() string {
	return loopdName
}

func (l Loopd) StreamURL() string {
	return loopdURL
}

func (l Loopd) RegisterForUpdates(updates chan models.TrackUpdate) tea.Cmd {
	return utils.SetupUpdateRegister(
		l.Name(),
		infoURL(loopdStationID, "d5791ac233c9aae42aac9342a2d6d70ec0886b5f"),
		parseTrackInfo,
		updates,
	)
}

// LiveStreamConfig was generated 2024-11-06 13:43:34 by https://xml-to-go.github.io/ in Ukraine.
type LiveStreamConfig struct {
	XMLName     xml.Name `xml:"live_stream_config"`
	Text        string   `xml:",chardata"`
	Version     string   `xml:"version,attr"`
	Xmlns       string   `xml:"xmlns,attr"`
	Mountpoints struct {
		Text       string `xml:",chardata"`
		Mountpoint struct {
			Text   string `xml:",chardata"`
			Status struct {
				Text          string `xml:",chardata"`
				StatusCode    string `xml:"status-code"`
				StatusMessage string `xml:"status-message"`
			} `xml:"status"`
			Transports struct {
				Text      string `xml:",chardata"`
				Transport string `xml:"transport"`
			} `xml:"transports"`
			Metadata struct {
				Text        string `xml:",chardata"`
				ShoutcastV1 struct {
					Text        string `xml:",chardata"`
					Enabled     string `xml:"enabled,attr"`
					MountSuffix string `xml:"mountSuffix,attr"`
				} `xml:"shoutcast-v1"`
				ShoutcastV2 struct {
					Text        string `xml:",chardata"`
					Enabled     string `xml:"enabled,attr"`
					MountSuffix string `xml:"mountSuffix,attr"`
				} `xml:"shoutcast-v2"`
				SseSideband struct {
					Text           string `xml:",chardata"`
					Enabled        string `xml:"enabled,attr"`
					StreamSuffix   string `xml:"streamSuffix,attr"`
					MetadataSuffix string `xml:"metadataSuffix,attr"`
				} `xml:"sse-sideband"`
			} `xml:"metadata"`
			Servers struct {
				Text   string `xml:",chardata"`
				Server []struct {
					Text  string `xml:",chardata"`
					Sid   string `xml:"sid,attr"`
					IP    string `xml:"ip"`
					Ports struct {
						Text string `xml:",chardata"`
						Port struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"port"`
					} `xml:"ports"`
				} `xml:"server"`
			} `xml:"servers"`
			Metrics struct {
				Text             string `xml:",chardata"`
				ListenerTracking struct {
					Text         string `xml:",chardata"`
					URL          string `xml:"url,attr"`
					WcmStationID string `xml:"wcm-station-id,attr"`
				} `xml:"listener-tracking"`
				Tag struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"tag"`
			} `xml:"metrics"`
			Mount       string `xml:"mount"`
			Format      string `xml:"format"`
			Bitrate     string `xml:"bitrate"`
			MediaFormat struct {
				Text        string `xml:",chardata"`
				Container   string `xml:"container,attr"`
				Cuepoints   string `xml:"cuepoints,attr"`
				TrackScheme string `xml:"trackScheme,attr"`
				Audio       struct {
					Text       string `xml:",chardata"`
					Index      string `xml:"index,attr"`
					Samplerate string `xml:"samplerate,attr"`
					Codec      string `xml:"codec,attr"`
					Bitrate    string `xml:"bitrate,attr"`
					Channels   string `xml:"channels,attr"`
				} `xml:"audio"`
			} `xml:"media-format"`
			Authentication string `xml:"authentication"`
			Timeout        string `xml:"timeout"`
			SendPageURL    string `xml:"send-page-url"`
		} `xml:"mountpoint"`
	} `xml:"mountpoints"`
}
