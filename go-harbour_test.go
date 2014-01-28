package main

import (
   "testing"
   "github.com/vmihailenco/pg"
   "fmt"
)

func Test1(t *testing.T) {
//   HrbLoad("file.hrb", 4)    
}

type SifTest struct {
   Id string
   Name string
}

type SifTests  []*SifTest

func (sifTests *SifTests) New() interface{} {
  s := &SifTest{}
  *sifTests = append(*sifTests, s)
  return s
}


func TestPgCon(t *testing.T) { 

   var sifTests SifTests

   db := pg.Connect(&pg.Options{
        User: "bringout", Password: "b", Database: "bringout",
   }) 

   _, err := db.Query(&sifTests, `SELECT * FROM siftest`)

   for _, sifTest := range sifTests {
        fmt.Printf("Record# %s: %s\n", sifTest.Id, sifTest.Name)
   }
   err = db.Close()
   fmt.Println(err)

}
