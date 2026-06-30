#include<fcntl.h>
#include<stdio.h>
#include<unistd.h>
#include<errno.h>
#include<stdlib.h>
#define BUFFER 20

int main(){
	char buf[BUFFER];
	int fd1=open("source.dat",O_RDONLY);
	if(fd1==-1){
		perror("error read:");
		return 1;
	}
	int fd2=open("replica.dat",O_CREAT | O_EXCL | O_WRONLY | O_SYNC, 0644);
	if(fd2==-1){
		perror("error open:");
		close(fd1);
		return 1;
	}
	ssize_t numread;
	while((numread=read(fd1,buf,BUFFER))>0){
		int totalwritten =0;
		while(totalwritten<numread){
			ssize_t written = write(fd2,buf+totalwritten,numread-totalwritten);
			if(written==-1){
				perror("error in write");
				close(fd2);
				close(fd1);
				return 1;
			}
			totalwritten+=written;
		}

	}
	if(numread==-1){
		perror("error read:");
	}

	close(fd1);
	close(fd2);
	return 0;
}



