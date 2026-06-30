#include<stdio.h>
#include<errno.h>
#include<unistd.h>


int main(){
    int sizeleft = truncate("result.txt",20);
    if(sizeleft==-1){
        perror("can't truncate");
    }
    return 0;
}
