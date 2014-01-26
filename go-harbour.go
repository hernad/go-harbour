package harbour

/*
  #cgo CFLAGS: -I/usr/local/include/harbour -I/usr/include/postgresql
  #cgo LDFLAGS: -L/usr/local/lib/harbour -lpq -lhbpgsql  -lharbour
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
//  "unsafe"
)

//export HbGo
func HbGo() {
    fmt.Println("Harbor function written in go")
}

func HrbLoad(file string, mode int) {

   C.hb_vmSetDefaultGT(C.CString("xwc"));
   //C.hb_vmSetDefaultGT(C.CString("trm"));
   C.hb_vmInit( C.HB_BOOL(0) )

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

   LoadDo("pg_test.hrb") 
   fmt.Println("ucitavam tpostgre harbour klasu")   
   LoadDo("tpostgre.hrb")

   //fmt.Println("test pgtest")
   //LoadDo("hbgo.hrb") 
   
   //C.hb_conOutStd( C.CString("Ernad Husremovic\n"), C.HB_SIZE(20) ) 
   fmt.Println("kraj")

}


func LoadDo(file string) {

  pBody := C.hb_hrbLoadFromFile( C.CString(file), C.HB_USHORT(1))
  C.hb_hrbDo( pBody, 0, nil );
 
}

func LoadHbPostgresqlC() {
        var conn *C.PGconn
        C.PQreset(conn)
}


