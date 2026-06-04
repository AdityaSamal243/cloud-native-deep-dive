#include<stdio.h>
#include<fcntl.h>
#include<unistd.h>
#include<errno.h>

int main(){
    int fd = open("hey.txt",O_RDONLY );
    if(fd == -1){
        perror("open");
        close(fd);
        return -1;
    }
    int flags = fcntl(fd, F_GETFL);
    int accessMode = flags & O_ACCMODE;
    if(accessMode == O_WRONLY){
        printf("File is opened in write only mode\n");
    }
    else{
        printf("File is not opened in write only mode\n");
    }
    close(fd);
    return 0;
}