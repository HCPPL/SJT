package main

import (
	//"bufio"
//	"fmt"
	"log"
	"os"
	"time"
	//"net/http"
	//"net"
	
)

	
		
		
var enable bool= true


func CallingLog (method string,api string, user string,email string, detail string){
	if(enable==true){
	currentTime := time.Now()
	var date string=currentTime.Format("01-02-2006")
	//fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
	filename:="callingresultlog"+date+".txt"
	var path = "/home/ubuntu/source/services/logfiles/"+filename	
	
	//log.Println(filename)

	// detect if file exists
var _, err = os.Stat(path)
//log.Println(path)


	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		
		//checkError(err)
		//log.Println(path)
		log.Println(err)
		defer file.Close()
}
f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

if err != nil {
	ErrorLog("ErrorLog", "LogFile.go"," "," ","failed opening file ")
}
defer f.Close()
		
		//printing calling log
		_, err = f.WriteString("\r\n\t\t\t CALLING LOG   ")
		_, err = f.WriteString("\r\n method name =====>"+ method)
		_, err = f.WriteString("\r\n api name =====>"+ api)
		_, err = f.WriteString("\r\n user address =====>"+ user)
		_, err = f.WriteString("\r\n email-id =====>"+ email)
		_, err = f.WriteString("\r\n time =====>"+ currentTime.Format("3:4:5 PM"))
		_, err = f.WriteString("\r\n date =====>"+ date)
		_, err = f.WriteString("\r\n String =====>"+  detail)
		if err != nil {
			ErrorLog("CallingResultLog", "LogFile.go"," "," ","failed in writing file ")
		}
		_, err = f.WriteString("\r\n \t\t ================")
	}
	
}
		func ResultLog(method string,api string, result string, detail string){
		
			if(enable==true){
			currentTime := time.Now()
			var date string=currentTime.Format("01-02-2006")
			//fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
			filename:="callingresultlog"+date+".txt"
			var path = "/home/ubuntu/source/services/logfiles/"+filename	
			
			//log.Println(filename)
		
			// detect if file exists
		var _, err = os.Stat(path)
		//log.Println(path)
		
		
			// create file if not exists
			if os.IsNotExist(err) {
				var file, err = os.Create(path)
				
				//checkError(err)
				//log.Println(path)
				log.Println(err)
				defer file.Close()
		}
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		
		if err != nil {
			ErrorLog("ErrorLog", "LogFile.go"," "," ","failed opening file ")
		}
		defer f.Close()
			//printing Error log
		_, err = f.WriteString("\r\n\t\t\t RESULT LOG   ")
		_, err = f.WriteString("\r\n method name =====>"+ method)
		_, err = f.WriteString("\r\n api name =====>"+ api)
		_, err = f.WriteString("\r\n time =====>"+ currentTime.Format("3:4:5 PM"))
		_, err = f.WriteString("\r\n date =====>"+ date)
		_, err = f.WriteString("\r\n result =====>"+ result)
		_, err = f.WriteString("\r\n String =====>"+  detail)
		if err != nil {
			ErrorLog("CallingResultLog", "LogFile.go"," "," ","failed in writing file ")
		}
		_, err = f.WriteString("\r\n \t\t ================")
	}
	
}


	func ErrorLog(method string,api string, user string,email string,errors string){
		
		if(enable==true){			
		currentTime := time.Now()
		var date string=currentTime.Format("01-02-2006")
		//fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
		filename:="errorlog"+date+".txt"
		var path = "/home/ubuntu/source/services/logfiles/"+filename	
		//log.Println(filename)
		
		// detect if file exists
		var _, err = os.Stat(path)
		//log.Println(path)
	
		// create file if not exists
		if os.IsNotExist(err) {
			var file, err = os.Create(path)
			
			//checkError(err)
			//log.Println(path)
			log.Println(err)
			defer file.Close()
	}
		
			f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			
			if err != nil {
				ErrorLog("ErrorLog", "LogFile.go"," "," ","failed opening file ")
			}
			defer f.Close()
			
			_, err = f.WriteString("\r\n\t\t\t Error LOG   ")
			_, err = f.WriteString("\r\n method name =====>"+ method)
			_, err = f.WriteString("\r\n api name =====>"+ api)
			_, err = f.WriteString("\r\n user address =====>"+ user)
			_, err = f.WriteString("\r\n email =====>"+ email)
			_, err = f.WriteString("\r\n time =====>"+ currentTime.Format("3:4:5 PM"))
			_, err = f.WriteString("\r\n date =====>"+ date)
			_, err = f.WriteString("\r\n Error =====>"+ errors)
			
		
		if err != nil {
				ErrorLog("ErrorLog", "LogFile.go"," "," ","failed in writing file")
			}
			_, err = f.WriteString("\r\n \t\t ================")
		}
		
	}




