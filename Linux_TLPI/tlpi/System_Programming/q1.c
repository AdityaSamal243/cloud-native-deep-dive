#include<stdio.h>
#include<fcntl.h>
#include<unistd.h>
int main(){
	int fd = open("file1.txt",O_RDONLY|O_CREAT);
	printf("fd is: %d\n",fd);
	close(fd);
	return 0;
}

