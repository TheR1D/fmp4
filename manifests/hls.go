package manifests

import (
	"fmp4/utils"
	"fmt"
	"os"
)

const manifestHeader = `#EXTM3U
#EXT-X-TARGETDURATION:%d
#EXT-X-VERSION:7
#EXT-X-MEDIA-SEQUENCE:1
#EXT-X-PLAYLIST-TYPE:VOD
#EXT-X-INDEPENDENT-SEGMENTS
#EXT-X-MAP:URI="%s",BYTERANGE="%d@%d"
`

const manifestSegment = `#EXTINF:%.5f,	
#EXT-X-BYTERANGE:%d@%d
%s
`

type Hls struct {
	Content string
}

func (h *Hls) String() string {
	return h.Content
}

func NewHls(duration int, uri string, br utils.ByteRange) *Hls {
	hls := Hls{Content: fmt.Sprintf(manifestHeader, duration, uri, br.Length, br.Start)}
	return &hls
}

func (h *Hls) AppendSegment(duration float64, br utils.ByteRange, uri string) {
	h.Content += fmt.Sprintf(manifestSegment, duration, br.Length, br.Start, uri)
}

func (h *Hls) WriteToFile(path string) error {
	err := os.WriteFile(path, []byte(h.Content), 0644)
	return err
}

func (h *Hls) Finalize() {
	h.Content += "#EXT-X-ENDLIST\n"
}
