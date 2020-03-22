package OGameBotRoutine

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
)

func RestartLogic() {
	log.Info("Restart Logic Started")

	var buffer bytes.Buffer
	for i, elm := range os.Args {
		if i == 0 {
			continue
		}
		buffer.WriteString(elm)

	}

	arguemntString := buffer.String()
	if runtime.GOOS == "windows" {
		//args := []string {"/c","..\\run.bat"}

		s := []string{"cmd.exe", "/C", "start", "run.bat"}
		execErr := exec.Command(s[0], s[1:]...).Run()
		if execErr != nil {
			panic(execErr)
		}

	} else {
		s := []string{"/bin/bash", "-c", "make", "all"}
		execErr := exec.Command(s[0], s[1:]...).Run()
		if execErr != nil {
			panic(execErr)
		}

		s = []string{"/bin/bash", "-c", "killall", "ogamebot", "&&", "./ogamebot", arguemntString}
		execErr = exec.Command(s[0], s[1:]...).Run()
		if execErr != nil {
			panic(execErr)
		}

	}

}
