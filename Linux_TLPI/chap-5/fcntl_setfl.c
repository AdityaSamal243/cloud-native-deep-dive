#include<stdio.h>
#include<fcntl.h>
#include<unistd.h>
#include<errno.h>

int main(){
    int fd = open("hii.txt",O_CREAT |O_RDWR, S_IRUSR | S_IWUSR);
    if(fd==-1){
        perror("open");
        close(fd);
        return -1;
    }
    int flags = fcntl(fd,F_GETFL); //retrieve copy of existing flags
    if(flags == -1){
        perror("fcntl");
        return -1;
    }
    flags= flags | O_APPEND;
    int res = fcntl(fd,F_SETFL,flags); //set new flags
    if(res==-1){
        perror("fcntl");
        return -1;
    }

    int fla = fcntl(fd,F_GETFL);
    if(fla == -1){
        perror("fcntl");
        return -1;
    }
    if(fla & O_APPEND){
        printf("O_APPEND flag is set\n");
    }
    else{
        printf("O_APPEND flag is not set\n");
    }
    close(fd);
    return 0;
}
