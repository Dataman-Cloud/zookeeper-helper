package logger

import (
	"strconv"
)

//set log file name with size
// every log file max size is 5M
func loggerConfig() string {
	config := `<seelog type="asyncloop" minlevel="` + level + `">  
			    <outputs formatid="main">`
	if logToScreen {
		config += `<console/>`
	}

	config += `<buffered size="10000" flushperiod="1000">  
			            <rollingfile type="size" filename="` + logfile + `" maxsize="` + strconv.Itoa(fileSize) + `" maxrolls="` + strconv.Itoa(fileNum) + `"/>  
			        </buffered>  
			    </outputs>  
			    <formats>  
			        <format id="main" format="` + formatString + `" />  
			    </formats>  
			</seelog>`
	return config
}
