package main

import(
	"fmt"
	"log"
	"os"
	"time"
)

func main(){
	logfile,err:=os.OpenFile("./test.log",os.O_RDWR|os.O_CREATE,0666)
	if err!=nil{
		fmt.Printf("%s\r\n",err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger:=log.New(logfile,"\r\n",log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	logger.Println("INFO ", "hello")
	logger.Println("DEBUG ","oh....")
	logger.Fatal("test")
	logger.Fatal("test2")

	time.Sleep(time.Second*1)
}