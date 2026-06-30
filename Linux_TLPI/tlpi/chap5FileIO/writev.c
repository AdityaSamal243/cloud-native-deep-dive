#include<stdio.h>
#include<sys/uio.h>
#include<errno.h>
#include<fcntl.h>
#include<unistd.h>
#include<string.h>

int main(){
    char work[]="readv writev\n";
    char desc[]="for File I/O to/from multiple buffer\n";
    char foot[]="okay byee\n";

    struct iovec iov[3];

    int fd = open("result.txt",O_CREAT| O_WRONLY , 0644);
    if(fd==-1){
        perror("error");
        return 1;
    }

    iov[0].iov_base= work;
    iov[0].iov_len = strlen(work);

    iov[1].iov_base = desc;
    iov[1].iov_len = strlen(desc);

    iov[2].iov_base = foot;
    iov[2].iov_len = strlen(foot);

    ssize_t totalWritten = writev(fd,iov,3);

    if(totalWritten == -1){
        perror("error in writting");
        close(fd);
        return 1;
    }
    printf("successfully wrote %ld bytes\n",totalWritten);
    close(fd);
    return 0;
}
