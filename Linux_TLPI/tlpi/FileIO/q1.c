#include<stdio.h>
#include<fcntl.h>
#include<errno.h>
#include<unistd.h>
#include<string.h>

#define BUFFER 1024

int main(){
	int fd = open("abc.txt",O_WRONLY | O_APPEND);
	if(fd==-1){
		perror("error is:");
		return 1;
	}
	char buf[] = "good day\n";
	if(write(fd,buf,strlen(buf))==-1){
		perror("error is:");
		return 1;
	}
	close(fd);
	return 0;
}

