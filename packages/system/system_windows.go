// +build windows

package system

import (
	"os"
)

/*
#include <windows.h>
#include <stdio.h>
#include <stdlib.h>
#include <TlHelp32.h>

void kill_childproc( DWORD myprocID) {
	PROCESSENTRY32 pe;

	memset(&pe, 0, sizeof(PROCESSENTRY32));
	pe.dwSize = sizeof(PROCESSENTRY32);

	HANDLE hSnap = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0);
	if (Process32First(hSnap, &pe))
	{
	    BOOL bContinue = TRUE;
	
	    while (bContinue)
	    {
	        if (pe.th32ParentProcessID == myprocID && memcmp( pe.szExeFile, "tmp_dcoin", 9 ) != 0 &&
				memcmp(pe.szExeFile, "dcoin", 5) != 0)
	        {
	            HANDLE hChildProc = OpenProcess(PROCESS_ALL_ACCESS, FALSE, pe.th32ProcessID);
	
	            if (hChildProc)
	            {
					kill_childproc(GetProcessId(hChildProc));
	                TerminateProcess(hChildProc, 1);
	                CloseHandle(hChildProc);
	            }               
	        }
	        bContinue = Process32Next(hSnap, &pe);
	    }	
	}
}
*/
import "C"

// lstrcmp( pe.szExeFile, TEXT("tmp_dcoin.exe")) != 0 && lstrcmp( pe.szExeFile, TEXT("dcoin.exe")) != 0
				
func killChildProc() {
	C.kill_childproc( C.DWORD(os.Getpid()))
}

