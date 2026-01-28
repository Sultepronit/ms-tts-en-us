package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	ssml := `
<speak version="1.0" xml:lang="ja-JP">
  <voice name="ja-JP-NanamiNeural">
    こいびと
  </voice>
</speak>`
	//en-US-JennyNeural
	ssml = `
<speak version="1.0" xml:lang="en-US">
  <voice name="en-US-NancyMultilingualNeural">
    <express-as style="assistant">
      "lest"
    </express-as>
  </voice>
</speak>`
	// Andrew
	// Ava
	// Davis
	// Steffan
	// think back
	// en-US-AIGenerate1Neural
	// en-US-MonicaNeural
	ssml = `
<speak version="1.0" xml:lang="en-US">
	<voice name="en-US-MonicaNeural">
		<express-as style="chat">
			<prosody rate="1" volume="150">
				"think back"
			</prosody>
		</express-as>
  </voice>
</speak>`
	// 	ssml = `
	// <speak version="1.0" xml:lang="en-US">
	// 	<voice name="en-US-Ava:DragonHDLatestNeural">
	// 		<prosody rate="1" volume="150">
	// 			at a loss for words
	// 		</prosody>
	//   </voice>
	// </speak>`

	req, _ := http.NewRequest(
		"POST",
		"https://"+os.Getenv("SPEECH_REGION")+".tts.speech.microsoft.com/cognitiveservices/v1",
		bytes.NewBufferString(ssml),
	)

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

	out, _ := os.Create("think-monica-chat.mp3")
	defer out.Close()

	io.Copy(out, resp.Body)
}
