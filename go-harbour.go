package harbour

// #cgo CFLAGS: -I/usr/local/include/harbour
// #cgo LDFLAGS: -lharbour
// #include <hbvmint.h>
// #include <hbapi.h>
// #include <hbstack.h>
// #include <hbapiitm.h>
// #include <hbapierr.h>
// #include <hbapifs.h>
// #include <hbvm.h>
// #include <hbpcode.h>
// #include <hbset.h>
// #include <hb_io.h>
// #include <hbhrb.ch>
// #include "go.h"
import ("C")

import(
  "fmt"
)

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
   

   //C.hb_conOutStd( C.CString("Ernad Husremovic\n"), C.HB_SIZE(20) ) 
   fmt.Println("kraj")

}
