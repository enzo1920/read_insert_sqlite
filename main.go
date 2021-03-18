package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "path"
    "path/filepath"
    "io/ioutil"
    "strings"
    "database/sql"
    "time"
    _ "github.com/mattn/go-sqlite3"

)
//    
func main() {
    short_fname := ""
    files, err := ioutil.ReadDir("./users")
    if err != nil {
       log.Fatal(err)
    }
    for _, f := range files {
         short_fname = strings.TrimSuffix(f.Name(),filepath.Ext(f.Name()))
         fmt.Println( short_fname)
         file, err := os.Open(path.Join("./users/",f.Name()))
         if err != nil {
              log.Fatal(err)
         }
         defer file.Close()

         scanner := bufio.NewScanner(file)
         for scanner.Scan() {
              fmt.Println(scanner.Text())
              inserter(scanner.Text(), short_fname)
         }

         if err := scanner.Err(); err != nil {
         log.Fatal(err)
    
        }

    }
}

func inserter(user string,project string) {
        db, err := sql.Open("sqlite3", "./PU.db")
        checkErr(err)
        dt := time.Now()
        // insert
        stmt, err := db.Prepare("INSERT INTO project_users(username, project, created) values(?,?,?)")
        checkErr(err)

        res, err := stmt.Exec(user, project, dt.Format("01-01-2021"))
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)

        fmt.Println(id)

}

func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }

