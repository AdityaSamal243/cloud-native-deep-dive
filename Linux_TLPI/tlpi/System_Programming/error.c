#include<stdio.h>
#include<errno.h>
#include<fcntl.h>
#include<unistd.h>
#include<string.h>
int main(){
	int fd = open("baby.txt",O_RDONLY);
	if(fd==-1){
		printf("error occured\n");
		printf("errno :%d\n",errno);
		printf("error: %s",strerror(errno));
		return 1;
	}
	printf("file opened");
	close(fd);
	return 0;
}
