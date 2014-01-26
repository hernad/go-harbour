#include "hbvmint.h"
#include "hbapi.h"
#include "hbstack.h"
#include "hbapiitm.h"
#include "hbapierr.h"
#include "hbapifs.h"
#include "hbvm.h"
#include "hbpcode.h"
#include "hbset.h"
#include "hb_io.h"
#include "hbpgsql.h"
#include "hbhrb.ch"



typedef struct
{
   char *        szName;                        /* Name of the function     */
   PHB_PCODEFUNC pCodeFunc;                     /* Dynamic function info    */
   HB_BYTE *     pCode;                         /* P-code                   */
} HB_DYNF, * PHB_DYNF;

typedef struct
{
   HB_ULONG    ulSymbols;                       /* Number of symbols        */
   HB_ULONG    ulFuncs;                         /* Number of functions      */
   HB_BOOL     fInit;                           /* should be INIT functions executed */
   HB_BOOL     fExit;                           /* should be EXIT functions executed */
   HB_LONG     lSymStart;                       /* Startup Symbol           */
   PHB_SYMB    pSymRead;                        /* Symbols read             */
   PHB_DYNF    pDynFunc;                        /* Functions read           */
   PHB_SYMBOLS pModuleSymbols;
} HRB_BODY, * PHRB_BODY;

extern void hb_hrbDo( PHRB_BODY pHrbBody, int iPCount, PHB_ITEM * pParams );
extern PHRB_BODY hb_hrbLoadFromFile( const char * szHrb, HB_USHORT usMode );

// postgres.c
extern void HB_FUN_PQSETDBLOGIN();
extern void HB_FUN_PQDB();
extern void HB_FUN_PQUSER();
extern void HB_FUN_PQPASS();
extern void HB_FUN_PQHOST();


//extern void HB_FUN_PQDB();
extern void HB_FUN_PQRESET();
extern void HB_FUN_PQSETDBLOGIN ();

