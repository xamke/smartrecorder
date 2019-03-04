package main

import (
        "fmt"
        "time"
        "syscall"
        "net/http"
        "os"
        "log"
        "strings"
        "encoding/base64"
        "encoding/json"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
        "github.com/gorilla/mux"
)

type SystemInformation struct {
        //Daemon Status
        AppName         string
        Version         string
        Connected       uint64
        StartTime       time.Time
        TotalRequestGet uint64
        TotalRequestPut uint64

        //Request
        RequestIP       string
        RequestData     string
        RequestUser     string
        RequestPass     string
        DBConnected     bool
}

type Configuration struct {
        Database struct {
                Type    string
                Host    string
                User    string
                Port    string
                Password string
                Name string
        }

        Listen struct {
                Address string
                Port    string
        }
        Auth struct {
                User    string
                Password        string
        }

        Log struct {
                Type    string
                Path    string
                FileName        string
        }
}


const (
        version = "1.0.0.0"
        program = "UNICAST Data Logger"
)


var                     (
        status  SystemInformation
        config  Configuration
)
var username = []byte("hello")
var password = []byte("password")

func BasicAuth(w http.ResponseWriter, r *http.Request, user, pass []byte) bool {
         s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
         if len(s) != 2 {
                 return false
         }

         b, err := base64.StdEncoding.DecodeString(s[1])
         if err != nil {
                 return false
         }
login as: sebae
sebae@dev.unicast.co.kr's password:
Linux dev 4.9.0-7-amd64 #1 SMP Debian 4.9.110-3+deb9u2 (2018-08-13) x86_64

The programs included with the Debian GNU/Linux system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent
permitted by applicable law.
Last login: Sat Feb 16 23:40:04 2019 from 14.33.176.5
sebae@dev:~$
sebae@dev:~$ ls
go
sebae@dev:~$ cd go
sebae@dev:~/go$
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$ cd go
-bash: cd: go: No such file or directory
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$ vi unilog.go
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$ vi unilog.go
}


const (
        version = "1.0.0.0"
        program = "UNICAST Data Logger"
)


var                     (
        status  SystemInformation
        config  Configuration
)
var username = []byte("hello")
var password = []byte("password")

func BasicAuth(w http.ResponseWriter, r *http.Request, user, pass []byte) bool {
         s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
         if len(s) != 2 {
                 return false
         }

         b, err := base64.StdEncoding.DecodeString(s[1])
         if err != nil {
                 return false
         }

         pair := strings.SplitN(string(b), ":", 2)
         if len(pair) != 2 {
                 return false
         }

         return pair[0] == string(user) && pair[1] == string(pass)
 }



func init() {
        status.StartTime = time.Now()
        status.Version = version
        status.AppName = program
        status.TotalRequestPut = 0
        status.TotalRequestGet = 0
        status.DBConnected = false

        readConfig()


        writeTextLog("SYS", "Daemon Initialize")
}



func writeMySqlLog(msg string) {
//      db, err := sql.Open("mysql", config.Database.Username+ ":" + config.DataBase.Password +
//      "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Databse.Name + ")"
db, err := sql.Open("mysql", "sebae:s445850@tcp(127.0.0.1:3306)/unicast")
     if err != nil {
                log.Fatal(err)
         }
        defer db.Close()


        result, err := db.Exec("INSERT INTO log (data) VALUES(?)", msg)
        if err != nil {
                log.Println("mysql connect error")
                log.Fatal(err)
        }
        n, err := result.RowsAffected()
        if n == 1 {
                fmt.Println("1 row insertd.")
        }

}


func connectDatabase(driver string, constr string) {
        db, err := sql.Open("mysql", "sebae:s445850@tcp(127.0.0.1:3306)/unicast")
         if err != nil {
               // log.Fatal(err)
         } else  {
         }
        defer db.Close()
        
        
        login as: sebae
sebae@dev.unicast.co.kr's password:
Linux dev 4.9.0-7-amd64 #1 SMP Debian 4.9.110-3+deb9u2 (2018-08-13) x86_64

The programs included with the Debian GNU/Linux system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent
permitted by applicable law.
Last login: Sat Feb 16 23:40:04 2019 from 14.33.176.5
sebae@dev:~$
sebae@dev:~$ ls
go
sebae@dev:~$ cd go
sebae@dev:~/go$
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$ cd go
-bash: cd: go: No such file or directory
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$ vi unilog.go
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$ vi unilog.go
        }

        Log struct {
                Type    string
                Path    string
                FileName        string
        }
}


const (
        version = "1.0.0.0"
        program = "UNICAST Data Logger"
)


var                     (
        status  SystemInformation
        config  Configuration
)
var username = []byte("hello")
var password = []byte("password")

func BasicAuth(w http.ResponseWriter, r *http.Request, user, pass []byte) bool {
         s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
         if len(s) != 2 {
                 return false
         }

         b, err := base64.StdEncoding.DecodeString(s[1])
         if err != nil {
                 return false
         }

         pair := strings.SplitN(string(b), ":", 2)
         if len(pair) != 2 {
                 return false
         }

         return pair[0] == string(user) && pair[1] == string(pass)
 }



func init() {
        status.StartTime = time.Now()
        status.Version = version
        status.AppName = program
        status.TotalRequestPut = 0
        status.TotalRequestGet = 0
        status.DBConnected = false

        readConfig()


        writeTextLog("SYS", "Daemon Initialize")
}



func writeMySqlLog(msg string) {
//      db, err := sql.Open("mysql", config.Database.Username+ ":" + config.DataBase.Password +
//      "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Databse.Name + ")"
db, err := sql.Open("mysql", "ide:pwd0@tcp(127.0.0.1:3306)/database")
     if err != nil {
                log.Fatal(err)
         }
        defer db.Close()


        result, err := db.Exec("INSERT INTO log (data) VALUES(?)", msg)
        if err != nil {
                log.Println("mysql connect error")
                log.Fatal(err)
        }
        n, err := result.RowsAffected()
        if n == 1 {
                fmt.Println("1 row insertd.")
        }

}


func connectDatabase(driver string, constr string) {

login as: sebae
sebae@dev.unicast.co.kr's password:
Linux dev 4.9.0-7-amd64 #1 SMP Debian 4.9.110-3+deb9u2 (2018-08-13) x86_64

The programs included with the Debian GNU/Linux system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent
permitted by applicable law.
Last login: Sat Feb 16 23:40:04 2019 from 14.33.176.5
sebae@dev:~$
sebae@dev:~$ ls
go
sebae@dev:~$ cd go
sebae@dev:~/go$
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$ cd go
-bash: cd: go: No such file or directory
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$ vi unilog.go
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$ vi unilog.go

}


func connectDatabase(driver string, constr string) {
        db, err := sql.Open("mysql", "sebae:s445850@tcp(127.0.0.1:3306)/unicast")
         if err != nil {
               // log.Fatal(err)
         } else  {
         }
        defer db.Close()


}

func readConfig() {
        file, err := os.Open("unilog.conf")

        if err != nil {
                log.Fatal("can't open config file: ", err)
        }
        defer file.Close()
        decoder := json.NewDecoder(file)
        err = decoder.Decode(&config)
        if err != nil {
                log.Fatal("can't decode config JSON: ", err)
        }
                log.Println(config.Database.Host)
}

func uptime() time.Duration {
        return time.Since(status.StartTime)
}


func writeTextLog(prefix string, msg string) {
        curdate := time.Now()
        filename := config.Log.Path + curdate.Format("2016-01-02")
        //f, err := os.OpenFile(config.Log.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
                log.Println(err)
        }
        defer f.Close()

        logger := log.New(f, prefix + ": ", log.LstdFlags)
        logger.Println(msg)
}

func writeSysLog(msg string) {
}

func writeSysStatus(w http.ResponseWriter) {
        info := syscall.Sysinfo_t{}
        err := syscall.Sysinfo(&info)

        if err != nil {
                fmt.Fprintf(w, "Error:%s\n", err)
        }

        fmt.Fprintf(w, "----------------------------- \n")
        fmt.Fprintf(w, "System Uptime: %d\n", info.Uptime)
        fmt.Fprintf(w, "System Loads %d\n", info.Loads)
        fmt.Fprintf(w, "Total Memory: %d\n", info.Totalram)
        fmt.Fprintf(w, "Free Memory: %d\n", info.Freeram)
        fmt.Fprintf(w, "Shared Memory: %d\n", info.Sharedram)
        fmt.Fprintf(w, "Total Swap: %d\n", info.Sharedram)
        fmt.Fprintf(w, "Free Swap:%d\n", info.Freeswap)
}



func writeHttpStatus(w http.ResponseWriter) {
        fmt.Fprintf(w, "%s\n\n", status.AppName)
        fmt.Fprintf(w, "----------------------------- \n")
        fmt.Fprintf(w, "Version %s\n", status.Version)
        fmt.Fprintf(w, "Service Start: %s\n", status.StartTime)
        fmt.Fprintf(w, "Service UpTime: %s\n", uptime())
        fmt.Fprintf(w, "Log Type: %s\n", config.Log.Type)
        fmt.Fprintf(w, "LogFile %s\n", config.Log.FileName)
        fmt.Fprintf(w, "Total Request Get:%d\n", status.TotalRequestGet)
        fmt.Fprintf(w, "Total Request Put:%d\n", status.TotalRequestPut)
}


login as: sebae
sebae@dev.unicast.co.kr's password:
Linux dev 4.9.0-7-amd64 #1 SMP Debian 4.9.110-3+deb9u2 (2018-08-13) x86_64

The programs included with the Debian GNU/Linux system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent
permitted by applicable law.
Last login: Sat Feb 16 23:40:04 2019 from 14.33.176.5
sebae@dev:~$
sebae@dev:~$ ls
go
sebae@dev:~$ cd go
sebae@dev:~/go$
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$ cd go
-bash: cd: go: No such file or directory
sebae@dev:~/go$ ls
data  log  pkg  src  test  test.go  unilog  unilog.conf  unilog.go
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$ vi unilog.go
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$
sebae@dev:~/go$ vi unilog.go

func writeHttpStatus(w http.ResponseWriter) {
        fmt.Fprintf(w, "%s\n\n", status.AppName)
        fmt.Fprintf(w, "----------------------------- \n")
        fmt.Fprintf(w, "Version %s\n", status.Version)
        fmt.Fprintf(w, "Service Start: %s\n", status.StartTime)
        fmt.Fprintf(w, "Service UpTime: %s\n", uptime())
        fmt.Fprintf(w, "Log Type: %s\n", config.Log.Type)
        fmt.Fprintf(w, "LogFile %s\n", config.Log.FileName)
        fmt.Fprintf(w, "Total Request Get:%d\n", status.TotalRequestGet)
        fmt.Fprintf(w, "Total Request Put:%d\n", status.TotalRequestPut)
}



func httpWebroot(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%s\n\n", status.AppName)
        fmt.Fprintf(w, "----------------------------- \n")
        fmt.Fprintf(w, "Version %s\n", status.Version)
        fmt.Fprintf(w, "Service Start: %s\n", status.StartTime)
        fmt.Fprintf(w, "Service UpTime: %s\n", uptime())
}
func httpStatus(w http.ResponseWriter, r *http.Request) {
        status.TotalRequestGet++
        writeHttpStatus(w)
        writeSysStatus(w)
}

func httpProtect(w http.ResponseWriter, r *http.Request) {
 // pass from global variables
         if BasicAuth(w, r, username, password) {
                 w.Write([]byte("Welcome to the Protected Page!!"))
                 return
         }

         w.Header().Set("WWW-Authenticate", `Basic realm="Beware! Protected REALM! "`)
         w.WriteHeader(401)
         w.Write([]byte("401 Unauthorized\n"))

}
func httpPut(w http.ResponseWriter, r *http.Request) {
        status.TotalRequestPut++
        vars := mux.Vars(r)

        status.RequestData = vars["data"]

        w.WriteHeader(http.StatusOK)

        writeHttpStatus(w)

        switch config.Log.Type {
        case "text":
                writeTextLog("DATA", vars["data"])
        case "mysql":
                writeMySqlLog(vars["data"])
        }


}




func main() {


          r := mux.NewRouter()

        r.HandleFunc("/", httpWebroot)
        r.HandleFunc("/get", httpStatus)
        r.HandleFunc("/put/{data}", httpPut)
        r.HandleFunc("/get/status", httpStatus)
        r.HandleFunc("/protect/", httpProtect)

        if err:= http.ListenAndServe(config.Listen.Address + ":" + config.Listen.Port, r);  err != nil {
                panic(err)
        }
}

