package harbour

/*
  #cgo CFLAGS: -I/usr/local/include/harbour -I/usr/include/postgresql
  #cgo LDFLAGS: -L/usr/local/lib/harbour -lpq -lhbmzip -lz -lminizip -lhbct -lhbtip -lhbpgsql  -lharbour
  #include <hbvmint.h>
  #include <hbapi.h>
  #include <hbstack.h>
  #include <hbapiitm.h>
  #include <hbapierr.h>
  #include <hbapifs.h>
  #include <hbvm.h>
  #include <hbpcode.h>
  #include <hbset.h>
  #include <hb_io.h>
  #include <hbhrb.ch>
  #include "go.h"
  #include <stdio.h>
  extern void HB_FUNC_GO_FUN();
*/
import "C"

import(
  "fmt"
  "archive/zip"
  "log"
  "os"
  "io"
//  "unsafe"
)

//export HbGo
func HbGo(i int, cstr *C.char) {
    fmt.Printf("i: %d str: %s\n", i, C.GoString(cstr))
    fmt.Println("Harbor function written in go - pokrenuta iz HB_FUNC_GO_FUN (go.c)")
}

func HrbLoad(file string, mode int) {

   C.hb_vmSetDefaultGT(C.CString("xwc"));
   //C.hb_vmSetDefaultGT(C.CString("trm"));
   C.hb_vmInit( C.HB_BOOL(0) )

/*
   LoadDo("hrbext.hrb")     
 
   pDynSym := C.hb_dynsymGet(C.CString("B"))
   pItem := C.hb_dynsymGetMemvar(pDynSym)


   fmt.Println("go string -> harbour memvar")
   C.hb_itemPutC(pItem, C.CString("GOGOGOGOGOOOOOOOOO"))
   
   fmt.Printf("Uzimam iz harbour memvar B(1): %s\n", C.GoString(C.hb_itemGetC(pItem))) 
 

   LoadDo("hrbext2.hrb")     
 
   fmt.Println("harbour memvar B (nakon inputa) -> go string")
   pItem = C.hb_dynsymGetMemvar(pDynSym)
   fmt.Printf("Uzimam iz harbour memvar B(2): %s\n", C.GoString(C.hb_itemGetC(pItem))) 
*/

   //LoadDo("pg_test.hrb") 

   //fmt.Println("ucitavam tpostgre harbour klasu")   
   //LoadDo("tpostgre.hrb")

   //fmt.Println("test pgtest")
   //LoadDo("hbgo.hrb") 
   
   fmt.Println("ucitavam tpostgre harbour klasu")   
   LoadDoHrb("tpostgre.hrb")



//   LoadZip("common")
//   LoadHrb("common/" + "f18_ini_utils.hrb")
//   LoadHrb("common/" + "f18_init_login.hrb")
//   LoadHrb("common/" + "f18_init_semaphores.hrb")
     LoadDoHrb("common.hrb")


   LoadZip("fin")

   //C.hb_conOutStd( C.CString("Ernad Husremovic\n"), C.HB_SIZE(20) ) 
   fmt.Println("kraj")



}


func LoadZip(file string) {

r, err := zip.OpenReader(file + ".zip")
if err != nil {
     log.Fatal(err)
}
defer r.Close()

// Iterate through the files in the archive,
// printing some of their contents.
for _, f := range r.File {
            fmt.Printf("Contents of %s:\n", f.Name)
            rc, err := f.Open()
            if err != nil {
                    log.Fatal(err)
            }


            d, errd := os.Create(file + "/" + f.Name)
	    if errd != nil {
		log.Fatal(errd)
	    }
	    if _, errc := io.Copy(d, rc); errc != nil {
		d.Close()
		log.Fatal(errc)
	    }
            rc.Close()
            fmt.Println()
}

}

func LoadDoHrb(file string) {

  fmt.Println("load and do hrb ", file)
  pBody := C.hb_hrbLoadFromFile( C.CString(file), C.HB_USHORT(1))
  C.hb_hrbDo( pBody, 0, nil );
 
}

func LoadHbPostgresqlC() {
        var conn *C.PGconn
        C.PQreset(conn)
}

