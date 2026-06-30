#include<stdio.h>
#include<fcntl.h>
#include<unistd.h>
#include<errno.h>

int main(){
	int fd1 = open("abc1.txt",O_CREAT|O_WRONLY|O_EXCL);
	if(fd1==-1){
		perror("error is :");
	}
	int fd2= open("abc1.txt",O_CREAT|O_WRONLY|O_EXCL);
	if(fd2==-1){
		perror("error is :");
	}
	close(fd1);
	close(fd2);
}
