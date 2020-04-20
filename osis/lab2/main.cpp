#include <stdlib.h>
#include <locale.h>  
#include <stdio.h>
#include <windows.h>
#include <tchar.h>

BOOL FindFileByClaster(TCHAR* volume,LONGLONG cluster){

    HANDLE hDevice = CreateFile(volume,      
        GENERIC_READ ,                       
        FILE_SHARE_READ | FILE_SHARE_WRITE,  
        NULL,                                
        OPEN_EXISTING,                       
        FILE_FLAG_BACKUP_SEMANTICS,        
        NULL);

    if(hDevice == INVALID_HANDLE_VALUE)
    {   
          return FALSE;
    }

    LOOKUP_STREAM_FROM_CLUSTER_INPUT input={0};
    input.NumberOfClusters = 1;
    input.Cluster[0].QuadPart = cluster;        

    BYTE output[5000]={};
    DWORD dwRes=0;
    LOOKUP_STREAM_FROM_CLUSTER_OUTPUT result={0};   

    BOOL bRes = DeviceIoControl( (HANDLE)       hDevice,
                 FSCTL_LOOKUP_STREAM_FROM_CLUSTER, 
                 (LPVOID)       &input,        
                 (DWORD)        sizeof(input),     
                 (LPVOID)       output,      
                 (DWORD)        5000,    
                 (LPDWORD)      &dwRes,   
                 NULL ); 

    if(bRes == FALSE){      
          return FALSE;
    }

    memcpy(&result,output,sizeof(LOOKUP_STREAM_FROM_CLUSTER_OUTPUT));

    if(result.NumberOfMatches == 0){
        wprintf( L"Файл не найден\n");
        return FALSE;
    }   

    wprintf( L"Информация о файле\n");

    BYTE* p = (BYTE*)output + result.Offset;
    LOOKUP_STREAM_FROM_CLUSTER_ENTRY* pentry = (LOOKUP_STREAM_FROM_CLUSTER_ENTRY*)p;    

    wprintf( L"Flags: 0x%x ",(UINT)pentry->Flags);

    if((pentry->Flags & LOOKUP_STREAM_FROM_CLUSTER_ENTRY_FLAG_PAGE_FILE) > 0) wprintf(L"(Pagefile)");
    else if((pentry->Flags & LOOKUP_STREAM_FROM_CLUSTER_ENTRY_FLAG_FS_SYSTEM_FILE) > 0)  wprintf(L"(Internal filesystem file)");
    else if((pentry->Flags & LOOKUP_STREAM_FROM_CLUSTER_ENTRY_FLAG_TXF_SYSTEM_FILE) > 0) wprintf(L"(Internal TXF file)");
    else wprintf(L"(Normal file)"); 

    wprintf( L"\nFile: %s\n",pentry->FileName); 
    return TRUE;
}

int _tmain(int argc, _TCHAR* argv[])
{
    setlocale(LC_ALL,"Russian");

    LONGLONG inp=0;
    wprintf( L"Введите номер кластера: \n");
    scanf("%llu",&inp);

    FindFileByClaster(L"\\\\.\\C:",inp);        

    system("PAUSE");
    return 0;
}