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
	// en-US-JennyNeural
	ssml = `
<speak version="1.0" xml:lang="en-US">
	<voice name="en-US-JennyNeural">
			<prosody rate="1" volume="200">
				<express-as style="whispering" styledegree="2">
					at a loss for words
				</express-as>
			</prosody>
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
	// xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="https://www.w3.org/2001/mstts"
	// 	ssml = `
	// <speak version="1.0" // xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="https://www.w3.org/2001/mstts" xml:lang="en-US">
	// 	<voice name="en-US-JennyNeural">
	// 				<mstts:express-as style="whispering">
	// 					at a loss for words
	// 				</mstts:express-as>
	//   </voice>
	// </speak>`
	// 	ssml = `
	// <speak version="1.0" xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="https://www.w3.org/2001/mstts" xml:lang="en-US">
	//     <voice name="my-custom-voice">
	//         <mstts:express-as style="cheerful" styledegree="2">
	//             That'd be just amazing!
	//         </mstts:express-as>
	//         <mstts:express-as style="my-custom-style" styledegree="0.01">
	//             What's next?
	//         </mstts:express-as>
	//     </voice>
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

	// bodyBytes, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(bodyBytes)) // подивись, що реально прийшло

	out, _ := os.Create("loss-jenny-whispering.mp3")
	defer out.Close()

	io.Copy(out, resp.Body)
}
