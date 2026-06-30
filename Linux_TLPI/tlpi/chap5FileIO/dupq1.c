#include<stdio.h>
#include<fcntl.h>
#include<unistd.h>
#include<errno.h>
#include<string.h>

int main(){
	int fd = open("hello.txt",O_WRONLY | O_CREAT | O_EXCL,00777);
	if(fd==-1){
		perror("open");
		return 1;
	}
	int newfd=dup2(fd,0);
	char buffer[] = "hello darling";
	int written=write(newfd,buffer,strlen(buffer));
	if(written==-1){
               perror("write");
	       return 1;
	}
	close(newfd);
	close(fd);
	return 0;
}
	

