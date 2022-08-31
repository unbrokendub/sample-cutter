package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func checkSubstrings(str string, subs ...string) bool {

	isCompleteMatch := false

	//   fmt.Printf("String: \"%s\", Substrings: %s\n", str, subs)

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			isCompleteMatch = true
		} else {

		}
	}

	return isCompleteMatch
}

func GetFileSize(filepath string) int64 {
	fi, err := os.Stat(filepath)
	if err != nil {
		return 0
	}
	// get the size
	return fi.Size()
}

func main() {
	var PATH2 string
	err := filepath.Walk("C:\\mpc_samples\\Analog_SamplepackMPC1",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !checkSubstrings(path, ".wav", ".WAV", ".jpg", ".db", ".rtf", ".aif", ".aiff", ".pdf", ".DS_Store") {
				//	fmt.Println(path)
				PATH2 = strings.Replace(path, "Analog_SamplepackMPC1", "Analog_SamplepackCutted", -1)
				PATH2 = strings.Replace(PATH2, " ", "_", -1)
				os.MkdirAll(PATH2, os.ModePerm)
			} else {
				if checkSubstrings(path, ".wav", ".WAV", ".aif", ".aiff") {
					path3 := strings.Replace(path, "Analog_SamplepackMPC1", "Analog_SamplepackCutted", -1)
					path3 = strings.Replace(path3, " ", "_", -1)
					fmt.Println("path=", path)
					fmt.Println("size=", GetFileSize(path)/1000)
					time.Sleep(100)

					if GetFileSize(path)/1000 > 20000 {
						fmt.Println("LARGE FILE")
						fmt.Println("LARGE FILE")
						fmt.Println("LARGE FILE")
						fmt.Println("LARGE FILE")
						fmt.Println("LARGE FILE")
						fmt.Println("LARGE FILE")
						fmt.Println("LARGE FILE")
						path4 := strings.Replace(path3, ".wav", "", -1)
						cmd := exec.Command("ffmpeg.exe", "-i", path, "-f", "segment", "-segment_time", "120", path4+"%03d"+".wav")
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						err = cmd.Run()
					} else {

						cmd := exec.Command("ffmpeg.exe", "-i", path, "-c", "copy", path3)
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						err = cmd.Run()
					}
					//fmt.Println(path3)
					//////cmd := exec.Command("ffmpeg.exe", "-i", path, "-c:a", "pcm_s16le", "-ar", "44100", "-stats", path3)
					//////cmd.Stdout = os.Stdout
					/////cmd.Stderr = os.Stderr
					/////err = cmd.Run()
				}
			}
			//time.Sleep(1)
			return nil
		})
	if err != nil {

	}
}
