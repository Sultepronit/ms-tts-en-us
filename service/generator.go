package service

import (
	"bytes"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"strings"
)

func Generate(text string, voice string) ([]byte, error) {
	ssml := fmt.Sprintf(`
<speak version="1.0" xml:lang="en-US">
	<voice name="%s">
		<prosody volume="+100%%">
			%s
		</prosody>
  </voice>
</speak>`, voice, text)

	ssml = fmt.Sprintf(`
	<speak version="1.0" xml:lang="en-US">
		<voice name="%s">
			%s
	  </voice>
	</speak>`, voice, text)
	// fmt.Println(ssml)

	url := "https://" + os.Getenv("SPEECH_REGION") + ".tts.speech.microsoft.com/cognitiveservices/v1"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(ssml))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Ocp-Apim-Subscription-Key", os.Getenv("SPEECH_KEY"))
	req.Header.Set("Content-Type", "application/ssml+xml")
	req.Header.Set("X-Microsoft-OutputFormat", "audio-24khz-48kbitrate-mono-mp3")
	req.Header.Set("User-Agent", "go-tts")
	req.Header.Set("Accept", "audio/mpeg")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := fmt.Errorf("bad tts service's status: %d", resp.StatusCode)
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

func Generate0(text string, voice string) {
	ssml := fmt.Sprintf(`
<speak version="1.0" xml:lang="en-US">
	<voice name="%s">
		<prosody volume="200">
			%s
		</prosody>
  </voice>
</speak>`, voice, text)

	url := "https://" + os.Getenv("SPEECH_REGION") + ".tts.speech.microsoft.com/cognitiveservices/v1"
	req, _ := http.NewRequest("POST", url, bytes.NewBufferString(ssml))
	req.Header.Set("Ocp-Apim-Subscription-Key", os.Getenv("SPEECH_KEY"))
	req.Header.Set("Content-Type", "application/ssml+xml")
	req.Header.Set("X-Microsoft-OutputFormat", "audio-24khz-48kbitrate-mono-mp3")
	req.Header.Set("User-Agent", "go-tts")
	req.Header.Set("Accept", "audio/mpeg")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic(resp.Status)
	}

	// bodyBytes, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(bodyBytes)) // подивись, що реально прийшло

	esc := strings.ReplaceAll(text, " ", "_")
	ri := rand.IntN(10)
	err = os.MkdirAll("records/"+esc, 0755)
	if err != nil {
		panic(err)
	}
	fn := fmt.Sprintf("records/%s/%d.mp3", esc, ri)
	// out, err := os.Create("records/" + esc + ".mp3")
	out, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	io.Copy(out, resp.Body)
}
