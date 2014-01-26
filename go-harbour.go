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

   pBody := C.hb_hrbLoadFromFile( C.CString("hrbext.hrb"), C.HB_USHORT(1))
   C.hb_hrbDo( pBody, 0, nil );

 
   pDynSym := C.hb_dynsymGet(C.CString("B"))
   pItem := C.hb_dynsymGetMemvar(pDynSym)


   fmt.Println("go string -> harbour memvar")
   C.hb_itemPutC(pItem, C.CString("GOGOGOGOGOOOOOOOOO"))
   
   fmt.Printf("Uzimam iz harbour memvar B(1): %s\n", C.GoString(C.hb_itemGetC(pItem))) 
 
     
   pBody = C.hb_hrbLoadFromFile( C.CString("hrbext2.hrb"), C.HB_USHORT(1))
   C.hb_hrbDo( pBody, 0, nil );
 
   fmt.Println("harbour memvar B (nakon inputa) -> go string")
   pItem = C.hb_dynsymGetMemvar(pDynSym)
   fmt.Printf("Uzimam iz harbour memvar B(2): %s\n", C.GoString(C.hb_itemGetC(pItem))) 

   LoadDo("pg_test.hrb") 
   fmt.Println("ucitavam tpostgre harbour klasu")   
   LoadDo("tpostgre.hrb")

/*

   fmt.Println("pokusavam injectovati PQSETDBLOGIN")
   type HbSymbol struct { 
       name string
       hbtype int
       hbfun  unsafe.Pointer
       dummy int
   }


   var symb2, symb3, symb4, symb5, symb6, symb7 HbSymbol
   symb2 = HbSymbol{ name: "PQSETDBLOGIN",  hbtype: int(C.HB_FS_PUBLIC), hbfun:  C.HB_FUN_PQSETDBLOGIN, dummy: 0 }
   C.hb_dynsymNew( (C.PHB_SYMB)(unsafe.Pointer(&symb2)))
   symb3 = HbSymbol{ name: "PQDB",  hbtype: int(C.HB_FS_PUBLIC), hbfun:  C.HB_FUN_PQDB, dummy: 0 }
   C.hb_dynsymNew( (C.PHB_SYMB)(unsafe.Pointer(&symb3)))
   symb4 = HbSymbol{ name: "PQUSER",  hbtype: int(C.HB_FS_PUBLIC), hbfun:  C.HB_FUN_PQUSER, dummy: 0 }
   C.hb_dynsymNew( (C.PHB_SYMB)(unsafe.Pointer(&symb4)))
   symb5 = HbSymbol{ name: "PQPASS",  hbtype: int(C.HB_FS_PUBLIC), hbfun:  C.HB_FUN_PQPASS, dummy: 0 }
   C.hb_dynsymNew( (C.PHB_SYMB)(unsafe.Pointer(&symb5)))

   symb6 = HbSymbol{ name: "PQHOST",  hbtype: int(C.HB_FS_PUBLIC), hbfun:  C.HB_FUN_PQHOST, dummy: 0 }
   C.hb_dynsymNew( (C.PHB_SYMB)(unsafe.Pointer(&symb6)))

   
   symb7 = HbSymbol{ name: "GO_FUN",  hbtype: int(C.HB_FS_PUBLIC), hbfun: C.HB_FUNC_GO_FUN, dummy: 0 }
   C.hb_dynsymNew( (C.PHB_SYMB)(unsafe.Pointer(&symb7)))


*/


   //fmt.Println("test pgtest")
   //LoadDo("hbgo.hrb") 
   
   //C.hb_conOutStd( C.CString("Ernad Husremovic\n"), C.HB_SIZE(20) ) 
   fmt.Println("kraj")

   //C.HB_FUN_PQSETDBLOGIN()
   
   var conn *C.PGconn
   C.PQreset(conn)
  
   //C.HB_FUN_PQRESET()
   //C.HB_FUN_PQSETDBLOGIN()

     //pSymb :=  C.hb_vmProcessSymbols( C.hb_vm_SymbolInit_TPOSTGRE, C.HB_USHORT(10), C.CString("test.prg"), C.HB_ULONG(0), C.HB_USHORT(3) )

}


func LoadDo(file string) {

  pBody := C.hb_hrbLoadFromFile( C.CString(file), C.HB_USHORT(1))
  C.hb_hrbDo( pBody, 0, nil );
 
}
